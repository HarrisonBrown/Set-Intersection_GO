package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Takes an error and fatally logs if it warrants it
func check(e error) {

	if e != nil {
		log.Fatal(e)
	}
}

// Returns true/false whether the given string is numeric
func isNumeric(s *string) bool {

	_, err := strconv.ParseFloat(*s, 64)
	return err == nil
}

// Given a filepath, returns:
//    - a count of all udprn values
//    - a map of distinct udprn values to their number of occurences
func countEntriesInFile(filename string) (uint, map[string]uint8) {

	// Output data
	var entryCount uint
	distinctEntrySet := make(map[string]uint8)

	// Open provided filepath into an os.File - f
	f, err := os.Open(filename)
	check(err)

	// Scan each line, checking if it is numeric, increment the entry counters if so
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scannedLine := scanner.Text(); isNumeric(&scannedLine) {
			distinctEntrySet[scannedLine]++
			entryCount++
		}

		// Check for any non-EOF errors
		check(scanner.Err())
	}

	f.Close()

	return entryCount, distinctEntrySet
}

// Given two maps of udprn entries to their occurence count, return the number of distinct and total overlaps
func findOverlap(entriesA map[string]uint8, entriesB map[string]uint8) (uint, uint) {
	// Find overlap
	var distinctOverlap uint
	var totalOverlap uint
	for entry, occurences := range entriesA {
		if entriesB[entry] > 0 {
			distinctOverlap++
			totalOverlap += uint(occurences) + uint(entriesB[entry])
		}
	}

	return distinctOverlap, totalOverlap
}

func main() {

	filenameA := "testData/short_a.csv"
	filenameB := "testData/short_b.csv"

	numEntriesA, distinctEntriesA := countEntriesInFile(filenameA)
	numEntriesB, distinctEntriesB := countEntriesInFile(filenameB)
	distinctOverlap, totalOverlap := findOverlap(distinctEntriesA, distinctEntriesB)

	fmt.Printf("File %q contains %d entries, of which %d are distinct.\n", filenameA, numEntriesA, len(distinctEntriesA))
	fmt.Printf("File %q contains %d entries, of which %d are distinct.\n", filenameB, numEntriesB, len(distinctEntriesB))
	fmt.Printf("Total overlap: %d, distinct overlap: %d\n", totalOverlap, distinctOverlap)
}
