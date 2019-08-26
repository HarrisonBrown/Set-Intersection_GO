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

func isNumeric(s string) bool {

	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func countEntriesInFile(filename string) (uint, uint) {

	var entryCount uint
	distinctEntrySet := make(map[string]bool)

	// Open file into an os.File - f
	f, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scannedLine := scanner.Text(); isNumeric(scannedLine) {
			distinctEntrySet[scannedLine] = true
			fmt.Println(scannedLine)
			entryCount++
		}
	}

	check(scanner.Err())

	f.Close()

	return entryCount, uint(len(distinctEntrySet))
}

func main() {

	filename := "testData/A_f.csv"

	numEntries, numDistinctEntries := countEntriesInFile(filename)

	fmt.Printf("%d entries in file %q, of which %d are distinct.", numEntries, filename, numDistinctEntries)
}
