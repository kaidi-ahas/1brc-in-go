# 1BRC with Go

Implementation of the **1 Billion Row Challenge** in Go.

Original challenge: [1 Billion Row Challenge](https://github.com/gunnarmorling/1brc)

---

### Problem

Given a file containing temperature measurements in the format:
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

Print results in sorted order:
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

The main goal of this project is not code correctness, but performance analysis and optimization at large scale (1 billion rows).

---

### Project structure

```go
.
├── main.go // orchestration + timing + profiling
├── measurement.go // parsing logic
├── reader.go // streaming + scanning
├── stats.go // aggreggation logic
```

### Design overview
* `ParseLine` handles parsing and validation
* `ReadMeasurements` streams the file line-by-line using `bufio.Scanner`
* `StationStats` encapsulates aggregation state
* `Stats` tracks min, max, sum, and count
* `main` wires everything together and enables CPU profiling

Current implementation is **sequential** (single-threaded).

---

### How to run

Place `measurements.txt` file in the project root.

Dataset generation instructions are available in the original challenge repository: [1 Billion Row Challenge](https://github.com/gunnarmorling/1brc)

Build and run:
```bash
go build -o 1brc
./1brc
```

Measure execution time:
```bash
/usr/bin/time -l ./1brc
```
---

## Performance baseline (Sequential Version)

Dataset: 1,000,000,000 rows<br> Implementation: sequential<br>
Machine: (your machine here)

### Runtime
```
real    54.28s
user    51.99s
sys     2.93s
```

Throughput:<br>
~18.4 million rows/second

---

### CPU profiling

Top output:
```
51.84s  96.86%  syscall.syscall
```

Interpretation:

- ~97% of execution time spent in syscall.syscall
- The workload is strongly IO-bound
- Compute logic (parsing + aggregation) is negligible relative to input reading

---

### Memory profiling

#### In-use memory (heap)

~3.2 MB retained at program end
- No memory leak
- Aggregation map remains small

#### Total allocations
15.52 GB allocated during execution
100% attributed to:

```
bufio.(*Scanner).Text
```

This indicates:
* One string allocation per input line
* 1B string allocations
* Significant GC churn despite small retained heap

### Peak RSS (OS-level memory usage)

~9 MB

The program maintains a very small real memory footprint.

---

### Performance baseline summary

* The implementation is IO-bound.
* Major allocation hotspot: Scanner.Text().
* Memory retention is minimal.
* Allocation churn is extremely high (15GB).
* GC likely contributes overhead.

---

## Next optimization steps

Replace bufio.Scanner with bufio.Reader to:
* Eliminate per-line string allocations
* Reduce total allocations
* Reduce GC pressure
* Potentially improve throughput

**Future improvements will be measured against this baseline.**