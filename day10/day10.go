package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"sort"
)

func main() {

	// Read all data into memory
	lines := fileload.FileLoadInts("day10/data.txt")

	// Part One
	sort.Ints(lines)

	nCurrVolts, nOne, nThree := 0, 0, 0

	for _, line := range lines {
		fmt.Println(nCurrVolts, line)
		switch line - nCurrVolts {
		case 1:
			nOne++
		case 2:
			panic("Should not get here")
		case 3:
			nThree++
		}
		nCurrVolts = line
	}

	// Last is always 3 away
	nThree++

	fmt.Printf("Results: 1=%d 3=%d final=%d ", nOne, nThree, nOne*nThree)

}
