package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// Read all data into memory
	line := fileload.ReadFile("day04/data.txt")

	// Split by empty line
	recs := strings.Split(line, "\n\n")
	fmt.Printf("Found %d recs\n", len(recs))

	// Turn each record into map
	validPass := 0
	for _, rec := range recs {
		recData := strings.Split(strings.ReplaceAll(rec, "\n", " "), " ")
		passport := make(map[string]string)
		for _, data := range recData {
			dataParts := strings.Split(data, ":")
			passport[dataParts[0]] = dataParts[1]
		}
		if IsValidPassport(passport) {
			validPass++
		}
	}

	fmt.Println("Valid passports: ", validPass)

}

// IsValidPassport - Tell whether a passport is valid
func IsValidPassport(passport map[string]string) bool {

	// Must have all but cid
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.

	fields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	hgtTest := regexp.MustCompile("^(\\d+)([ic][nm])$")
	hclTest := regexp.MustCompile("^#[0-9a-f]{6}$")
	eclTest := regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
	pidTest := regexp.MustCompile("^[0-9]{9}$")

	for _, field := range fields {
		value, exists := passport[field]
		if !exists {
			//fmt.Println("Missing field: ", field)
			return false
		}

		// validate fields
		passportGood := true
		switch field {
		case "byr":
			n, err := strconv.Atoi(value)
			if err != nil || (n < 1920 || n > 2002) {
				passportGood = false
			}

		case "iyr":
			n, err := strconv.Atoi(value)
			if err != nil || (n < 2010 || n > 2020) {
				passportGood = false
			}

		case "eyr":
			n, err := strconv.Atoi(value)
			if err != nil || (n < 2020 || n > 2030) {
				passportGood = false
			}

		case "hgt":
			parts := hgtTest.FindStringSubmatch(value)
			passportGood = (len(parts) == 3)
			if passportGood {
				n, err := strconv.Atoi(parts[1])
				if err != nil || (parts[2] == "in" && (n < 59 || n > 76)) || (parts[2] == "cm" && (n < 150 || n > 193)) {
					passportGood = false
				}
			}

		case "hcl":
			passportGood = hclTest.MatchString(value)

		case "ecl":
			passportGood = eclTest.MatchString(value)

		case "pid":
			passportGood = pidTest.MatchString(value)
		}

		if !passportGood {
			fmt.Printf("%s invalid: %s\n", field, value)
			return false
		}
	}

	return true
}
