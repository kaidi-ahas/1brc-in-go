package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	start := time.Now()

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

	fmt.Println("took: ", time.Since(start))
}