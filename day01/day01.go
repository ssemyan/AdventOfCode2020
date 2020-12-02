package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Read all data into memory
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//fmt.Println("Contents of file:", string(data))

	// Convert to array of strings
	lines := strings.Split(string(data), "\n")

	// convert to array of ints (make sure no empty lines)
	nums := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error in file: ", line)
		}
		nums[i] = n
	}

	fmt.Println("Lines:", len(nums))

	FindSum(nums, 2020, true)
	FindSum(nums, 2020, false)
}

// FindSum - find two or three numbers that add up to a given value
func FindSum(nums []int, sumToFind int, findTwo bool) {

	numLen := len(nums)

	// Start loop
	for i, val1 := range nums {
		for v := i + 1; v < numLen; v++ {
			val2 := nums[v]
			if findTwo {
				// check for two nums that add to sum
				if (i != v) && ((val1 + val2) == sumToFind) {
					fmt.Println("Found values: ", val1, val2, val1*val2)
					return
				}
			} else {
				// check for three nums that add to sum
				for z := v + 1; z < numLen; z++ {
					val3 := nums[z]
					if (i != v) && (i != z) && ((val1 + val2 + val3) == sumToFind) {
						fmt.Println("Found values: ", val1, val2, val3, val1*val2*val3)
						return
					}
				}
			}
		}
	}
	fmt.Println("Not found.")
}
