package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
)

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day12/data.txt")

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
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

		//fmt.Printf("Line %s %d %d x:%d y:%d\n", string(cmd), num, head, x, y)
	}

	fmt.Println("Pt 1 Pos ", x, y, Abs(x)+Abs(y))
}

func partTwo(lines []string) {

	// Action N means to move the waypoint north by the given value.
	// Action S means to move the waypoint south by the given value.
	// Action E means to move the waypoint east by the given value.
	// Action W means to move the waypoint west by the given value.
	// Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
	// Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
	// Action F means to move forward to the waypoint a number of times equal to the given value.

	x, y, wpX, wpY := 0, 0, 10, -1
	for _, line := range lines {
		cmd := line[0]
		num, _ := strconv.Atoi(line[1:])

		switch cmd {
		case 'N':
			wpY -= num
		case 'S':
			wpY += num
		case 'E':
			wpX += num
		case 'W':
			wpX -= num
		case 'F':
			// Move ship and waypoint
			x += wpX * num
			y += wpY * num
		case 'R', 'L':
			// Rotate waypoint around ship
			if (cmd == 'R' && num == 90) || (cmd == 'L' && num == 270) {
				tmp := wpX
				wpX = wpY * -1
				wpY = tmp
			} else if num == 180 {
				wpX = wpX * -1
				wpY = wpY * -1
			} else { // R 270 or L 90
				tmp := wpY
				wpY = wpX * -1
				wpX = tmp
			}
		}

		//fmt.Printf("Line %s %d x:%d y:%d wpX:%d wpY:%d\n", string(cmd), num, x, y, wpX, wpY)
	}

	fmt.Println("Pt 2 Pos ", x, y, Abs(x)+Abs(y))
}

// Abs - return absolute value the easy way
func Abs(num int) int {

	if num < 0 {
		return (-1 * num)
	}
	return num

}
