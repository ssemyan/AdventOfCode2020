package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
)

func main() {

	// Read all data into memory
	lines := fileload.Fileload("day03/data.txt")

	// put trees in a map
	trees := make(map[string]bool)
	numCols, numRows := len(lines[0]), len(lines)
	for row, line := range lines {
		for col, square := range line {
			if square == '#' {
				key := strconv.Itoa(row) + "_" + strconv.Itoa(col)
				fmt.Println("Found tree: ", key)
				trees[key] = true
			}
		}
	}

	// Now traverse down
	row, col, rowMv, colMv, treesHit := 0, 0, 1, 3, 0
	for {
		row += rowMv
		col += colMv
		if col >= numCols {
			col = col - numCols
		}
		if row > numRows {
			break
		}
		key := strconv.Itoa(row) + "_" + strconv.Itoa(col)
		fmt.Println("Checking ", key)
		_, exists := trees[key]
		if exists {
			treesHit++
			fmt.Println("Hit tree")
		}
	}
	fmt.Printf("Hit %d trees\n", treesHit)
}
