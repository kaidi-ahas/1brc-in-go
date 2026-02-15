package main

import (
	"bufio"
	"fmt"
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

	// name it to a measurement
	// one measurement is station and it's statistics
	stationMeasurements := make(map[string]*Stats)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// rename separated to parts or fields and use Cut instead
		station, value, found := strings.Cut(line, ";")
		if !found {
			log.Printf("bad line: %s", line)
			continue
		}
		// Create a separate function for parsing
		temperature, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Printf("failed to parse %s. Err: %v", value, err)
			continue
		}
		s, exists := stationMeasurements[station]
		if !exists {
			s = &Stats{}
			stationMeasurements[station] = s
		}
		s.Add(temperature)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	var stations []string

	for s := range stationMeasurements {
		stations = append(stations, s)
	}


	fmt.Println("{")
	for station, stat := range stationMeasurements {
		avg := stat.Avg()
		fmt.Printf("%s=%.1f/%.1f/%.1f,\n", station, stat.Min, avg, stat.Max)
	}
	fmt.Println("}")
}