package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	OccurLow  int
	OccurHigh int
	Letter    byte
	Pass      string
}

func main() {

	// Read all data into memory
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))

	// Convert to array of strings
	lines := strings.Split(string(data), "\n")

	// convert to array of passwords (make sure no empty lines)
	passwords := make([]Password, len(lines))

	// regex to struct example:
	// 1-3 a: abcde
	re := regexp.MustCompile("([0-9]+)-([0-9]+)\\s(\\w):\\s(\\w+)")

	for i, line := range lines {

		// 0 element will be full string
		parts := re.FindStringSubmatch(line)

		lowVal, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error in file: ", parts[1])
		}

		highVal, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Error in file: ", parts[2])
		}

		pass := Password{
			OccurLow:  lowVal,
			OccurHigh: highVal,
			Letter:    parts[3][0],
			Pass:      parts[4],
		}
		passwords[i] = pass
	}

	fmt.Printf("Loaded %d passwords\n", len(passwords))

	nValid, nInvalid := 0, 0

	for _, pass := range passwords {

		// need to validate char occurs between max and min num of times
		nCount := 0
		for z := 0; z < len(pass.Pass); z++ {
			if pass.Letter == pass.Pass[z] {
				nCount++
			}
		}
		if (nCount >= pass.OccurLow) && (nCount <= pass.OccurHigh) {
			nValid++
		} else {
			nInvalid++
		}
	}

	fmt.Printf("Done. Valid: %d InValid: %d", nValid, nInvalid)
}
