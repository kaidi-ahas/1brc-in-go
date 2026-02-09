package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		stationData[station]= append(stationData[station], temp)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	
	fmt.Println(stationData)
}