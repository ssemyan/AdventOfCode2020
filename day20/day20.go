package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
	"strings"
)

type tile struct {
	tileNum   int
	topNum    uint32
	rightNum  uint32
	bottomNum uint32
	leftNum   uint32
	pos       string
}

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day20/testdata.txt")

	// Load tiles
	tiles := []tile{}
	leftSide, rightSide := "", ""
	lineNum, topNum, tileNumber := 0, 0, 0
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

		} else if line != "" {

			// Calculate top if at top
			if lineNum == 0 {
				topNum = getCode(line)
			}

			// Calculate each side
			leftSide += string(line[0])
			rightSide += string(line[len(line)-1])

			if n == len(lines)-1 || lines[n+1] == "" {
				// End of tile
				newTile := tile{
					bottomNum: uint32(getCode(line)),
					leftNum:   uint32(getCode(leftSide)),
					rightNum:  uint32(getCode(rightSide)),
					topNum:    uint32(topNum),
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
	cornerTotal := 0
	for i, tile := range tiles {
		bTopMatch, bBottomMatch, bLeftMatch, bRightMatch := false, false, false, false
		nMatchedEdges := 0
		for n, ntile := range tiles {

			if n != i {
				// top
				tileSide := tile.topNum
				if !bTopMatch && ((tileSide == (^ntile.topNum)) || (tileSide == ntile.bottomNum) || (tileSide == (^ntile.rightNum)) || (tileSide == ntile.leftNum)) {
					bTopMatch = true
					nMatchedEdges++
				}

				// bottom
				tileSide = tile.bottomNum
				if !bBottomMatch && ((tileSide == ntile.topNum) || (tileSide == (^ntile.bottomNum)) || (tileSide == ntile.rightNum) || (tileSide == (^ntile.leftNum))) {
					bBottomMatch = true
					nMatchedEdges++
				}

				// left
				tileSide = tile.leftNum
				if !bLeftMatch && ((tileSide == ntile.topNum) || (tileSide == (^ntile.bottomNum)) || (tileSide == ntile.rightNum) || (tileSide == (^ntile.leftNum))) {
					bLeftMatch = true
					nMatchedEdges++
				}

				// right
				tileSide = tile.rightNum
				if !bRightMatch && ((tileSide == (^ntile.topNum)) || (tileSide == ntile.bottomNum) || (tileSide == (^ntile.rightNum)) || (tileSide == ntile.leftNum)) {
					bRightMatch = true
					nMatchedEdges++
				}
			}
		}
		if nMatchedEdges == 2 {
			// Found a corner
			corners[tile.tileNum] = true
			cornerTotal += tile.tileNum
			fmt.Println("Found corner: ", tile.tileNum)
		}
	}
	fmt.Printf("Found %d corners with a total of %d", len(corners), cornerTotal)

}

func getCode(line string) int {

	// Create bit array of line to use as thumbprint for side
	thumbprint := 0
	for i, ch := range line {
		if ch == '#' {
			thumbprint += (2 ^ i)
		}
	}
	return thumbprint
}
