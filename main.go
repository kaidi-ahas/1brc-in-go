package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileHandle, err := os.Open("measurements.txt")
	if err != nil {
		log.Println(err)
	}
	defer fileHandle.Close()

	stationData := make(map[string][]float64)

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		line := scanner.Text()
		separated := strings.Split(line, ";")
		if len(separated) != 2 {
			log.Printf("bad line: %q", line)
			continue
		}
		temp, err := strconv.ParseFloat(separated[1], 64)
		if err != nil {
			log.Println(err)
			continue
		}

		station := separated[0]

		stationData[station] = append(stationData[station], temp)

		// accumulate min, max, ave, count for each station

	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	output := make(map[string][]float64)

	for station, temps := range stationData {
		output[station] = calculate(temps)
	}

	stations := make([]string, 0, len(output))
	for station := range output {
		stations = append(stations, station)
	}

	slices.Sort(stations)

	fmt.Print("{")
	for i, station := range stations {
		stats := output[station]

		if i > 0 {
			fmt.Print(", ")
		}

		fmt.Printf(
			"%s=%.1f/%.1f/%.1f",
			station,
			stats[0],
			stats[1],
			stats[2],
		)
	}
	fmt.Println("}")

}

func calculate(temps []float64) []float64 {
	var results []float64

	min := slices.Min(temps)
	max := slices.Max(temps)
	var sum float64
	for _, temp := range temps {
		sum += temp
	}
	count := len(temps)

	average := sum / float64(count)

	results = append(results, min, average, max)
	return results
}
