package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strings"
)

func main() {

	// Read all data into memory
	line := fileload.ReadFile("day04/data.txt")

	// Split by empty line
	recs := strings.Split(line, "\n\n")
	fmt.Printf("Found %d recs\n", len(recs))

	// Turn each record into map
	//passports := make([]map[string]string, len(recs))
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
		//passports[i] = passport
	}

	fmt.Println("Valid passports: ", validPass)

}

// IsValidPassport - Tell whether a passport is valid
func IsValidPassport(passport map[string]string) bool {

	// Must have all but cid
	// byr (Birth Year)
	// iyr (Issue Year)
	// eyr (Expiration Year)
	// hgt (Height)
	// hcl (Hair Color)
	// ecl (Eye Color)
	// pid (Passport ID)
	// cid (Country ID)

	fields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range fields {
		_, exists := passport[field]
		if !exists {
			fmt.Println("Missing field: ", field)
			return false
		}
	}
	return true
}
