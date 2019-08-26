package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func countEntriesInFile(filename string) uint {

	var entryCount uint = 0

	// Open file into an os.File - f
	f, err := os.Open(filename)
	check(err)

	reader := bufio.NewReader(f)
	for {
		udprn, err := reader.ReadBytes('\n')

		if len(udprn) > 2 {
			fmt.Printf("udprn: %s\n", udprn)
			entryCount++
		}

		if err != nil {
			fmt.Println(err)
			break
		}
	}
	f.Close()

	return entryCount
}

func main() {

	filename := "testData/short.csv"
	numEntries := countEntriesInFile(filename)

	fmt.Printf("%d entries in file %q.", numEntries, filename)
}
