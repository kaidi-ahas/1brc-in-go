package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Goal: Parse one line
// create a struct for measurement
// change countLines to return the lines slice

func main() {
	line := "Hamburg;12.3"
	m, err := parseLine(line)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%+v\n", m)

}

type Measurement struct {
	Station string
	Temperature float64
}

func parseLine(line string) (Measurement, error) {
	station, value, found := strings.Cut(line, ";")
	if !found {
		return Measurement{}, fmt.Errorf("Bad line: %q", line)
	}

	temperature, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return Measurement{}, fmt.Errorf("invalid temperature: %q: %w", value, err)
	}

	return Measurement{
		Station: station,
		Temperature: temperature,
	}, nil
}






// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type Stats struct {
// 	Min float64
// 	Max float64
// 	Sum float64
// 	Count int
// }

// func (s *Stats) Add(value float64) {

// 	if s.Count == 0 {
// 		s.Min = value
// 		s.Max = value
// 	} else {
// 		if value < s.Min {
// 			s.Min = value
// 		}
// 		if value > s.Max {
// 			s.Max = value
// 		}
// 	}
// 	s.Sum += value
// 	s.Count++
// }

// func main() {
// 	fmt.Println(readMeasurements("measurements.txt"))

// }

// func readMeasurements(path string) map[string]*Stats {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
	
// 	statsByStation := make(map[string]*Stats)

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		station, value, found := strings.Cut(line, ";")
// 		if !found {
// 			log.Printf("bad line: %q", line)
// 			continue
// 		}
		
// 		temperature, err := strconv.ParseFloat(value, 64)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}

// 		stats, exists := statsByStation[station]

// 		if !exists {
// 			stats = &Stats{}
// 			statsByStation[station]=stats
// 		}

// 		stats.Add(temperature)

// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Println(err)
// 	}
// 	return statsByStation
// }

// func printStats(stats map[string]*Stats) {
	
// }