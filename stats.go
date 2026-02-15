package main

import (
	"fmt"
	"io"
)

type StationStats struct {
	Data map[string]*Stats
}

type Stats struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func NewStationStats() *StationStats {
	return &StationStats{
		Data: make(map[string]*Stats),
	}
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
	s.Sum += value
}

func (s *Stats) Avg() float64 {
	return s.Sum / float64(s.Count)
}

func (ss *StationStats) Add(m Measurement) {
	stats := ss.Data[m.Station]
	if stats == nil {
		stats = &Stats{}
		ss.Data[m.Station] = stats
	}
	stats.Add(m.Temperature)
}

func (ss *StationStats) Results() map[string]*Stats {
	return ss.Data
}

func (ss *StationStats) Print(w io.Writer) {
	fmt.Fprintln(w, "{")
	for station, stat := range ss.Results() {
		avg := stat.Avg()
		fmt.Fprintf(w, "%s=%.1f/%.1f/%.1f,\n", station, stat.Min, avg, stat.Max)
	}
	fmt.Fprintln(w, "}")
}