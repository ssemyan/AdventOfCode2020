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
	lines := fileload.Fileload("day21/data.txt")

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

	// Make list of ingred that do not appear in other lines with same allergens
	// kfcds, nhms, sbzzf, trh cannot contain an allergen.
	unknownFoods := []string{}
	soloFoods := make(map[string]bool)
	for _, foodList := range foods {

		// Look up ingredients in other foods
		for ing := range foodList.ingred {

			// Look through the other foods for ingred that have same allerg
			bIsInOther := false
			bFoundSameAleg := false
			for _, foodList2 := range foods {
				if foodList.lineNum != foodList2.lineNum {
					bSameAllerg := false
					for alg := range foodList.allerg {
						if exists(alg, foodList2.allerg) {
							bSameAllerg = true
							bFoundSameAleg = true
							break
						}
					}

					if bSameAllerg {
						if exists(ing, foodList2.ingred) {
							bIsInOther = true
							break
						}
					}
				}
			}
			if !bIsInOther && bFoundSameAleg {
				unknownFoods = append(unknownFoods, ing)
			} else if !bIsInOther && !bFoundSameAleg {
				// If we find something again that is not in any other ingred, and alerg is unique, then it is valid
				soloFoods[ing] = true
			}
		}
	}
	fmt.Println("Unknown foods: ", len(unknownFoods))
	for _, uf := range unknownFoods {
		fmt.Printf("%s, ", uf)
	}
	fmt.Println()
	fmt.Println("Solo foods: ", len(soloFoods))
	for uf := range soloFoods {
		fmt.Printf("%s, ", uf)
	}

	finalFoods := []string{}
	for _, food := range unknownFoods {
		if !exists(food, soloFoods) {
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
