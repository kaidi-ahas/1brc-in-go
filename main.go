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
	if value < s.Min {
		s.Min = value
	}
	if value > s.Max {
		s.Max = value
	}
	s.Sum += value
	s.Count++	
}

func main() {
	file, err := os.Open("measurements.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	
	statsByStation := make(map[string]*Stats)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")
		if len(fields) != 2 {
			log.Printf("bad line: %q", line)
			continue
		}
		
		temp, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Println(err)
			continue
		}

		stats := &Stats{}

		stats.Add(temp)

		station := fields[0]

		statsByStation[station] = stats
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println("}")

}
