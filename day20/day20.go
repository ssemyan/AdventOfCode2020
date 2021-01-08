package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
	"strings"
)

type tile struct {
	tileNum     int
	topStr      string
	bottomStr   string
	leftStr     string
	rightStr    string
	fullStr     [][]rune
	rotated     bool
	matched     bool
	topMatch    int
	rightMatch  int
	bottomMatch int
	leftMatch   int
	isCorner    bool
}

func (t *tile) updateSides() {
	// Update the sides so they are cached
	t.topStr = string(t.fullStr[0])
	t.bottomStr = string(t.fullStr[len(t.fullStr)-1])
	t.leftStr = string(t.sideStr(true))
	t.rightStr = string(t.sideStr(false))
}

func (t tile) sideStr(isLeft bool) []rune {

	retStr := make([]rune, len(t.fullStr))
	colNum := 0
	if !isLeft {
		colNum = len(t.fullStr) - 1
	}
	for i := 0; i < len(t.fullStr); i++ {
		retStr[i] = t.fullStr[i][colNum]
	}
	return retStr
}

func (t tile) getMap() [][]rune {

	// for the map skip the sides
	retStr := make([][]rune, len(t.fullStr)-2)
	for y := 1; y < len(t.fullStr)-1; y++ {
		retStr[y-1] = t.fullStr[y][1 : len(t.fullStr[0])-1]
	}
	return retStr
}

func (t tile) print() {
	printArr(t.fullStr)
}

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day20/testdata.txt")

	// Load tiles into array
	tiles := make(map[int]*tile)
	var fullStr [][]rune
	var lineNum, tileNumber int
	for n, line := range lines {

		// Tile line starts new tile
		if strings.HasPrefix(line, "Tile ") {
			// Create new tile
			strTileNum := line[5:]
			tileNum, _ := strconv.Atoi(strTileNum[:len(strTileNum)-1])
			tileNumber = tileNum
			lineNum = 0
			fullStr = make([][]rune, 10)

		} else if line != "" {

			// Add to full str
			fullStr[lineNum] = []rune(line)

			if n == len(lines)-1 || lines[n+1] == "" {
				// End of tile
				newTile := tile{
					tileNum: tileNumber,
					fullStr: fullStr,
				}
				newTile.updateSides()
				tiles[tileNumber] = &newTile
			}
			lineNum++
		}
	}

	fmt.Println("Tiles loaded: ", len(tiles))

	// start matching tiles from the last tile found above
	// mark it as rotated so we don't rotate it again
	start := tiles[tileNumber]
	start.rotated = true
	findTileMatches(start, tiles)

	// Count corners and find top left for part two
	cornerTotal := 1
	var topLeft int
	for tileNum, ctile := range tiles {
		// Calculate corner
		if ctile.isCorner {
			cornerTotal *= tileNum
		}
		// look for top left (where we will start building the map from)
		if ctile.topMatch == 0 && ctile.leftMatch == 0 {
			topLeft = ctile.tileNum
		}
	}
	fmt.Printf("Total of corners: %d\n", cornerTotal)

	// Part two -- put the tiles together by rotating or flipping following previous matching

	// make the master map - start by cloning the first tile, then adding on to it
	seamap := [][]rune{}
	curY := 0
	rowMatch := topLeft
	for {
		// add start of the row
		fmt.Printf("%d ", topLeft)
		start = tiles[rowMatch]
		startMap := start.getMap()
		seamap = append(seamap, startMap...)

		// add tiles to the right
		colMatch := start.rightMatch
		for {
			if colMatch == 0 {
				break
			}
			fmt.Printf("%d ", colMatch)
			tile := tiles[colMatch]
			tileMap := tile.getMap()
			for y := 0; y < len(tileMap); y++ {
				seamap[curY+y] = append(seamap[curY+y], tileMap[y]...)
			}
			colMatch = tiles[tile.tileNum].rightMatch
		}
		fmt.Println()

		// add the next tile down
		rowMatch = start.bottomMatch
		if rowMatch == 0 {
			break
		}
		start = tiles[rowMatch]
		curY += len(start.getMap())
	}

	printArr(seamap)
}

func findTileMatches(tile *tile, tiles map[int]*tile) {

	// Don't match previously matched tiles
	if !tile.matched {
		bTopMatch, bBottomMatch, bLeftMatch, bRightMatch := false, false, false, false
		matchCount := 0

		fmt.Println("Matching ", tile.tileNum)
		//tile.print()

		for _, ntile := range tiles {

			if tile.tileNum != ntile.tileNum {

				if !bTopMatch && checkSide(tile.topStr, *ntile) { // top
					bTopMatch = true
					tile.topMatch = ntile.tileNum
					matchCount++
					//fmt.Printf("Tile %d found match on top with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bBottomMatch && checkSide(tile.bottomStr, *ntile) {
					// bottom
					bBottomMatch = true
					tile.bottomMatch = ntile.tileNum
					matchCount++
					//fmt.Printf("Tile %d found match on bottom with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bLeftMatch && checkSide(tile.leftStr, *ntile) {
					// left
					bLeftMatch = true
					tile.leftMatch = ntile.tileNum
					matchCount++
					//fmt.Printf("Tile %d found match on left with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bRightMatch && checkSide(tile.rightStr, *ntile) {
					// right
					bRightMatch = true
					tile.rightMatch = ntile.tileNum
					matchCount++
					//fmt.Printf("Tile %d found match on right with %d\n", tile.tileNum, ntile.tileNum)
				}
			}
		}
		// Corners only have two matches
		if matchCount == 2 {
			tile.isCorner = true
		}

		// Mark current tile as matched
		tile.matched = true

		// Rotate surrounding tiles and match
		if tile.topMatch != 0 {
			rotateMatch(tile.topStr, tiles[tile.topMatch], "top")
			findTileMatches(tiles[tile.topMatch], tiles)
		}
		if tile.bottomMatch != 0 {
			rotateMatch(tile.bottomStr, tiles[tile.bottomMatch], "bottom")
			findTileMatches(tiles[tile.bottomMatch], tiles)
		}
		if tile.leftMatch != 0 {
			rotateMatch(tile.leftStr, tiles[tile.leftMatch], "left")
			findTileMatches(tiles[tile.leftMatch], tiles)
		}
		if tile.rightMatch != 0 {
			rotateMatch(tile.rightStr, tiles[tile.rightMatch], "right")
			findTileMatches(tiles[tile.rightMatch], tiles)
		}
	}
}

func rotateMatch(tileSide string, tile *tile, side string) {

	if !tile.rotated {
		fmt.Println("Rotating ", tile.tileNum)
		//tile.print()

		// top
		flip := ""
		rotateLeft := 0
		if side == "top" {
			if tileSide == tile.topStr {
				flip = "horz"
			} else if tileSide == tile.bottomStr {
				// nothing needed
			} else if tileSide == tile.leftStr {
				rotateLeft = 1 // Rot left 90
			} else if tileSide == tile.rightStr {
				rotateLeft = 3 // Rot right 90
				flip = "vert"  // flip vertically
			} else if tileSide == rev(tile.bottomStr) {
				flip = "vert" // flip vertically
			} else if tileSide == rev(tile.topStr) {
				rotateLeft = 2 // rot 180
			} else if tileSide == rev(tile.leftStr) {
				rotateLeft = 1 // Rot left 90
				flip = "vert"  // flip vertically
			} else if tileSide == rev(tile.rightStr) {
				rotateLeft = 3 // Rot right 90
				flip = "vert"  // flip vertically
			}

		} else if side == "bottom" {
			if tileSide == tile.topStr {
				// nothing needed
			} else if tileSide == tile.bottomStr {
				flip = "horz"
			} else if tileSide == tile.leftStr {
				rotateLeft = 3 // Rot right 90
				flip = "vert"
			} else if tileSide == tile.rightStr {
				rotateLeft = 1 // Rot left 90
			} else if tileSide == rev(tile.bottomStr) {
				rotateLeft = 2 // rot 180
			} else if tileSide == rev(tile.topStr) {
				flip = "vert" // flip vertically
			} else if tileSide == rev(tile.leftStr) {
				rotateLeft = 3 // Rot left 270
			} else if tileSide == rev(tile.rightStr) {
				rotateLeft = 1 // Rot left 90
				flip = "vert"  // flip vertically
			}

		} else if side == "left" {
			if tileSide == tile.topStr {
				rotateLeft = 3 // Rot left 270
			} else if tileSide == tile.bottomStr {
				rotateLeft = 1 // Rot left 90
				flip = "horz"
			} else if tileSide == tile.leftStr {
				flip = "vert" // flip vertically
			} else if tileSide == tile.rightStr {
				// nothing needed
			} else if tileSide == rev(tile.bottomStr) {
				rotateLeft = 1 // Rot left 90
			} else if tileSide == rev(tile.topStr) {
				rotateLeft = 3 // Rot left 270
				flip = "horz"  // flip vertically
			} else if tileSide == rev(tile.leftStr) {
				rotateLeft = 2 // Rot left 180
			} else if tileSide == rev(tile.rightStr) {
				flip = "horz" // flip vertically
			}

		} else if side == "right" {
			if tileSide == tile.topStr {
				rotateLeft = 1
				flip = "horz"
			} else if tileSide == tile.bottomStr {
				rotateLeft = 3
			} else if tileSide == tile.leftStr {
				// nothing needed
			} else if tileSide == tile.rightStr {
				flip = "vert" // flip vertically
			} else if tileSide == rev(tile.bottomStr) {
				rotateLeft = 3
			} else if tileSide == rev(tile.topStr) {
				rotateLeft = 1 // Rot left 270
				flip = "horz"  // flip vertically
			} else if tileSide == rev(tile.leftStr) {
				flip = "horz" // flip vertically
			} else if tileSide == rev(tile.rightStr) {
				rotateLeft = 2 // Rot left 180
			}
		}

		// do the actual rotation
		if rotateLeft > 0 {
			tile.fullStr = rotate(tile.fullStr, rotateLeft)
		}
		if flip != "" {
			tile.fullStr = doFlip(tile.fullStr, flip)
		}
		// need to recalc the sides
		tile.updateSides()
		tile.rotated = true
	}
}

func rev(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func checkSide(tileSide string, ntile tile) bool {

	// can potentially match any side normal or flipped
	return (tileSide == rev(ntile.topStr)) || (tileSide == ntile.topStr) ||
		(tileSide == rev(ntile.bottomStr)) || (tileSide == ntile.bottomStr) ||
		(tileSide == rev(ntile.rightStr)) || (tileSide == ntile.rightStr) ||
		(tileSide == rev(ntile.leftStr)) || (tileSide == ntile.leftStr)
}

func rotate(arr [][]rune, rotateLeft int) [][]rune {

	// Do the nubmer of rotations left asked for
	var newArr [][]rune
	for i := 0; i < rotateLeft; i++ {
		newArr = cloneArr(arr)
		for y := 0; y < len(arr); y++ {
			for x := 0; x < len(arr[0]); x++ {
				newArr[y][x] = arr[x][len(arr)-y-1]
			}
		}
	}
	return newArr
}

func doFlip(arr [][]rune, flipType string) [][]rune {

	// Do the flip
	newArr := cloneArr(arr)
	for y := 0; y < len(arr); y++ {
		for x := 0; x < len(arr[0]); x++ {
			if flipType == "vert" {
				newArr[y][x] = arr[y][len(arr)-x-1]
			} else if flipType == "horz" {
				newArr[y][x] = arr[len(arr)-y-1][x]
			}
		}
	}
	return newArr
}

func cloneArr(arr [][]rune) [][]rune {

	// Multi dim arrays are a pain in golang, must roll our own deep copy
	newArr := make([][]rune, len(arr))
	for y := 0; y < len(arr); y++ {
		line := make([]rune, len(arr[0]))
		copy(line, arr[y])
		newArr[y] = line
	}
	return newArr
}

func printArr(arr [][]rune) {
	for y := 0; y < len(arr); y++ {
		for i := 0; i < len(arr[0]); i++ {
			fmt.Printf(string(arr[y][i]))
		}
		fmt.Println()
	}
	fmt.Println()
}
