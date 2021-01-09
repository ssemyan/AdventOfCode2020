package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strings"
)

type food struct {
	lineNum int
	ingred  map[string]bool
	allerg  map[string]bool
}

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day21/testdata.txt")

	// Load tiles into array
	foods := []food{}
	for n, line := range lines {

		splitPt := strings.Index(line, " (contains ")
		foodsStr := line[:splitPt]
		foodArr := strings.Split(foodsStr, " ")
		foodMap := make(map[string]bool)
		for _, fd := range foodArr {
			foodMap[fd] = true
		}

		alergStr := line[splitPt+11 : len(line)-1]
		alergArray := strings.Split(alergStr, ", ")
		alergMap := make(map[string]bool)
		for _, alg := range alergArray {
			alergMap[alg] = true
		}

		newFood := food{
			lineNum: n,
			ingred:  foodMap,
			allerg:  alergMap,
		}
		foods = append(foods, newFood)
	}

	fmt.Println("Foods loaded: ", len(foods))

	// Make list of ingred that also appear in other lines with same allergens
	// kfcds, nhms, sbzzf, trh cannot contain an allergen.
	knownFoods := make(map[string]bool)
	allFoods := []string{}
	knownAlergens := make(map[string]string)
	for _, foodList := range foods {

		for ing := range foodList.ingred {

			// Add to list of known foods - just overwrite if existing
			allFoods = append(allFoods, ing)

			// Look through the other foods for ingred that have same allerg
			bIsInOther := false
			bFoundSameAleg := false
			alergn := ""
			for _, foodList2 := range foods {
				if foodList.lineNum != foodList2.lineNum {
					bSameAllerg := false
					currAlg := ""
					for alg := range foodList.allerg {
						if exists(alg, foodList2.allerg) {
							bSameAllerg = true
							bFoundSameAleg = true
							currAlg = alg
							break
						}
					}

					if bSameAllerg {
						if exists(ing, foodList2.ingred) {
							bIsInOther = true
							if alergn == "" {
								alergn = currAlg
							}
							if alergn != currAlg {
								alergn += "_" + currAlg
							}
						} else {
							bIsInOther = false
							break
						}
					}
				}
			}
			if bIsInOther || (!bFoundSameAleg && !bIsInOther) {
				knownFoods[ing] = true
				if bIsInOther {
					fmt.Println("Known: ", ing, alergn)
					knownAlergens[ing] = alergn
				}
			}
		}
	}
	fmt.Println("Known foods: ", len(knownFoods))
	for uf := range knownFoods {
		fmt.Printf("%s, ", uf)
	}
	finalFoods := []string{}
	for _, food := range allFoods {
		if !exists(food, knownFoods) {
			finalFoods = append(finalFoods, food)
		}
	}
	fmt.Println()
	fmt.Println("Final foods: ", len(finalFoods))
	for _, uf := range finalFoods {
		fmt.Printf("%s, ", uf)
	}
}

func exists(key string, mp map[string]bool) bool {
	_, exists := mp[key]
	return exists
}
