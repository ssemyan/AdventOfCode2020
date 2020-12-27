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
	validRules := make(map[int]map[int]bool)
	stage := "rules"
	myTicket := ""
	rulerex := regexp.MustCompile("^([a-zA-Z ]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)$")
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
			badTicket := false
			for _, num := range numsa {
				v, _ := strconv.Atoi(num)

				bBadNum := true
				for _, r := range rules {
					// Does number pass at least one rule?
					if (v >= r.lowMin && v <= r.lowMax) || (v >= r.hiMin && v <= r.hiMax) {
						bBadNum = false
					}
				}
				if bBadNum {
					fmt.Println("Bad num: ", v)
					sumInvalid += v
					badTicket = true
				}
			}

			if !badTicket {

				// Create list of potential rules for each position
				for pos, num := range numsa {

					v, _ := strconv.Atoi(num)

					potentialRules := make(map[int]bool)
					for ruleNum, r := range rules {
						// which rules pass for this position?
						if (v >= r.lowMin && v <= r.lowMax) || (v >= r.hiMin && v <= r.hiMax) {
							potentialRules[ruleNum] = true
						}
					}

					posPotentialRules, exists := validRules[pos]
					if !exists {
						// First time finding potential rules for this pos
						validRules[pos] = potentialRules
					} else {
						// Remove any rules that are not in the join of the two sets
						commonRules := make(map[int]bool)
						for ruleNum := range posPotentialRules {
							_, exists := potentialRules[ruleNum]
							if exists {
								commonRules[ruleNum] = true
							}
						}
						validRules[pos] = commonRules
					}
				}
			}
		}
	}

	// Narrow down potential rules
	for {

		bFoundMulti := false
		for pos, posRules := range validRules {
			if len(posRules) == 1 {
				// since this pos has only one possible rule, remove that rule from all other potentials
				for ruleNum := range posRules {
					for poso, posRuleso := range validRules {
						if pos != poso {
							_, exists := posRuleso[ruleNum]
							if exists {
								delete(posRuleso, ruleNum)
								bFoundMulti = true
							}
						}
					}
				}
			}
		}
		if !bFoundMulti {
			break
		}
	}

	// Print rules for each position
	myTicketNums := strings.Split(myTicket, ",")
	deptVal := 1
	for pos, posRules := range validRules {
		fmt.Printf("Pos: %d Rule: ", pos)
		for ruleNum := range posRules {
			fmt.Printf("%d %s \n", ruleNum, rules[ruleNum].name)
			if strings.HasPrefix(rules[ruleNum].name, "departure") {
				posVal, _ := strconv.Atoi(myTicketNums[pos])
				deptVal *= posVal
			}
		}
	}

	fmt.Println("Sum invalid: ", sumInvalid)

	fmt.Println("Sum my ticket departures: ", deptVal)
}
