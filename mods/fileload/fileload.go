package fileload

import (
	"fmt"
	"io/ioutil"
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
