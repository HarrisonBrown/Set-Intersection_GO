package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// Takes an error and fatally logs if it warrants it
func check(e error) {

	if e != nil {
		log.Fatal(e)
	}
}

// Asks the user for two file paths consecutively, returns each as individual string
func promptAndReadFilePaths() (string, string) {

	var filepaths [2]string

	for i := 0; i < 2; i++ {
		fmt.Printf("Filepath %d: ", i+1)
		scanner := bufio.NewReader(os.Stdin)
		filepaths[i], _ = scanner.ReadString('\n')

		// Remove line ending from filepath
		filepaths[i] = strings.TrimSuffix(filepaths[i], "\n")

		// Windows lines are terminated with \r\n, so remove \r too
		if runtime.GOOS == "windows" {
			filepaths[i] = strings.TrimSuffix(filepaths[i], "\r")
		}
	}

	return filepaths[0], filepaths[1]
}

// Given a filepath, returns:
//    - a count of all udprn values
//    - a map of distinct udprn values to their number of occurences
func countEntriesInFile(filepath string) (uint, map[uint32]uint8) {

	// Output data
	var entryCount uint
	distinctEntrySet := make(map[uint32]uint8)

	// Open provided filepath into an os.File - f
	f, err := os.Open(filepath)
	check(err)

	// Scan each line, attempting to convert it to a numeric value,
	// Store the value as a key in the map and increment the entry counters if so
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		scannedLine := scanner.Text()
		if value, err := strconv.ParseUint(scannedLine, 10, 32); err == nil {
			distinctEntrySet[uint32(value)]++
			entryCount++
		}

		// Check for any non-EOF errors
		check(scanner.Err())
	}

	f.Close()

	return entryCount, distinctEntrySet
}

// Given two maps of udprn entries to their occurence count, return the number of distinct and total overlaps
func findOverlap(entriesA map[uint32]uint8, entriesB map[uint32]uint8) (uint, uint) {
	// Find overlap
	var distinctOverlap uint
	var totalOverlap uint
	for entry, occurences := range entriesA {
		if entriesB[entry] > 0 {
			distinctOverlap++
			totalOverlap += uint(occurences) * uint(entriesB[entry])
		}
	}

	return distinctOverlap, totalOverlap
}

func main() {

	numArgs := len(os.Args[1:])

	if numArgs == 1 {

		// Only one argument passed, treat it as just one filepath to be analysed
		filepathA := os.Args[1]

		numEntriesA, distinctEntriesA := countEntriesInFile(filepathA)
		fmt.Printf("File %q contains %d entries, of which %d are distinct.\n", filepathA, numEntriesA, len(distinctEntriesA))

	} else {

		var filepathA, filepathB string

		if numArgs >= 2 {

			// At least two arguments passed, treat the first two as filepaths to be analysed and ignore the rest
			filepathA = os.Args[1]
			filepathB = os.Args[2]

		} else {

			// No arguments passed, prompt the user for some filepaths
			filepathA, filepathB = promptAndReadFilePaths()
		}

		numEntriesA, distinctEntriesA := countEntriesInFile(filepathA)
		numEntriesB, distinctEntriesB := countEntriesInFile(filepathB)
		distinctOverlap, totalOverlap := findOverlap(distinctEntriesA, distinctEntriesB)

		fmt.Printf("File %q contains %d entries, of which %d are distinct.\n", filepathA, numEntriesA, len(distinctEntriesA))
		fmt.Printf("File %q contains %d entries, of which %d are distinct.\n", filepathB, numEntriesB, len(distinctEntriesB))
		fmt.Printf("Total overlap: %d, distinct overlap: %d\n", totalOverlap, distinctOverlap)
	}
}
