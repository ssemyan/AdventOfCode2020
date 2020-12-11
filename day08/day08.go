package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
	"strings"
)

type RunResult struct {
	accumulator int
	infinite    bool
}

func main() {

	// Read all data into memory
	lines := fileload.Fileload("day08/data.txt")

	// Part One
	result := testCode(lines, -1)
	fmt.Printf("Result acc: %v inf: %v \n", result.accumulator, result.infinite)

	// Part Two
	for i := range lines {
		result = testCode(lines, i)
		if !result.infinite {
			fmt.Printf("Found wrong line %v add %v", i, result.accumulator)
			break
		}
	}
}

func testCode(lines []string, rowToTry int) RunResult {

	accumulator := 0
	infinite := false

	// keep track of where we have been
	ran := make([]bool, len(lines))

	// run program
	curIns := 0
	for {
		if curIns == len(lines) {
			// program successfully ran
			break
		}

		curLine := lines[curIns]
		if ran[curIns] {
			// infinite loop detected
			infinite = true
			break
		}
		ran[curIns] = true

		fmt.Println(curLine)
		parts := strings.Split(curLine, " ")
		n, _ := strconv.Atoi(parts[1])

		// Swap instruction if indicated
		inst := parts[0]
		if curIns == rowToTry {
			if inst == "jmp" {
				inst = "nop"
			} else if inst == "nop" {
				inst = "jmp"
			}
		}

		switch inst {
		case "acc":
			accumulator += n
			curIns++
		case "jmp":
			curIns += n
		case "nop":
			curIns++
		default:
			panic(fmt.Sprintf("Unknown command %s", inst))
		}
	}

	return RunResult{
		accumulator: accumulator,
		infinite:    infinite,
	}
}
