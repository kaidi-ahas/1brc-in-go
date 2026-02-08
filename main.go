package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {
	file, err := os.Open("measurements.txt")
	if err != nil {
		log.Println(err)
	}

	line := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
	}

	fmt.Println(line)
}