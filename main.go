package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {

	if e != nil {
		log.Fatal(e)
	}
}

func isNumeric(s *string) bool {

	_, err := strconv.ParseFloat(*s, 64)
	return err == nil
}

func countEntriesInFile(filename string) (uint, map[string]uint8) {

	var entryCount uint
	distinctEntrySet := make(map[string]uint8)

	// Open file into an os.File - f
	f, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scannedLine := scanner.Text(); isNumeric(&scannedLine) {
			distinctEntrySet[scannedLine]++
			//fmt.Println(scannedLine)
			entryCount++
		}
	}

	check(scanner.Err())

	f.Close()

	return entryCount, distinctEntrySet
}

func main() {

	filenameA := "testData/short_a.csv"
	filenameB := "testData/short_b.csv"

	numEntriesA, distinctEntriesA := countEntriesInFile(filenameA)
	numEntriesB, distinctEntriesB := countEntriesInFile(filenameB)

	fmt.Printf("File %q contains %d entries, of which %d are distinct.\n", filenameA, numEntriesA, len(distinctEntriesA))
	fmt.Printf("File %q contains %d entries, of which %d are distinct.\n", filenameB, numEntriesB, len(distinctEntriesB))

	// Find overlap
	var distinctOverlap uint
	var totalOverlap uint
	for entry, occurences := range distinctEntriesA {
		if distinctEntriesB[entry] > 0 {
			distinctOverlap++
			totalOverlap += uint(occurences) + uint(distinctEntriesB[entry])
		}
	}

	fmt.Printf("Total Overlap: %d, Distinct Overlap: %d\n", totalOverlap, distinctOverlap)
}
