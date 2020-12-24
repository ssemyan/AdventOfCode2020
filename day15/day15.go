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

	// partOne
	//doGame(nums, 2020)
	// partTwo
	doGame(nums, 30000000)
}

func doGame(nums []int, nMaxTurn int) {

	lastSpoken := make(map[int]int)

	// Speak all nums first
	for i, v := range nums {
		// add all but last num
		if i < len(nums)-1 {
			lastSpoken[v] = i
		}
		//fmt.Println(i, v)
	}

	// start with last num
	nCount := len(nums)
	nPrev := nums[nCount-1]

	for {
		// Look for earlier num
		nextNum := 0
		foundNum, exists := lastSpoken[nPrev]
		if exists {
			nextNum = nCount - foundNum - 1
		}
		lastSpoken[nPrev] = nCount - 1
		//fmt.Println(nCount, nextNum)
		nPrev = nextNum
		nCount++

		if nCount == nMaxTurn {
			fmt.Println(nCount, nextNum)
			break
		}
	}

}
