package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"math"
	"strconv"
)

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day12/data.txt")

	// Action N means to move north by the given value.
	// Action S means to move south by the given value.
	// Action E means to move east by the given value.
	// Action W means to move west by the given value.
	// Action L means to turn left the given number of degrees.
	// Action R means to turn right the given number of degrees.
	// Action F means to move forward by the given value in the direction the ship is currently facing.

	x, y, head := 0, 0, 90
	for _, line := range lines {
		cmd := line[0]
		num, _ := strconv.Atoi(line[1:])

		// translate forward into proper dir
		if cmd == 'F' {
			switch head {
			case 0:
				cmd = 'N'
			case 90:
				cmd = 'E'
			case 180:
				cmd = 'S'
			case 270:
				cmd = 'W'
			}
		}

		switch cmd {
		case 'N':
			y -= num
		case 'S':
			y += num
		case 'E':
			x += num
		case 'W':
			x -= num
		case 'R':
			head += num
		case 'L':
			head -= num
		}

		// fix heading
		for {
			if head >= 360 {
				head -= 360
			} else {
				break
			}
		}
		for {
			if head < 0 {
				head += 360
			} else {
				break
			}
		}

		fmt.Printf("Line %s %d %d x:%d y:%d\n", string(cmd), num, head, x, y)
	}

	fmt.Println("Pos ", x, y, math.Abs(float64(x))+math.Abs(float64(y)))
}
