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
	invalidNum := 0

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
			invalidNum = lines[nCurrent]
			break
		}
		nCurrent++
	}

	// part two
	nMin, nMax := 0, 0
	nCurrent = 0
	for {
		if nCurrent >= len(lines) {
			break
		}

		nStart := nCurrent
		nEnd := len(lines)
		bFoundCombo := false
		for i := nStart; i < nEnd; i++ {
			nSum := lines[i]
			for z := i + 1; z < nEnd; z++ {
				nSum += lines[z]
				if invalidNum == nSum {
					fmt.Println("Found contiguous nums: ", i, z)
					nMin = lines[i] // give it something other than 0
					for y := i; y <= z; y++ {
						if lines[y] < nMin {
							nMin = lines[y]
						}
						if lines[y] > nMax {
							nMax = lines[y]
						}
					}
					bFoundCombo = true
					break
				}
				if nSum > invalidNum {
					break
				}
			}
			if bFoundCombo {
				break
			}
		}
		if bFoundCombo {
			fmt.Println("Found lines ", nMin, nMax, (nMin + nMax))
			break
		}
		nCurrent++
	}

}
