package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strings"
)

func main() {

	// Read all data into memory
	line := fileload.ReadFile("day06/data.txt")

	// Split by empty line
	recs := strings.Split(line, "\n\n")
	fmt.Printf("Found %d recs\n", len(recs))

	// Part one Count unique answers
	totUnique := 0
	for _, rec := range recs {
		answers := make(map[rune]bool)
		for _, ans := range rec {
			if ans != '\n' {
				_, exists := answers[ans]
				if !exists {
					answers[ans] = true
				}
			}
		}
		fmt.Println("Unique answers: ", len(answers))
		totUnique += len(answers)
	}
	fmt.Println("Total of answers: ", totUnique)

	// Part two Count answers all said yes
	totAllYes := 0
	for _, rec := range recs {
		answers := make(map[rune]int)
		numPass := 1
		for _, ans := range rec {
			if ans != '\n' {
				curVal, exists := answers[ans]
				if !exists {
					answers[ans] = 1
				} else {
					answers[ans] = curVal + 1
				}
			} else {
				numPass++
			}
		}

		// Find numb of answers all said yes to
		groupYes := 0
		for _, passCnt := range answers {
			if passCnt == numPass {
				groupYes++
			}
		}
		fmt.Println("Answers all said yes: ", groupYes)
		totAllYes += groupYes
	}
	fmt.Println("Total of yes answers: ", totAllYes)

}
