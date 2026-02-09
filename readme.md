# How fast can you aggregate huge amounts of data with minimal memory and CPU waste?

file with rows:
    station;temperature_float

1. Read the file
Line by line (streaming — not loading the whole file into memory).

2. Group by station name
Each station accumulates:
min
max
sum
count

map the station and all the temperatures

"station": [1.3, 2.7, 3.0, 4.3, 4.2]
map type:
key - string
value - []float64

For a slice of floats you must compute for each station:
1. minimum temperature
slices.Min()

2. maximum temperature
slices.Max()

3. sum
range over slice and build the sum

4. count
map can give a count

5. average temperature
avg = sum / count

Print the result to stdout in this format:

One line per station, like this:
{
Station=min/avg/max, 
Station=min/avg/max, 
...}