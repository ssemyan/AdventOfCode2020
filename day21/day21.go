package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"sort"
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
			if exists(fd, foodMap) {
				fmt.Println("Dupe: ", fd)
			}
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

	// Make list of ingred alergns
	ingredAlergn := make(map[string]string)
	for {
		// Make list of ingred that also appear in every other lines with same allergens
		knownAlergens := make(map[string]map[string]bool)
		for _, foodList := range foods {

			// loop through aleg
			for alg := range foodList.allerg {

				// Find ingred that are in every instance of the aleg
				for ing := range foodList.ingred {

					bIsInEveryOther := true
					for _, foodList2 := range foods {
						if foodList.lineNum != foodList2.lineNum {
							bSameAllerg := false
							for alg := range foodList.allerg {
								if exists(alg, foodList2.allerg) {
									bSameAllerg = true
									break
								}
							}

							if bSameAllerg {
								if !exists(ing, foodList2.ingred) {
									bIsInEveryOther = false
									break
								}
							}
						}
					}
					if bIsInEveryOther {
						//fmt.Println("Known: ", ing, alg)
						possIng, exis := knownAlergens[alg]
						if !exis {
							possIng = make(map[string]bool)
						}
						possIng[ing] = true
						knownAlergens[alg] = possIng
					}
				}
			}
		}

		// if no alergens found, we are done
		if len(knownAlergens) == 0 {
			break
		}

		// Find any alg with only one ing, remove that ingred and aleg from the other foods and repeat
		for alg, ingds := range knownAlergens {
			if len(ingds) == 1 {
				ing := getFirstKey(ingds)
				fmt.Println("Found aleg match: ", alg, ing)
				ingredAlergn[alg] = ing
				// remove this ing, and aleg from all foods
				for _, foodList := range foods {
					if exists(ing, foodList.ingred) { // need if?
						delete(foodList.ingred, ing)
					}
					if exists(alg, foodList.allerg) { // need if?
						delete(foodList.allerg, alg)
					}
				}
				break
			}
		}
	}

	// all the ingred left are inert
	unknownFoods := []string{}
	for _, foodList := range foods {
		for ing := range foodList.ingred {
			unknownFoods = append(unknownFoods, ing)
		}
	}

	// Part One
	fmt.Println()
	fmt.Println("Unknown foods: ", len(unknownFoods))

	// Part two
	// create alpha list of alergns
	alerg := []string{}
	for alg := range ingredAlergn {
		alerg = append(alerg, alg)
	}

	// sort list
	sort.Strings(alerg)

	// print list
	fmt.Println("canonical dangerous ingredient list: ")
	for _, alg := range alerg {
		fmt.Printf("%s,", ingredAlergn[alg])
	}
	fmt.Println()
}

func exists(key string, mp map[string]bool) bool {
	_, exists := mp[key]
	return exists
}

func getFirstKey(mp map[string]bool) string {
	retKey := ""
	for key := range mp {
		retKey = key
		break
	}
	return retKey
}
