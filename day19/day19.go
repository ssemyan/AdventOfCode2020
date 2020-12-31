package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	ruleTxt     string
	isProcessed bool
}

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day19/data.txt")

	// Load rules
	stage := "rules"
	rules := make(map[int]rule)
	validLines := 0
	strRules := ""
	var reRule *regexp.Regexp
	for _, line := range lines {
		if line == "" {
			fmt.Println("Loaded rules: ", len(rules))
			strRules = "^" + getRule(0, rules) + "$"
			fmt.Println("Rule: ", strRules)
			reRule = regexp.MustCompile(strRules)
			stage = "messages"
		} else if stage == "rules" {
			// get rule number
			numEnd := strings.Index(line, ":")
			idx, _ := strconv.Atoi(line[:numEnd])
			ruleTxt := line[numEnd+2:]
			isChar := false
			if ruleTxt[0] == '"' {
				ruleTxt = string(ruleTxt[1])
				isChar = true
			}
			newRule := rule{
				ruleTxt:     ruleTxt,
				isProcessed: isChar,
			}
			rules[idx] = newRule
		} else {
			// Process messages
			fmt.Println("Processing: ", line)

			// Process rules
			if reRule.MatchString(line) {
				fmt.Println("Valid")
				validLines++
			} else {
				fmt.Println("Invalid")
			}
		}
	}
	fmt.Println("Valid: ", validLines)
}

func getRule(ruleNum int, rules map[int]rule) string {

	// Get rule
	rle := rules[ruleNum]
	if rle.isProcessed {
		return rle.ruleTxt
	}

	sides := strings.Split(rle.ruleTxt, "|")
	str := "("
	for i, side := range sides {
		if i > 0 {
			str += "|"
		}
		side1 := strings.Split(side, " ")
		for _, s := range side1 {
			if s != "" {
				sInt, _ := strconv.Atoi(s)
				str += getRule(sInt, rules)
			}
		}
	}
	str += ")"
	rle.isProcessed = true
	rle.ruleTxt = str
	return str
}
