package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// goal read the temperature measurements per weather station, aggregate the statistics and print to the standard output

type Stats struct {
	Min float64
	Max float64
	Sum float64
	Count int
}

func (s *Stats) Add(value float64) {
	if s.Count == 0 {
		s.Min = value
		s.Max = value
	}
	if s.Min > value {
		s.Min = value
	}
	if s.Max < value {
		s.Max = value
	}
	s.Count++
	s.Sum+=value
}

func (s *Stats) Avg() float64{
	return s.Sum / float64(s.Count)
}

func main() {
	path := "measurements.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open the file %s. %s", path, err)
	}
	defer file.Close()

	measurement, err := readMeasurements(file)
	if err != nil {
		log.Println(err)
	}

	printResults(measurement)
}

func readMeasurements(r io.Reader) (map[string]*Stats, error) {
	measurement := make(map[string]*Stats)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		station, temperature, err := parseLine(line)
		if err != nil {
			continue
		}

		s := measurement[station]
		if s == nil {
			s = &Stats{}
			measurement[station] = s
		}
		s.Add(temperature)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return measurement, nil
}

func parseLine(line string) (station string, temperature float64, err error) {
	station, value, found := strings.Cut(line, ";")
	if !found {
		log.Printf("bad line: %s", line)
	}

	temperature, err = strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("failed to parse %s. Err: %v", value, err)
	}
	return station, temperature, nil
}

func printResults(measurement map[string]*Stats) {
	var stations []string

	for s := range measurement {
		stations = append(stations, s)
	}

	fmt.Println("{")
	for station, stat := range measurement {
		avg := stat.Avg()
		fmt.Printf("%s=%.1f/%.1f/%.1f,\n", station, stat.Min, avg, stat.Max)
	}
	fmt.Println("}")
}