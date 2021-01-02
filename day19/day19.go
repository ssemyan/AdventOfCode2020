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

	// Part one without special rules
	valid := doRules(lines, false)
	fmt.Println("Part One Valid: ", valid)

	valid = doRules(lines, true)
	fmt.Println("Part Two Valid: ", valid)

}

func doRules(lines []string, doSpecialRules bool) int {

	// Load rules
	stage := "rules"
	rules := make(map[int]rule)
	validLines := 0
	strRules := ""
	var reRule, rule42, rule31 *regexp.Regexp
	for _, line := range lines {
		if line == "" {
			//fmt.Println("Loaded rules: ", len(rules))
			strRules = "^" + getRule(0, rules, doSpecialRules) + "$"
			//fmt.Println("Final Rule: ", strRules)
			reRule = regexp.MustCompile(strRules)

			// Load special rules for part two
			rule42txt := "(" + rules[42].ruleTxt + ")"
			//fmt.Println("Rule 42: ", rule42txt)
			rule31txt := "(" + rules[31].ruleTxt + ")"
			//fmt.Println("Rule 31: ", rule31txt)
			rule42 = regexp.MustCompile(rule42txt)
			rule31 = regexp.MustCompile(rule31txt)

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
			//fmt.Println("Processing: ", line)

			// Process rules
			if reRule.MatchString(line) {
				// Need to validate rule 11 - captured groups must have correct count
				matches := reRule.FindStringSubmatch(line)
				matchNames := reRule.SubexpNames()
				rule31Cnt, rule42Cnt := 0, 0
				for i, match := range matches[1:] {
					//fmt.Println(matchNames[i+1], match)
					if matchNames[i+1] == "rule11_42" || matchNames[i+1] == "rule8_42" {
						rule42Matches := rule42.FindAllString(match, -1)
						rule42Cnt += len(rule42Matches)
					} else if matchNames[i+1] == "rule11_31" {
						rule31Matches := rule31.FindAllString(match, -1)
						rule31Cnt = len(rule31Matches)
					}
				}
				if (rule31Cnt < rule42Cnt) || !doSpecialRules {
					//fmt.Println("Valid")
					validLines++
				} else {
					//fmt.Println("Different match nums")
				}
			} else {
				//fmt.Println("Invalid")
			}
		}
	}
	//fmt.Println("Valid: ", validLines)
	return validLines
}

func getRule(ruleNum int, rules map[int]rule, doSpecialRules bool) string {

	// Get rule
	rle := rules[ruleNum]
	if rle.isProcessed {
		return rle.ruleTxt
	}

	sides := strings.Split(rle.ruleTxt, "|")

	// Make all non-capturing
	str := "(?:"

	for i, side := range sides {
		if i > 0 {
			str += "|"
		}
		side1 := strings.Split(side, " ")
		for _, s := range side1 {
			if s != "" {
				sInt, _ := strconv.Atoi(s)
				strTemp := getRule(sInt, rules, doSpecialRules)

				// Handle loop in rules 8 and 11
				if (ruleNum == 11 || ruleNum == 8) && doSpecialRules {
					str += "(?P<rule" + fmt.Sprint(ruleNum) + "_" + s + ">(?:" + strTemp + ")+)"
				} else {
					str += strTemp
				}
			}
		}
	}
	str += ")"

	//fmt.Println("Rule ", ruleNum, str)
	newRule := rule{
		isProcessed: true,
		ruleTxt:     str,
	}
	rules[ruleNum] = newRule
	return str
}
