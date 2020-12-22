package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day14/data.txt")

	//partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {

	mask := ""
	andMask := 0
	var orMask uint64
	re := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")
	memory := make(map[int]int)
	for _, line := range lines {

		// Get mask
		if strings.HasPrefix(line, "mask = ") {
			mask = line[7:]
			fmt.Println("Found mask ", mask)

			// Make the AND and OR masks
			andMask = 0
			orMask = 0
			for i := len(mask) - 1; i >= 0; i-- {
				if mask[i] == '1' {
					// set the bit in the AND mask
					andMask |= (1 << (len(mask) - i - 1))
				} else if mask[i] == '0' {
					orMask |= (1 << (len(mask) - i - 1))
				}
			}
			// finally flip the or mask
			orMask = ^orMask

			fmt.Println("AND mask: ", strconv.FormatInt(int64(andMask), 2))
			fmt.Println("OR mask: ", strconv.FormatInt(int64(orMask), 2))
		} else {

			// Write to mem
			parts := re.FindStringSubmatch(line)
			memLoc, _ := strconv.Atoi(parts[1])
			memVal, _ := strconv.Atoi(parts[2])
			fmt.Println("Found mem at ", memLoc, memVal, strconv.FormatInt(int64(memVal), 2))

			newVal := memVal | andMask
			fmt.Println("After AND mask: ", strconv.FormatInt(int64(newVal), 2))
			newVal = int(uint64(newVal) & orMask)
			fmt.Println("After OR mask: ", strconv.FormatInt(int64(newVal), 2))
			memory[memLoc] = newVal
		}

		nTotal := 0
		for _, val := range memory {
			nTotal += val
		}
		fmt.Println("Total: ", nTotal)
	}

}

func partTwo(lines []string) {

	mask := ""
	orMask := 0
	re := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")
	memory := make(map[int]int)
	for _, line := range lines {

		// Get mask
		if strings.HasPrefix(line, "mask = ") {
			mask = line[7:]
			//fmt.Println("Found mask ", mask)

			// Make the OR mask
			orMask = 0
			for i := len(mask) - 1; i >= 0; i-- {
				if mask[i] == '1' {
					// set the bit in the OR mask
					orMask |= (1 << (len(mask) - i - 1))
				}
			}
			//fmt.Println("OR mask: ", strconv.FormatInt(int64(orMask), 2))
		} else {

			// Write to mem
			parts := re.FindStringSubmatch(line)
			memLoc, _ := strconv.Atoi(parts[1])
			memVal, _ := strconv.Atoi(parts[2])
			//fmt.Println("Found loc at ", memLoc, memVal, strconv.FormatInt(int64(memLoc), 2))

			newLoc := memLoc | orMask
			//fmt.Println("After OR mask: ", strconv.FormatInt(int64(newLoc), 2))

			// Write to all memory locations
			writeMem(newLoc, []byte(mask), memVal, memory, 0)
		}
	}

	nTotal := 0
	for _, val := range memory {
		nTotal += val
	}
	fmt.Println("Total: ", nTotal)

}

func writeMem(newLoc int, mask []byte, memVal int, memory map[int]int, nStart int) {

	branched := false
	for i := nStart; i < len(mask); i++ {
		if mask[i] == 'X' {
			// Create two new masks with each value set
			newMaskZero := make([]byte, len(mask))
			newMaskOne := make([]byte, len(mask))
			copy(newMaskOne, mask)
			copy(newMaskZero, mask)
			newMaskZero[i] = '1'
			newMaskOne[i] = '0'
			writeMem(newLoc, newMaskZero, memVal, memory, i+1)
			writeMem(newLoc, newMaskOne, memVal, memory, i+1)
			branched = true
			break
		} else {
			// Set to Y to ignore
			mask[i] = 'Y'
		}
	}

	if !branched {
		// Reached the end of a variation of mem
		//fmt.Println("Mask var: ", string(mask))

		// Make the OR and Not OR masks
		orMask := 0
		var notOrMask uint64
		notOrMask = 0
		for i := len(mask) - 1; i >= 0; i-- {
			if mask[i] == '1' {
				// set the bit in the OR mask
				orMask |= (1 << (len(mask) - i - 1))
			} else if mask[i] == '0' {
				notOrMask |= (1 << (len(mask) - i - 1))
			}
		}
		// finally flip the or mask
		notOrMask = ^notOrMask
		//fmt.Println("Mem loc: ", strconv.FormatInt(int64(newLoc), 2))
		newVal := newLoc | orMask
		//fmt.Println("After OR mask: ", strconv.FormatInt(int64(newVal), 2))
		newVal = int(uint64(newVal) & notOrMask)
		//fmt.Println("After NOR mask: ", strconv.FormatInt(int64(newVal), 2))
		//fmt.Println("Writing to: ", newVal)
		memory[newVal] = memVal
	}
}
