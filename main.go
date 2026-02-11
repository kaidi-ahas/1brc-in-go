package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	} else {
		if value < s.Min {
			s.Min = value
		}
		if value > s.Max {
			s.Max = value
		}
	}
	s.Sum += value
	s.Count++
}

func main() {
	fmt.Println(readMeasurements("measurements.txt"))

}

func readMeasurements(path string) map[string]*Stats {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	statsByStation := make(map[string]*Stats)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		station, value, found := strings.Cut(line, ";")
		if !found {
			log.Printf("bad line: %q", line)
			continue
		}
		
		temperature, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println(err)
			continue
		}

		stats, exists := statsByStation[station]

		if !exists {
			stats = &Stats{}
			statsByStation[station]=stats
		}

		stats.Add(temperature)

	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return statsByStation
}

func printStats(stats map[string]*Stats) {
	
}