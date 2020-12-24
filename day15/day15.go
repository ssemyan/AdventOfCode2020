package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day15/data.txt")

	nlines := strings.Split(lines[0], ",")
	nums := make([]int, len(nlines))
	for i, line := range nlines {
		v, _ := strconv.Atoi(line)
		nums[i] = v
	}

	partOne(nums)
	//partTwo(lines)
}

func partOne(nums []int) {

	nMaxTurn := 2020
	spoken := make([]int, nMaxTurn)

	// Speak all nums first
	for i, v := range nums {
		spoken[i] = v
		//fmt.Println(i, v)
	}

	// start with next num
	nCount := len(nums)

	for {
		nPrev := spoken[nCount-1]
		// Look for earlier num
		nextNum := 0
		for i := nCount - 2; i >= 0; i-- {
			if spoken[i] == nPrev {
				nextNum = nCount - i - 1
				break
			}
		}
		spoken[nCount] = nextNum
		nCount++

		if nCount == nMaxTurn {
			fmt.Println(nCount, nextNum)
			break
		}
	}

}
