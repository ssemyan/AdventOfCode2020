package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"math"
)

func main() {

	// Read all data into memory
	lines := fileload.Fileload("day05/data.txt")

	maxSeatCode := 0.0

	for _, line := range lines {

		rowMin, rowMax := 0.0, 127.0
		colMin, colMax := 0.0, 7.0

		for _, chr := range line {

			switch chr {
			case 'F':
				rowMax -= math.Round((rowMax - rowMin) / 2)

			case 'B':
				rowMin += math.Round((rowMax - rowMin) / 2)

			case 'L':
				colMax -= math.Round((colMax - colMin) / 2)

			case 'R':
				colMin += math.Round((colMax - colMin) / 2)
			}

		}
		// Verify
		if (rowMin != rowMax) || (colMin != colMax) {
			fmt.Println("Bad result: ", line, rowMin, rowMax, colMin, colMax)
			panic("Wrong data")
		}

		seatCode := rowMin*8 + colMin
		if seatCode > maxSeatCode {
			maxSeatCode = seatCode
		}
		fmt.Println("Result: ", line, rowMin, rowMax, colMin, colMax, seatCode)
	}
	fmt.Println("Max seat code: ", maxSeatCode)
}
