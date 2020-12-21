package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day13/data.txt")

	//partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	departTime, _ := strconv.Atoi(lines[0])
	fmt.Println("Finding best bus for ", departTime)
	buses := strings.Split(lines[1], ",")

	bestBus, bestBusTime := 0, 0
	for _, bus := range buses {

		if bus != "x" {
			busNum, _ := strconv.Atoi(bus)

			nextTime := busNum * (departTime / busNum)
			if nextTime < departTime {
				nextTime += busNum
			}

			if bestBusTime == 0 || bestBusTime > nextTime {
				bestBusTime = nextTime
				bestBus = busNum
			}
			fmt.Printf("Bus: %d Nearest Depart Time: %d Delta: %d\n", busNum, nextTime, (nextTime - departTime))
		}
	}
	fmt.Printf("Best Bus: %d Depart Time: %d Delta: %d\n", bestBus, bestBusTime, (bestBusTime-departTime)*bestBus)
}

func partTwo(lines []string) {
	fmt.Println("Part two")
	bus := strings.Split(lines[1], ",")

	// Number to multiply (first num)
	toMultiply, _ := strconv.Atoi(bus[0])
	i, multiplier, toAdd, toCompare := 1, 0, 1, 0

	for {

		// Get next number
		for {
			if bus[i] != "x" {
				toCompareTemp, _ := strconv.Atoi(bus[i])
				toCompare = toCompareTemp
				break
			}
			i++
		}
		fmt.Println("Finding next common time: ", toMultiply, i, toCompare)

		tmultiplier := 1
		for {
			nSum := toMultiply * (multiplier + (tmultiplier * toAdd))
			if (nSum+i)%toCompare == 0 {
				// Found common
				multiplier += (tmultiplier * toAdd)
				toAdd *= toCompare
				fmt.Println("Common found: ", multiplier, nSum)
				break
			}
			tmultiplier++
		}

		i++
		if i > len(bus)-1 {
			break
		}
	}

}
