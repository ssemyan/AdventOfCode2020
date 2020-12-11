package fileload

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Fileload - load file into array of strings
func Fileload(filename string) []string {

	// Read all data into memory
	data := ReadFile(filename)

	// Convert to array of strings
	lines := strings.Split(data, "\n")
	fmt.Printf("Loaded %d lines\n", len(lines))

	return lines
}

// FileLoadInts - load file as array of ints
func FileLoadInts(filename string) []int {
	lines := Fileload(filename)
	nums := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error in file: ", line)
		}
		nums[i] = n
	}
	return nums
}

// ReadFile - load a file into a string
func ReadFile(filename string) string {

	// Read all data into memory
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		panic("Ending")
	}

	return string(data)
}
