# 1BRC in Go

Implementation of the **1 Billion Row Challenge** in Go.

Original challenge: [1 Billion Row Challenge](https://github.com/gunnarmorling/1brc)

---

### Problem

Given a file containing temperature measurements for weather stations in the format:
```
Hamburg;12.3
Munich;3.6
Paris;13.8
Munich;6.1
```

Compute, for each station:
```
min/average/max
```

Print the result to the standard output in this format:
```
{
    Upington=-26.2/20.4/69.4,
    Honiara=-24.4/26.5/75.4,
    Chittagong=-24.8/25.9/76.6,
    Cotonou=-25.2/27.2/78.2,
    ...
    Skopje=-36.7/12.4/64.7,
    Stockholm=-41.8/6.6/57.6,
    Burnie=-34.7/13.1/61.0,
}
```

The main goal is to process very large datasets while measuring and understanding real performance bottlenecks.

---

### My Project Structure

```go
.
├── main.go // orchestration + timing + profiling
├── measurement.go // parsing logic
├── reader.go // streaming + scanning
├── stats.go // aggreggation logic
```

### Design Overview
* `ParseLine` handles parsing and validation
* `ReadMeasurements` streams the file line-by-line using `bufio.Scanner`
* `StationStats` encapsulates aggregation state
* `Stats` tracks min, max, sum, and count
* `main` wires everything together and enables CPU profiling

---

### How to Run

This implementation expects a `measurements.txt` file in the project root.

The dataset can be generated following the instructions in the Original challenge: [1 Billion Row Challenge](https://github.com/gunnarmorling/1brc)

Then run:
```bash
go run .
```

To measure runtime:
```bash
time go run .
```

---

### CPU Profiling

CPU profiling is enabled in `main.go`:
```go
f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

After execution:
```
go tool pprof cpu.prof
```

Then:
```
(pprof) top
```

---

### Profiling Results

Example output:

```
flat    flat%   function
49.38s  96.48%  syscall.syscall
```

Full output:
```
Showing nodes accounting for 50.17s, 98.03% of 51.18s total
      flat  flat%   sum%        cum   cum%
    49.38s 96.48% 96.48%     49.38s 96.48%  syscall.syscall
```

### Interpretation

* ~96% of CPU time is spent in syscall.syscall
* The workload is strongly IO-bound
* Parsing and aggregation contribute very little relative cost
* Optimizing compute logic would not significantly improve performance without addressing IO

Profiling changed the optimization direction.

Initial intuition suggested parsing would dominate.<br>
Measurement proved otherwise.

---

### Runtime Example

On my machine:
```
real    1m6s
user    1m16s
sys     0m4s
```
The higher `user` time compared to `real` indicates multi-core CPU utilization during execution.

---

### Implementation Details

#### Parsing

```go
station, value, found := strings.Cut(line, ";")
temperature, err := strconv.ParseFloat(value, 64)
```

#### Streaming

```go
scanner := bufio.NewScanner(r)
```

The file is processed incrementally — no full-file loading.

#### Aggregation

Each station maintains:
* Min
* Max
* Sum
* Count

Average is computed on demand:
```go
func (s *Stats) Avg() float64 {
    return s.Sum / float64(s.Count)
}
```

### Next Steps

Planned experiments:
* Bounded concurrency (worker pool)
* Alternative IO strategies
* Allocation reduction during parsing
* Benchmark comparisons