package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

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

	mf, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer mf.Close()

	runtime.GC()

	if err := pprof.WriteHeapProfile(mf); err != nil {
		log.Fatal(err)
	}

	fmt.Println("took: ", time.Since(start))
}
