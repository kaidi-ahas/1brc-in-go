package main

import (
	"errors"
	"strconv"
	"strings"
)

type Measurement struct {
	Station     string
	Temperature float64
}

func ParseLine(line string) (m Measurement, err error) {
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