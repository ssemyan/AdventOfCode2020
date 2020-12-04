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
				//fmt.Println("Found tree: ", key)
				trees[key] = true
			}
		}
	}

	// Right 1, down 1.
	// Right 3, down 1. (This is the slope you already checked.)
	// Right 5, down 1.
	// Right 7, down 1.
	// Right 1, down 2.

	treesHit := findTrees(1, 1, numCols, numRows, trees)
	treesHit *= findTrees(3, 1, numCols, numRows, trees)
	treesHit *= findTrees(5, 1, numCols, numRows, trees)
	treesHit *= findTrees(7, 1, numCols, numRows, trees)
	treesHit *= findTrees(1, 2, numCols, numRows, trees)

	fmt.Printf("Hit %d total trees\n", treesHit)
}

func findTrees(colMv int, rowMv int, numCols int, numRows int, trees map[string]bool) int {

	row, col, treesHit := 0, 0, 0
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
		//fmt.Println("Checking ", key)
		_, exists := trees[key]
		if exists {
			treesHit++
			//fmt.Println("Hit tree")
		}
	}
	fmt.Printf("Hit %d trees\n", treesHit)

	return treesHit
}
