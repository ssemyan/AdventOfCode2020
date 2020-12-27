package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	name   string
	lowMin int
	lowMax int
	hiMin  int
	hiMax  int
}

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day16/data.txt")

	sumInvalid := 0

	rules := make(map[int]rule)
	stage := "rules"
	myTicket := ""
	rulerex := regexp.MustCompile("^([a-zA-Z ])+: (\\d+)-(\\d+) or (\\d+)-(\\d+)$")
	for _, line := range lines {

		if stage == "rules" {
			if line != "" {
				if line == "your ticket:" {
					fmt.Println("Loaded rules: ", len(rules))
					stage = "yourticket"
				} else {
					parts := rulerex.FindStringSubmatch(line)
					lowMin, _ := strconv.Atoi(parts[2])
					lowMax, _ := strconv.Atoi(parts[3])
					hiMin, _ := strconv.Atoi(parts[4])
					hiMax, _ := strconv.Atoi(parts[5])
					newRule := rule{
						name:   parts[1],
						lowMin: lowMin,
						lowMax: lowMax,
						hiMin:  hiMin,
						hiMax:  hiMax,
					}
					rules[len(rules)] = newRule
				}
			}
		} else if stage == "yourticket" {
			if line != "" {
				if line == "nearby tickets:" {
					fmt.Println("Loaded ticket: ", myTicket)
					stage = "nearbytickets"
				} else {
					myTicket = line
				}
			}
		} else {

			// Test each ticket
			// Get numbers from each line
			numsa := strings.Split(line, ",")
			for _, num := range numsa {
				v, _ := strconv.Atoi(num)

				bBadNum := true
				for _, r := range rules {
					if (v >= r.lowMin && v <= r.lowMax) || (v >= r.hiMin && v <= r.hiMax) {
						bBadNum = false
					}
				}
				if bBadNum {
					fmt.Println("Bad num: ", v)
					sumInvalid += v
				}
			}
		}
	}

	fmt.Println("Sum invalid: ", sumInvalid)
}
