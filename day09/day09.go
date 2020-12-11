package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
)

func main() {

	// Read all data into memory
	lines := fileload.FileLoadInts("day09/data.txt")

	// Part One
	preamble := 25
	nCurrent := preamble

	for {
		if nCurrent >= len(lines) {
			break
		}

		nStart := nCurrent - preamble
		nEnd := nCurrent - 1
		bFoundCombo := false
		for i := nStart; i < nEnd; i++ {
			for z := i + 1; z <= nEnd; z++ {
				if lines[nCurrent] == (lines[i] + lines[z]) {
					fmt.Println("Found combo: ", lines[nCurrent], i, z)
					bFoundCombo = true
					break
				}
			}
			if bFoundCombo {
				break
			}
		}
		if !bFoundCombo {
			fmt.Println("No combo found line ", nCurrent, lines[nCurrent])
			break
		}
		nCurrent++
	}
}
