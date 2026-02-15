package main

import (
	"bufio"
	"io"
)

// scanning, parsing
func ReadMeasurements(r io.Reader, ss *StationStats) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		measurement, err := ParseLine(line)
		if err != nil {
			continue
		}
		ss.Add(measurement)
	}
	if err := scanner.Err(); err != nil {
		return scanner.Err()
	}
	return nil
}