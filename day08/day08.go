package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Read all data into memory
	lines := fileload.Fileload("day08/data.txt")

	accumulator := 0

	// keep track of where we have been
	ran := make([]bool, len(lines))

	// run program
	curIns := 0
	for {
		curLine := lines[curIns]
		if ran[curIns] {
			break
		}
		ran[curIns] = true

		fmt.Println(curLine)
		parts := strings.Split(curLine, " ")
		n, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "acc":
			accumulator += n
			curIns++
		case "jmp":
			curIns += n
		case "nop":
			curIns++
		default:
			panic(fmt.Sprintf("Unknown command %s", parts[0]))
		}
	}

	fmt.Println("Acc ", accumulator)
}
