package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {
	fileHandle, err := os.Open("measurements.txt")
	if err != nil {
		log.Println(err)
	}
	defer fileHandle.Close()

	line := 0
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		line++
	}
	err = scanner.Err()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(line)
}