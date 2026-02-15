package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// goal read the temperature measurements per weather station, aggregate the statistics and print to the standard output

type Measurement struct {
	Station string
	Temperature float64
}

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
	stationByStats := make(map[string]*Stats)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		measurement, err := parseLine(line)
		if err != nil {
			continue
		}
		
		station := measurement.Station
		temp := measurement.Temperature

		stats := stationByStats[station]
		if stats == nil {
			stats = &Stats{}
			stationByStats[station] = stats
		}
		stats.Add(temp)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return stationByStats, nil
}

func parseLine(line string) (m Measurement, err error) {
	station, value, found := strings.Cut(line, ";")
	if !found {
		return m, errors.New("bad line")
	}

	temperature, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return m, errors.New("failed to parse the value")
	}

	m.Station = station
	m.Temperature = temperature
	return m, nil
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