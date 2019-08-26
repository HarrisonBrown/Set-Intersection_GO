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

func countEntriesInFile(filename string) uint {

	var entryCount uint = 0

	// Open file into an os.File - f
	f, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scannedLine := scanner.Text(); isNumeric(scannedLine) {
			fmt.Println(scannedLine)
			entryCount++
		}
	}

	check(scanner.Err())

	f.Close()

	return entryCount
}

func main() {
	filename := "testData/short.csv"
	numEntries := countEntriesInFile(filename)

	fmt.Printf("%d entries in file %q.", numEntries, filename)
}
