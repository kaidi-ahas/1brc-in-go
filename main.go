package main

import (
	"log"
	"os"
)

// goal read the temperature measurements per weather station, aggregate the statistics and print to the standard output

func main() {
	path := "measurements.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open the file %s. %s", path, err)
	}
	defer file.Close()

	ss := NewStationStats()

	err = ReadMeasurements(file, ss)
	if err != nil {
		log.Println(err)
	}

	ss.Print(os.Stdout)
}