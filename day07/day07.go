package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BagChild struct {
	num   int
	color string
}

func main() {

	// Read all data into memory
	lines := fileload.Fileload("day07/data.txt")

	// create rule lookup
	parentBags := make(map[string][]BagChild)
	parentLookup := make(map[string]map[string]bool)

	// Regex for bag parts
	breg := regexp.MustCompile("([0-9]+) ([a-z ]+) bag[s\\.]*")
	for _, rule := range lines {
		parts := strings.Split(rule, " bags contain ")

		// save the parent lookup
		if parts[1] != "no other bags." {
			bagParts := strings.Split(parts[1], ",")
			childList := make([]BagChild, len(bagParts))
			for i, bagPart := range bagParts {
				bagRule := breg.FindStringSubmatch(bagPart)
				fmt.Printf("%s contains %s %s\n", parts[0], bagRule[1], bagRule[2])

				// Add child to list
				n, _ := strconv.Atoi(bagRule[1])
				newChild := BagChild{
					num:   n,
					color: bagRule[2],
				}
				childList[i] = newChild

				// Add to parent map
				parentMap, exists := parentLookup[bagRule[2]]
				if !exists {
					parentMap = make(map[string]bool)
					parentLookup[bagRule[2]] = parentMap
				}
				parentMap[parts[0]] = true
			}

			// save the child bags
			parentBags[parts[0]] = childList
		}
	}
	fmt.Println("Rules loaded: ", len(parentBags))

	// Print the lookups
	for bag, parents := range parentLookup {
		fmt.Printf("%s has parents: ", bag)
		for pmap := range parents {
			fmt.Printf(" %s, ", pmap)
		}
		fmt.Println()
	}

	// Part one: Find the unique parents
	uniqueParents := make(map[string]bool)
	bagToFind := "shiny gold"
	findParents(bagToFind, parentLookup, uniqueParents)

	fmt.Printf("Unique parents for %s: ", bagToFind)
	for par := range uniqueParents {
		fmt.Printf("%s, ", par)
	}
	fmt.Println()
	fmt.Println("Total unique parents: ", len(uniqueParents))

	// Part two: count child bags
	bagCount := countChildren(bagToFind, parentBags, 1)
	fmt.Println("Total children: ", bagCount-1)
}

func countChildren(currBag string, parentBags map[string][]BagChild, numBags int) int {
	// Lookup children
	childBags := 0
	children := parentBags[currBag]
	for _, child := range children {
		fmt.Println("Found child: ", child.num, child.color, childBags, numBags)
		childBags += (countChildren(child.color, parentBags, child.num) * numBags)
	}
	return numBags + childBags
}

func findParents(currBag string, parentLookup map[string]map[string]bool, uniqueParents map[string]bool) {

	parents := parentLookup[currBag]
	for parent := range parents {
		_, exists := uniqueParents[parent]
		if !exists {
			uniqueParents[parent] = true
			findParents(parent, parentLookup, uniqueParents)
		}
	}

}
