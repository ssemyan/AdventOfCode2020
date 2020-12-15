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

	fmt.Printf("Results: 1=%d 3=%d final=%d \n", nOne, nThree, nOne*nThree)

	// Part Two
	//nCombos := 1
	foundCombos := make(map[int]int)
	nCombos := FindNextAdaptor(lines, 0, 0, foundCombos)

	fmt.Println("Total Combos: ", nCombos+1)

}

// FindNextAdaptor - find and count the next adapter
func FindNextAdaptor(lines []int, nCurrVolts int, lineIndex int, foundCombos map[int]int) int {

	// See if we know how many combos from here
	val, exists := foundCombos[nCurrVolts]
	if exists {
		//fmt.Println("found ", val)
		return val
	}

	bOnFirst := true
	nCombos := 0
	for {
		// Check if at end
		if lineIndex >= len(lines) {
			break
		}

		// Is the next match in range
		if (lines[lineIndex] - nCurrVolts) < 4 {
			//fmt.Printf(" %d", lines[lineIndex])

			// is it the first match
			if bOnFirst {
				bOnFirst = false
			} else {
				nCombos++
				//fmt.Printf(" break")
			}

			nCombos += FindNextAdaptor(lines, lines[lineIndex], lineIndex+1, foundCombos)
		}
		lineIndex++
	}

	// Save combo
	foundCombos[nCurrVolts] = nCombos
	return nCombos
}
