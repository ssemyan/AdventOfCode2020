package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
	"strings"
)

type tile struct {
	tileNum   int
	topStr    string
	bottomStr string
	leftStr   string
	rightStr  string
	fullStr   [][]rune
}

func (t *tile) updateSides() {
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
	for y := 0; y < 10; y++ {
		for i := 0; i < 10; i++ {
			fmt.Printf(string(t.fullStr[y][i]))
		}
		fmt.Println()
	}
	fmt.Println()
}

type tileMatch struct {
	tileNum     int
	topMatch    int
	rightMatch  int
	bottomMatch int
	leftMatch   int
}

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day20/testdata.txt")

	// Load tiles into array
	tiles := make(map[int]tile)
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
				tiles[tileNumber] = newTile
			}
			lineNum++
		}
	}

	fmt.Println("Tiles loaded: ", len(tiles))

	// part one - find the corner tiles (where only two edges have matches)
	corners := []tile{}
	cornerTotal := 1

	for _, tile := range tiles {
		bTopMatch, bBottomMatch, bLeftMatch, bRightMatch := false, false, false, false
		nMatchedEdges := 0

		//fmt.Println("Matching ", tile.tileNum)
		//tile.print()

		for _, ntile := range tiles {

			if tile.tileNum != ntile.tileNum {

				if !bTopMatch && checkSide(tile.topStr, ntile) {
					// top
					bTopMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on top with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bBottomMatch && checkSide(tile.bottomStr, ntile) {
					// bottom
					bBottomMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on bottom with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bLeftMatch && checkSide(tile.leftStr, ntile) {
					// left
					bLeftMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on left with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bRightMatch && checkSide(tile.rightStr, ntile) {
					// right
					bRightMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on right with %d\n", tile.tileNum, ntile.tileNum)
				}
			}
		}
		if nMatchedEdges == 2 {
			// Found a corner
			corners = append(corners, tile)
			cornerTotal *= tile.tileNum
			fmt.Println("Found corner: ", tile.tileNum)
		}
	}
	fmt.Printf("Found %d corners with a total of %d\n", len(corners), cornerTotal)

	// Part two -- put the tiles together by rotating or flipping when match is found
	for _, tile := range tiles {
		bTopMatch, bBottomMatch, bLeftMatch, bRightMatch := false, false, false, false
		nMatchedEdges := 0

		fmt.Println("Matching ", tile.tileNum)
		tile.print()

		for _, ntile := range tiles {

			if tile.tileNum != ntile.tileNum {

				if !bTopMatch && checkSide2(tile.topStr, ntile, tiles, "top") {
					// top
					bTopMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on top with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bBottomMatch && checkSide2(tile.bottomStr, ntile, tiles, "bottom") {
					// bottom
					bBottomMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on bottom with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bLeftMatch && checkSide2(tile.leftStr, ntile, tiles, "left") {
					// left
					bLeftMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on left with %d\n", tile.tileNum, ntile.tileNum)

				} else if !bRightMatch && checkSide2(tile.rightStr, ntile, tiles, "right") {
					// right
					bRightMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on right with %d\n", tile.tileNum, ntile.tileNum)
				}
			}
		}
		if nMatchedEdges == 2 {
			// Found a corner
			corners = append(corners, tile)
			cornerTotal *= tile.tileNum
			fmt.Println("Found corner: ", tile.tileNum)
		}
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

	// 	rev(ntile.topStr)) || (tileSide == ntile.topStr) ||
	// (tileSide == rev(ntile.bottomStr)) || (tileSide == ntile.bottomStr) ||
	// (tileSide == rev(ntile.rightStr)) || (tileSide == ntile.rightStr) ||
	// (tileSide == rev(ntile.leftStr)) || (tileSide == ntile.leftStr)

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

func flip(arr [][]rune, isVert bool) [][]rune {

	// Do the flip
	newArr := cloneArr(arr)
	for y := 0; y < len(arr); y++ {
		for x := 0; x < len(arr[0]); x++ {
			if isVert {
				newArr[y][x] = arr[y][len(arr)-x-1]
			} else {
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
