package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
)

func main() {

	// Read all data into string array
	slines := fileload.Fileload("day11/data.txt")

	// Convert to two rune arrays
	lines := make([][]rune, len(slines))
	newLines := make([][]rune, len(slines))
	for n, line := range slines {
		lines[n] = []rune(line)
		newLines[n] = []rune(line)
	}

	// Part One
	for {

		// Run rules
		bSomethingChanged := false
		for nLine, line := range lines {
			for nCol, col := range line {
				// Ignore floor
				if col != '.' {
					// Count adjacent
					nOccupied := 0
					for row := nLine - 1; row < nLine+2; row++ {
						for c := nCol - 1; c < nCol+2; c++ {
							if row >= 0 && row < len(lines) && c >= 0 && c < len(line) && (c != nCol || row != nLine) && lines[row][c] == '#' {
								nOccupied++
							}
						}
					}
					if col == 'L' && nOccupied == 0 {
						newLines[nLine][nCol] = '#'
						bSomethingChanged = true
					}
					if col == '#' && nOccupied >= 4 {
						newLines[nLine][nCol] = 'L'
						bSomethingChanged = true
					}
				}
				//fmt.Printf(string(col))
			}
			//fmt.Println()
		}
		fmt.Println()

		if !bSomethingChanged {
			nSeats := 0
			for _, line := range newLines {
				for _, col := range line {
					if col == '#' {
						nSeats++
					}
				}
			}
			fmt.Println("Total seats: ", nSeats)
			break
		}

		// Copy back to original
		for n, line := range newLines {
			newLine := make([]rune, len(line))
			copy(newLine, line)
			lines[n] = newLine
		}
	}
}
