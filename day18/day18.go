package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day18/data.txt")

	// Part One
	nSum := 0
	for _, line := range lines {
		fmt.Println("Line: ", line)
		nSum += getVal(reverse(line))
	}
	fmt.Println("Part One Sum: ", nSum)
}

func getVal(equation string) int {

	num1, currPos := 0, 0

	fmt.Println("Solving: ", equation)

	// equation starts with num or (
	if equation[0] == '(' {
		// Find closing )
		openCnt := 0
		for pos, char := range equation {
			if char == '(' {
				openCnt++
			} else if char == ')' {
				openCnt--
			}
			if openCnt == 0 {
				num1 = getVal(equation[1:pos])
				currPos = pos + 1
				if currPos == len(equation) {
					return num1
				}
				break
			}
		}
	} else {
		currPos = strings.Index(equation, " ")
		if currPos == -1 {
			// At end of string
			num1, _ = strconv.Atoi(equation)
			fmt.Println("Returning num: ", num1)
			return num1
		}
		num1, _ = strconv.Atoi(equation[0:currPos])
	}
	//fmt.Println("Num 1: ", num1)
	num2 := 0
	switch equation[currPos+1] {
	case '+':
		num2 = num1 + getVal(equation[currPos+3:])
	case '*':
		num2 = num1 * getVal(equation[currPos+3:])
	default:
		panic(fmt.Sprint("Unknown char at pos : ", equation[currPos+1]))
	}
	fmt.Println("Returning: ", num2)
	return num2
}

func reverse(s string) (result string) {
	for _, v := range s {
		if v == ')' {
			v = '('
		} else if v == '(' {
			v = ')'
		}
		result = string(v) + result
	}
	return
}
