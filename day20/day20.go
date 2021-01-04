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
	rightStr  string
	bottomStr string
	leftStr   string
}

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day20/data.txt")

	// Load tiles into array
	tiles := []tile{}
	var leftSide, rightSide, topSide string
	var lineNum, tileNumber int
	for n, line := range lines {

		// Tile line starts new tile
		if strings.HasPrefix(line, "Tile ") {
			// Create new tile
			strTileNum := line[5:]
			tileNum, _ := strconv.Atoi(strTileNum[:len(strTileNum)-1])
			tileNumber = tileNum
			lineNum = 0
			leftSide = ""
			rightSide = ""
			topSide = ""

		} else if line != "" {

			// Calculate top if at top
			if lineNum == 0 {
				topSide = line
			}

			// Calculate each side
			leftSide += string(line[0])
			rightSide += string(line[len(line)-1])

			if n == len(lines)-1 || lines[n+1] == "" {
				// End of tile
				newTile := tile{
					bottomStr: line,
					leftStr:   leftSide,
					rightStr:  rightSide,
					topStr:    topSide,
					tileNum:   tileNumber,
				}
				tiles = append(tiles, newTile)
			}
			lineNum++
		}
	}

	fmt.Println("Tiles loaded: ", len(tiles))

	// find the corner tiles (where only two edges have matches)
	corners := make(map[int]bool)
	cornerTotal := 1
	for _, tile := range tiles {
		bTopMatch, bBottomMatch, bLeftMatch, bRightMatch := false, false, false, false
		nMatchedEdges := 0
		for _, ntile := range tiles {

			if tile.tileNum != ntile.tileNum {

				// top
				if !bTopMatch && checkSide(tile.topStr, ntile) {
					bTopMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on top with %d\n", tile.tileNum, ntile.tileNum)
				}

				// bottom
				if !bBottomMatch && checkSide(tile.bottomStr, ntile) {
					bBottomMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on bottom with %d\n", tile.tileNum, ntile.tileNum)
				}

				// left
				if !bLeftMatch && checkSide(tile.leftStr, ntile) {
					bLeftMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on left with %d\n", tile.tileNum, ntile.tileNum)
				}

				// right
				if !bRightMatch && checkSide(tile.rightStr, ntile) {
					bRightMatch = true
					nMatchedEdges++
					//fmt.Printf("Tile %d found match on right with %d\n", tile.tileNum, ntile.tileNum)
				}
			}
		}
		if nMatchedEdges == 2 {
			// Found a corner
			corners[tile.tileNum] = true
			cornerTotal *= tile.tileNum
			fmt.Println("Found corner: ", tile.tileNum)
		}
	}
	fmt.Printf("Found %d corners with a total of %d\n", len(corners), cornerTotal)

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
