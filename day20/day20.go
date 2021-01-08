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

func (t tile) print() {
	printArr(t.fullStr)
}

type tileMatch struct {
	tileNum     int
	topMatch    int
	rightMatch  int
	bottomMatch int
	leftMatch   int
	isCorner    bool
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

	// part one - find the corner tiles (where only two edges have matches)
	// var topLeft int
	// cornerTotal := 1

	// save the matches for part two
	matches := make(map[int]tileMatch)

	// start matching tiles from the last tile found above
	findTileMatches(*tiles[tileNumber], tiles, matches)

	// for _, tile := range tiles {
	// 	bTopMatch, bBottomMatch, bLeftMatch, bRightMatch := false, false, false, false
	// 	nMatchedEdges := 0
	// 	match := tileMatch{}
	// 	tileNums := []int{}

	// 	//fmt.Println("Matching ", tile.tileNum)
	// 	//tile.print()

	// 	for _, ntile := range tiles {

	// 		if tile.tileNum != ntile.tileNum {

	// 			if !bTopMatch && checkSide(tile.topStr, *ntile) {
	// 				// top
	// 				bTopMatch = true
	// 				nMatchedEdges++
	// 				tileNums = append(tileNums, ntile.tileNum)
	// 				match.topMatch = ntile.tileNum
	// 				//fmt.Printf("Tile %d found match on top with %d\n", tile.tileNum, ntile.tileNum)

	// 			} else if !bBottomMatch && checkSide(tile.bottomStr, *ntile) {
	// 				// bottom
	// 				bBottomMatch = true
	// 				nMatchedEdges++
	// 				match.bottomMatch = ntile.tileNum
	// 				//fmt.Printf("Tile %d found match on bottom with %d\n", tile.tileNum, ntile.tileNum)

	// 			} else if !bLeftMatch && checkSide(tile.leftStr, *ntile) {
	// 				// left
	// 				bLeftMatch = true
	// 				nMatchedEdges++
	// 				match.leftMatch = ntile.tileNum
	// 				//fmt.Printf("Tile %d found match on left with %d\n", tile.tileNum, ntile.tileNum)

	// 			} else if !bRightMatch && checkSide(tile.rightStr, *ntile) {
	// 				// right
	// 				bRightMatch = true
	// 				nMatchedEdges++
	// 				match.rightMatch = ntile.tileNum
	// 				//fmt.Printf("Tile %d found match on right with %d\n", tile.tileNum, ntile.tileNum)
	// 			}
	// 		}
	// 	}
	// 	if nMatchedEdges == 2 {
	// 		// Found a corner
	// 		if match.topMatch == 0 && match.leftMatch == 0 {
	// 			topLeft = tile.tileNum
	// 		}
	// 		cornerTotal *= tile.tileNum
	// 		fmt.Println("Found corner: ", tile.tileNum)
	// 	}
	// 	matches[tile.tileNum] = match
	// }
	// Count corners
	cornerTotal := 1
	var topLeft int
	for tileNum, match := range matches {
		// Calculate corner
		if match.isCorner {
			cornerTotal *= tileNum
		}
		// look for top left (where we will start rotating from)
		if match.topMatch == 0 && match.leftMatch == 0 {
			topLeft = tileNum
		}
	}
	fmt.Printf("Total of corners: %d\n", cornerTotal)

	// Part two -- put the tiles together by rotating or flipping following previous matching
	// start with the top left corner (mark it as set)
	start := tiles[topLeft]
	// start.topRotationSet = true
	// start.bottomRotationSet = true
	// start.leftRotationSet = true
	// start.rightRotationSet = true
	findMatches(*start, tiles, matches)

	// Everything should be in order now so we can start looking for monsters ;-)

	// make the master map - start by cloning the first tile, then adding on to it
	fmt.Println("Starting with tile ", start.tileNum)
	seamap := cloneArr(start.fullStr)
	curY := 0
	for {
		match := matches[start.tileNum]

		// add tiles to the right
		rightMatch := match.rightMatch
		for {
			if rightMatch == 0 {
				break
			}
			fmt.Println("Adding tile ", rightMatch)
			tile := tiles[rightMatch]
			fullStr := tile.fullStr
			for y := 0; y < len(fullStr); y++ {
				seamap[curY+y] = append(seamap[curY+y], tile.fullStr[y]...)
			}
			rightMatch = matches[tile.tileNum].rightMatch
		}
		break
		// // add the next tile down
		// bottomMatch := match.bottomMatch
		// if bottomMatch == 0 {
		// 	break
		// }

		// start = tiles[bottomMatch]

		// // add the current tile
		// nextRow := []rune{}
		// for y := 0; y < len(start.fullStr); y++ {
		// 	nextRow = append(nextRow, start.fullStr[y]...)
		// }

	}
	printArr(seamap)
}

func findTileMatches(tile tile, tiles map[int]*tile, matches map[int]tileMatch) {

	// Don't match previously matched tiles
	_, exists := matches[tile.tileNum]
	if !exists {
		bTopMatch, bBottomMatch, bLeftMatch, bRightMatch := false, false, false, false
		match := tileMatch{}
		matchCount := 0

		fmt.Println("Matching ", tile.tileNum)
		//tile.print()

		for _, ntile := range tiles {

			if tile.tileNum != ntile.tileNum {

				if !bTopMatch && checkSide(tile.topStr, *ntile) { // top
					bTopMatch = true
					match.topMatch = ntile.tileNum
					matchCount++
					//fmt.Printf("Tile %d found match on top with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bBottomMatch && checkSide(tile.bottomStr, *ntile) {
					// bottom
					bBottomMatch = true
					match.bottomMatch = ntile.tileNum
					matchCount++
					//fmt.Printf("Tile %d found match on bottom with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bLeftMatch && checkSide(tile.leftStr, *ntile) {
					// left
					bLeftMatch = true
					match.leftMatch = ntile.tileNum
					matchCount++
					//fmt.Printf("Tile %d found match on left with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bRightMatch && checkSide(tile.rightStr, *ntile) {
					// right
					bRightMatch = true
					match.rightMatch = ntile.tileNum
					matchCount++
					//fmt.Printf("Tile %d found match on right with %d\n", tile.tileNum, ntile.tileNum)
				}
			}
		}
		// Corners only have two matches
		if matchCount == 2 {
			match.isCorner = true
		}
		matches[tile.tileNum] = match

		// Now find matches for all the sides
		if match.topMatch != 0 {
			findTileMatches(*tiles[match.topMatch], tiles, matches)
		}
		if match.bottomMatch != 0 {
			findTileMatches(*tiles[match.bottomMatch], tiles, matches)
		}
		if match.leftMatch != 0 {
			findTileMatches(*tiles[match.leftMatch], tiles, matches)
		}
		if match.rightMatch != 0 {
			findTileMatches(*tiles[match.rightMatch], tiles, matches)
		}
	}
}

func rotateMatch(tile tile, tiles map[int]*tile, matches map[int]tileMatch) {

	fmt.Println("Rotating matches to ", tile.tileNum)
	//tile.print()

	// get matches
	matchInfo := matches[tile.tileNum]

	// top
	flip := ""
	rotateLeft := 0
	if matchInfo.topMatch != 0 {
		ntile := tiles[matchInfo.topMatch]
		if !ntile.topRotationSet {
			tileSide := tile.topStr
			if tileSide == ntile.topStr {
				flip = "horz"
			} else if tileSide == ntile.bottomStr {
				// nothing needed
			} else if tileSide == ntile.leftStr {
				rotateLeft = 1 // Rot left 90
			} else if tileSide == ntile.rightStr {
				rotateLeft = 3 // Rot right 90
				flip = "vert"  // flip vertically
			} else if tileSide == rev(ntile.bottomStr) {
				flip = "vert" // flip vertically
			} else if tileSide == rev(ntile.topStr) {
				rotateLeft = 2 // rot 180
			} else if tileSide == rev(ntile.leftStr) {
				rotateLeft = 1 // Rot left 90
				flip = "vert"  // flip vertically
			} else if tileSide == rev(ntile.rightStr) {
				rotateLeft = 3 // Rot right 90
				flip = "vert"  // flip vertically
			}
			ntile.topRotationSet = true
			processTile(ntile, matches, rotateLeft, flip, tiles)
		}
	}

	// bottom
	flip = ""
	rotateLeft = 0
	if matchInfo.bottomMatch != 0 {
		ntile := tiles[matchInfo.bottomMatch]
		if !ntile.bottomRotationSet {
			tileSide := tile.topStr
			if tileSide == ntile.topStr {
				// nothing needed
			} else if tileSide == ntile.bottomStr {
				flip = "horz"
			} else if tileSide == ntile.leftStr {
				rotateLeft = 3 // Rot right 90
				flip = "vert"
			} else if tileSide == ntile.rightStr {
				rotateLeft = 1 // Rot left 90
			} else if tileSide == rev(ntile.bottomStr) {
				rotateLeft = 2 // rot 180
			} else if tileSide == rev(ntile.topStr) {
				flip = "vert" // flip vertically
			} else if tileSide == rev(ntile.leftStr) {
				rotateLeft = 3 // Rot left 270
			} else if tileSide == rev(ntile.rightStr) {
				rotateLeft = 1 // Rot left 90
				flip = "vert"  // flip vertically
			}
			ntile.bottomRotationSet = true
			processTile(ntile, matches, rotateLeft, flip, tiles)
		}
	}

	// left
	flip = ""
	rotateLeft = 0
	if matchInfo.leftMatch != 0 {
		ntile := tiles[matchInfo.leftMatch]
		if !ntile.leftRotationSet {
			tileSide := tile.topStr
			if tileSide == ntile.topStr {
				rotateLeft = 3 // Rot left 270
			} else if tileSide == ntile.bottomStr {
				rotateLeft = 1 // Rot left 90
				flip = "horz"
			} else if tileSide == ntile.leftStr {
				flip = "vert" // flip vertically
			} else if tileSide == ntile.rightStr {
				// nothing needed
			} else if tileSide == rev(ntile.bottomStr) {
				rotateLeft = 1 // Rot left 90
			} else if tileSide == rev(ntile.topStr) {
				rotateLeft = 3 // Rot left 270
				flip = "horz"  // flip vertically
			} else if tileSide == rev(ntile.leftStr) {
				rotateLeft = 2 // Rot left 180
			} else if tileSide == rev(ntile.rightStr) {
				flip = "horz" // flip vertically
			}
			ntile.leftRotationSet = true
			processTile(ntile, matches, rotateLeft, flip, tiles)
		}
	}

	// right
	flip = ""
	rotateLeft = 0
	if matchInfo.rightMatch != 0 {
		ntile := tiles[matchInfo.rightMatch]
		if !ntile.rightRotationSet {
			tileSide := tile.topStr
			if tileSide == ntile.topStr {
				rotateLeft = 1
				flip = "horz"
			} else if tileSide == ntile.bottomStr {
				rotateLeft = 3
			} else if tileSide == ntile.leftStr {
				// nothing needed
			} else if tileSide == ntile.rightStr {
				flip = "vert" // flip vertically
			} else if tileSide == rev(ntile.bottomStr) {
				rotateLeft = 3
			} else if tileSide == rev(ntile.topStr) {
				rotateLeft = 1 // Rot left 270
				flip = "horz"  // flip vertically
			} else if tileSide == rev(ntile.leftStr) {
				flip = "horz" // flip vertically
			} else if tileSide == rev(ntile.rightStr) {
				rotateLeft = 2 // Rot left 180
			}
			ntile.rightRotationSet = true
			processTile(ntile, matches, rotateLeft, flip, tiles)
		}
	}
}

func findMatches(tile tile, tiles map[int]*tile, matches map[int]tileMatch) {

	fmt.Println("Rotating matches to ", tile.tileNum)
	//tile.print()

	// get matches
	matchInfo := matches[tile.tileNum]

	// top
	flip := ""
	rotateLeft := 0
	if matchInfo.topMatch != 0 {
		ntile := tiles[matchInfo.topMatch]
		if !ntile.topRotationSet {
			tileSide := tile.topStr
			if tileSide == ntile.topStr {
				flip = "horz"
			} else if tileSide == ntile.bottomStr {
				// nothing needed
			} else if tileSide == ntile.leftStr {
				rotateLeft = 1 // Rot left 90
			} else if tileSide == ntile.rightStr {
				rotateLeft = 3 // Rot right 90
				flip = "vert"  // flip vertically
			} else if tileSide == rev(ntile.bottomStr) {
				flip = "vert" // flip vertically
			} else if tileSide == rev(ntile.topStr) {
				rotateLeft = 2 // rot 180
			} else if tileSide == rev(ntile.leftStr) {
				rotateLeft = 1 // Rot left 90
				flip = "vert"  // flip vertically
			} else if tileSide == rev(ntile.rightStr) {
				rotateLeft = 3 // Rot right 90
				flip = "vert"  // flip vertically
			}
			ntile.topRotationSet = true
			processTile(ntile, matches, rotateLeft, flip, tiles)
		}
	}

	// bottom
	flip = ""
	rotateLeft = 0
	if matchInfo.bottomMatch != 0 {
		ntile := tiles[matchInfo.bottomMatch]
		if !ntile.bottomRotationSet {
			tileSide := tile.topStr
			if tileSide == ntile.topStr {
				// nothing needed
			} else if tileSide == ntile.bottomStr {
				flip = "horz"
			} else if tileSide == ntile.leftStr {
				rotateLeft = 3 // Rot right 90
				flip = "vert"
			} else if tileSide == ntile.rightStr {
				rotateLeft = 1 // Rot left 90
			} else if tileSide == rev(ntile.bottomStr) {
				rotateLeft = 2 // rot 180
			} else if tileSide == rev(ntile.topStr) {
				flip = "vert" // flip vertically
			} else if tileSide == rev(ntile.leftStr) {
				rotateLeft = 3 // Rot left 270
			} else if tileSide == rev(ntile.rightStr) {
				rotateLeft = 1 // Rot left 90
				flip = "vert"  // flip vertically
			}
			ntile.bottomRotationSet = true
			processTile(ntile, matches, rotateLeft, flip, tiles)
		}
	}

	// left
	flip = ""
	rotateLeft = 0
	if matchInfo.leftMatch != 0 {
		ntile := tiles[matchInfo.leftMatch]
		if !ntile.leftRotationSet {
			tileSide := tile.topStr
			if tileSide == ntile.topStr {
				rotateLeft = 3 // Rot left 270
			} else if tileSide == ntile.bottomStr {
				rotateLeft = 1 // Rot left 90
				flip = "horz"
			} else if tileSide == ntile.leftStr {
				flip = "vert" // flip vertically
			} else if tileSide == ntile.rightStr {
				// nothing needed
			} else if tileSide == rev(ntile.bottomStr) {
				rotateLeft = 1 // Rot left 90
			} else if tileSide == rev(ntile.topStr) {
				rotateLeft = 3 // Rot left 270
				flip = "horz"  // flip vertically
			} else if tileSide == rev(ntile.leftStr) {
				rotateLeft = 2 // Rot left 180
			} else if tileSide == rev(ntile.rightStr) {
				flip = "horz" // flip vertically
			}
			ntile.leftRotationSet = true
			processTile(ntile, matches, rotateLeft, flip, tiles)
		}
	}

	// right
	flip = ""
	rotateLeft = 0
	if matchInfo.rightMatch != 0 {
		ntile := tiles[matchInfo.rightMatch]
		if !ntile.rightRotationSet {
			tileSide := tile.topStr
			if tileSide == ntile.topStr {
				rotateLeft = 1
				flip = "horz"
			} else if tileSide == ntile.bottomStr {
				rotateLeft = 3
			} else if tileSide == ntile.leftStr {
				// nothing needed
			} else if tileSide == ntile.rightStr {
				flip = "vert" // flip vertically
			} else if tileSide == rev(ntile.bottomStr) {
				rotateLeft = 3
			} else if tileSide == rev(ntile.topStr) {
				rotateLeft = 1 // Rot left 270
				flip = "horz"  // flip vertically
			} else if tileSide == rev(ntile.leftStr) {
				flip = "horz" // flip vertically
			} else if tileSide == rev(ntile.rightStr) {
				rotateLeft = 2 // Rot left 180
			}
			ntile.rightRotationSet = true
			processTile(ntile, matches, rotateLeft, flip, tiles)
		}
	}
}

func processTile(ntile *tile, matches map[int]tileMatch, rotateLeft int, flip string, tiles map[int]*tile) {

	if rotateLeft > 0 {
		ntile.fullStr = rotate(ntile.fullStr, rotateLeft)
	}
	if flip != "" {
		ntile.fullStr = doFlip(ntile.fullStr, flip)
	}
	//ntile.rotationSet = true

	// fix the matches to match the flip and rotation
	match := matches[ntile.tileNum]

	for i := 0; i < rotateLeft; i++ {
		tmp := match.topMatch
		match.topMatch = match.rightMatch
		match.rightMatch = match.bottomMatch
		match.bottomMatch = match.leftMatch
		match.leftMatch = tmp
	}

	if flip == "horz" {
		tmp := match.topMatch
		match.topMatch = match.bottomMatch
		match.bottomMatch = tmp
	} else if flip == "horz" {
		tmp := match.leftMatch
		match.leftMatch = match.rightMatch
		match.rightMatch = tmp
	}

	matches[ntile.tileNum] = match

	// now find the match for the recently processed tile
	findMatches(*ntile, tiles, matches)
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

func checkSide2(tileSide string, ntile tile, tiles map[int]tile, side string) bool {

	// can potentially match any side normal or flipped
	rotateLeft := 0
	flip := ""

	// rotateBase := 0
	// leftRight := true

	// switch side {
	// case "top":
	// 	rotateBase = 0
	// 	leftRight = true
	// case "bottom":
	// 	rotateBase = 2
	// 	leftRight = true
	// case "left":
	// 	rotateBase = 1
	// 	leftRight = false
	// case "right":
	// 	rotateBase = 3
	// 	leftRight = false
	// }

	if side == "top" {
		if tileSide == ntile.topStr {
			flip = "horz"
		} else if tileSide == ntile.bottomStr {
			// nothing needed
		} else if tileSide == ntile.leftStr {
			rotateLeft = 1 // Rot left 90
		} else if tileSide == ntile.rightStr {
			rotateLeft = 3 // Rot right 90
			flip = "vert"  // flip vertically
		} else if tileSide == rev(ntile.bottomStr) {
			flip = "vert" // flip vertically
		} else if tileSide == rev(ntile.topStr) {
			rotateLeft = 2 // rot 180
		} else if tileSide == rev(ntile.leftStr) {
			rotateLeft = 1 // Rot left 90
			flip = "vert"  // flip vertically
		} else if tileSide == rev(ntile.rightStr) {
			rotateLeft = 3 // Rot right 90
			flip = "vert"  // flip vertically
		} else {
			// not a match
			return false
		}
	} else if side == "bottom" {
		if tileSide == ntile.topStr {
			// nothing needed
		} else if tileSide == ntile.bottomStr {
			flip = "horz"
		} else if tileSide == ntile.leftStr {
			rotateLeft = 3 // Rot right 90
			flip = "vert"
		} else if tileSide == ntile.rightStr {
			rotateLeft = 1 // Rot left 90
		} else if tileSide == rev(ntile.bottomStr) {
			rotateLeft = 2 // rot 180
		} else if tileSide == rev(ntile.topStr) {
			flip = "vert" // flip vertically
		} else if tileSide == rev(ntile.leftStr) {
			rotateLeft = 3 // Rot left 270
		} else if tileSide == rev(ntile.rightStr) {
			rotateLeft = 1 // Rot left 90
			flip = "vert"  // flip vertically
		} else {
			// not a match
			return false
		}
	} else if side == "left" {
		if tileSide == ntile.topStr {
			rotateLeft = 3 // Rot left 270
		} else if tileSide == ntile.bottomStr {
			rotateLeft = 1 // Rot left 90
			flip = "horz"
		} else if tileSide == ntile.leftStr {
			flip = "vert" // flip vertically
		} else if tileSide == ntile.rightStr {
			// nothing needed
		} else if tileSide == rev(ntile.bottomStr) {
			rotateLeft = 1 // Rot left 90
		} else if tileSide == rev(ntile.topStr) {
			rotateLeft = 3 // Rot left 270
			flip = "horz"  // flip vertically
		} else if tileSide == rev(ntile.leftStr) {
			rotateLeft = 2 // Rot left 180
		} else if tileSide == rev(ntile.rightStr) {
			flip = "horz" // flip vertically
		} else {
			// not a match
			return false
		}
	} else if side == "right" {
		if tileSide == ntile.topStr {
			rotateLeft = 1
			flip = "horz"
		} else if tileSide == ntile.bottomStr {
			rotateLeft = 3
		} else if tileSide == ntile.leftStr {
			// nothing needed
		} else if tileSide == ntile.rightStr {
			flip = "vert" // flip vertically
		} else if tileSide == rev(ntile.bottomStr) {
			rotateLeft = 3
		} else if tileSide == rev(ntile.topStr) {
			rotateLeft = 1 // Rot left 270
			flip = "horz"  // flip vertically
		} else if tileSide == rev(ntile.leftStr) {
			flip = "horz" // flip vertically
		} else if tileSide == rev(ntile.rightStr) {
			rotateLeft = 2 // Rot left 180
		} else {
			// not a match
			return false
		}
	} else {
		panic("Unknown side")
	}

	// Print before out
	fmt.Printf("Found match for %s: %d\n", side, ntile.tileNum)
	ntile.print()

	bTransformed := false
	var newFullStr [][]rune
	copy(newFullStr, ntile.fullStr)
	// Do any rotation
	for rot := 0; rot < rotateLeft; rot++ {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				newFullStr[y][x] = ntile.fullStr[x][9-y]
			}
		}
		copy(ntile.fullStr, newFullStr)
		ntile.print()
		bTransformed = true
	}

	// Do any flip
	if flip == "horz" {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				newFullStr[y][x] = ntile.fullStr[9-y][x]
			}
		}
		bTransformed = true

	} else if flip == "vert" {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				newFullStr[y][x] = ntile.fullStr[y][9-x]
			}
		}
		bTransformed = true
	}

	// Print after out if changed
	if bTransformed {
		copy(ntile.fullStr, newFullStr)
		tiles[ntile.tileNum] = ntile
		ntile.print()
	}

	return true
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
