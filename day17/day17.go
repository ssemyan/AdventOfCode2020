package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
)

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day17/data.txt")

	cubes := make(map[string]bool)

	// Load initial state and set min and max for each dim
	xmin, xmax, ymin, ymax, zmin, zmax := 0, len(lines[0])-1, 0, len(lines)-1, 0, 0

	// Start at 0,0,0
	for y, line := range lines {
		for x, chr := range line {
			if chr == '#' {
				key := getKey(x, y, 0)
				cubes[key] = true
			}
		}
	}
	fmt.Println("Loaded cubes: ", len(cubes))

	// for key := range cubes {
	// 	fmt.Println(key)
	// }

	// Do six cycles
	for i := 0; i < 6; i++ {

		// for z := zmin; z < zmax+1; z++ {
		// 	fmt.Println("Z Dim: ", z)
		// 	for y := ymin; y < ymax+1; y++ {
		// 		for x := xmin; x < xmax+1; x++ {
		// 			if isActive(x, y, z, cubes) {
		// 				fmt.Printf("#")
		// 			} else {
		// 				fmt.Printf(".")
		// 			}
		// 		}
		// 		fmt.Println()
		// 	}
		// }

		toActivate := make(map[string]bool)
		toDeactivate := make(map[string]bool)

		xmint, xmaxt, ymint, ymaxt, zmint, zmaxt := xmin, xmax, ymin, ymax, zmin, zmax
		for z := zmin - 1; z <= zmax+1; z++ {

			for y := ymin - 1; y <= ymax+1; y++ {

				for x := xmin - 1; x <= xmax+1; x++ {

					// Count active around
					nActiveCnt := 0
					for zi := z - 1; zi <= z+1; zi++ {
						for yi := y - 1; yi <= y+1; yi++ {
							for xi := x - 1; xi <= x+1; xi++ {
								// be sure to ignore current cell
								if !(xi == x && yi == y && z == zi) && isActive(xi, yi, zi, cubes) {
									nActiveCnt++
								}
							}
						}
					}

					// If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes inactive.
					if isActive(x, y, z, cubes) {
						if nActiveCnt < 2 || nActiveCnt > 3 {
							toDeactivate[getKey(x, y, z)] = true
						}
					} else if nActiveCnt == 3 {
						// If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.
						toActivate[getKey(x, y, z)] = true

						// may need to reset max/min
						if x < xmint {
							xmint = x
						}
						if x > xmaxt {
							xmaxt = x
						}
						if y < ymint {
							ymint = y
						}
						if y > ymaxt {
							ymaxt = y
						}
						if z < zmint {
							zmint = z
						}
						if z > zmaxt {
							zmaxt = z
						}
					}
				}
			}
		}

		// Activate and deactivate
		for key := range toDeactivate {
			delete(cubes, key)
		}

		for key := range toActivate {
			cubes[key] = true
		}

		// may need to reset max/min
		xmin, xmax, ymin, ymax, zmin, zmax = xmint, xmaxt, ymint, ymaxt, zmint, zmaxt

		fmt.Println("Total active: ", len(cubes))
	}
}

func getKey(x int, y int, z int) string {
	return fmt.Sprintf("%d|%d|%d", x, y, z)
}

func isActive(x int, y int, z int, cubes map[string]bool) bool {
	_, exists := cubes[getKey(x, y, z)]
	return exists
}
