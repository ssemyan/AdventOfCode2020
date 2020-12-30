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

	// Part Two
	nSum2 := 0
	for _, line := range lines {
		fmt.Println("Line: ", line)
		lineWPre := addPrecedence(line)
		fmt.Println("Line w precedence: ", lineWPre)
		nSum2 += getVal(reverse(lineWPre))
	}
	fmt.Println("Part Two Sum: ", nSum2)

}

func addPrecedence(line string) string {

	// Use () to add precedence for + over *
	// not the prettiest code but...
	precChar := "+"
	nCurrLoc := 0
	for {
		fmt.Println("Line: ", line)

		nextAdd := strings.Index(line[nCurrLoc:], precChar)
		if nextAdd == -1 {
			return line
		}
		nextAdd += nCurrLoc
		// Add parens before
		if line[nextAdd-2] == ')' {
			// walk back to find matching paren
			rCount := 1
			for i := nextAdd - 3; i >= 0; i-- {
				if line[i] == ')' {
					rCount++
				} else if line[i] == '(' {
					rCount--
				}
				if rCount == 0 {
					line = line[:i] + "(" + line[i:]
					break
				}
			}
			if rCount != 0 {
				panic("Did not find end of parens")
			}
			//line = line[:nextAdd-2] + ")" + line[nextAdd-2:]
		} else {
			// walk back to find end of num
			parenPos := 0
			for i := nextAdd - 3; i >= 0; i-- {
				if line[i] == ' ' || line[i] == '(' {
					parenPos = i
					if line[i] == ' ' {
						parenPos++
					}
					break
				}
			}
			if parenPos == 0 {
				line = "(" + line
			} else {
				line = line[:parenPos] + "(" + line[parenPos:]
			}
		}

		// Add parens after
		if line[nextAdd+3] == '(' {
			// walk forward to find matching paren
			rCount := 1
			for i := nextAdd + 4; i < len(line); i++ {
				if line[i] == '(' {
					rCount++
				} else if line[i] == ')' {
					rCount--
				}
				if rCount == 0 {
					line = line[:i] + ")" + line[i:]
					break
				}
			}
			if rCount != 0 {
				panic("Did not find end of parens")
			}
			//line = line[:nextAdd+3] + "(" + line[nextAdd+3:]
		} else {
			// walk forward to find end of num
			parenPos := 0
			for i := nextAdd + 3; i < len(line); i++ {
				if line[i] == ' ' || line[i] == ')' {
					parenPos = i
					break
				}
			}
			if parenPos == 0 {
				line = line + ")"
			} else {
				line = line[:parenPos] + ")" + line[parenPos:]
			}

		}

		nCurrLoc = nextAdd + 2
	}

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
