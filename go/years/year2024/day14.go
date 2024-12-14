package year2024

import (
	"os"
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 229632480
//
// Part 2: 7051
type Day14 struct {
	Input []string
}

func (d Day14) Part1() string {
	var width, height int
	var err error

	guards := make([]*Guard, len(d.Input)-1)
	for i, line := range d.Input {
		if i == 0 {
			parts := strings.Split(line, ",")
			width, err = strconv.Atoi(parts[0])
			utils.PanicErr(err)
			height, err = strconv.Atoi(parts[1])
			utils.PanicErr(err)
			continue
		}

		guards[i-1] = NewGuard(line)
	}

	simulateGuards(guards, 100, width, height, false)

	topLeft := 0
	topRight := 0
	bottomLeft := 0
	bottomRight := 0
	for _, guard := range guards {
		if guard.Pos.X < width/2 && guard.Pos.Y < height/2 {
			topLeft++
			continue
		}

		if guard.Pos.X > width/2 && guard.Pos.Y < height/2 {
			topRight++
			continue
		}

		if guard.Pos.X < width/2 && guard.Pos.Y > height/2 {
			bottomLeft++
			continue
		}

		if guard.Pos.X > width/2 && guard.Pos.Y > height/2 {
			bottomRight++
			continue
		}
	}

	if topRight == 0 {
		topRight = 1
	}

	if topLeft == 0 {
		topLeft = 1
	}

	if bottomLeft == 0 {
		bottomLeft = 1
	}

	if bottomRight == 0 {
		bottomRight = 1
	}

	answer := topRight * bottomLeft * bottomRight * topLeft

	return strconv.Itoa(answer)
}

func (d Day14) Part2() string {
	var width, height int
	var err error

	guards := make([]*Guard, len(d.Input)-1)
	for i, line := range d.Input {
		if i == 0 {
			parts := strings.Split(line, ",")
			width, err = strconv.Atoi(parts[0])
			utils.PanicErr(err)
			height, err = strconv.Atoi(parts[1])
			utils.PanicErr(err)
			continue
		}

		guards[i-1] = NewGuard(line)
	}

	simulateGuards(guards, 10_000, width, height, true)

	answer := 0
	return strconv.Itoa(answer)
}

type Guard struct {
	Pos utils.Pos
	Vel utils.Pos
}

func NewGuard(line string) *Guard {
	// left  : p=0,4
	// right : v=3,-1
	parts := strings.Split(line, " ")
	point := utils.ParseNumbers(parts[0])
	velocity := utils.ParseNumbers(parts[1])
	return &Guard{
		Pos: utils.Pos{
			X: point[0],
			Y: point[1],
		},
		Vel: utils.Pos{
			X: velocity[0],
			Y: velocity[1],
		},
	}
}

func simulateGuards(guards []*Guard, iterations int, width int, height int, showSimulation bool) {
	var file *os.File
	var err error

	if showSimulation {
		file, err = os.Create("output.txt")
		utils.PanicErr(err)
		defer file.Close()
	}

	for range iterations {
		if showSimulation {
			printGuards(guards, width, height, file)
		}
		for _, guard := range guards {
			guard.Pos.X += guard.Vel.X
			guard.Pos.Y += guard.Vel.Y

			if guard.Pos.X >= width {
				guard.Pos.X -= width
			}

			if guard.Pos.X < 0 {
				guard.Pos.X += width
			}

			if guard.Pos.Y >= height {
				guard.Pos.Y -= height
			}

			if guard.Pos.Y < 0 {
				guard.Pos.Y += height
			}
		}
	}

	if showSimulation {
		printGuards(guards, width, height, file)
	}
}

func printGuards(guards []*Guard, width int, height int, file *os.File) {
	_, err := file.WriteString("\n")
	utils.PanicErr(err)
	for r := range height {
		str := ""
		for c := range width {
			numGuards := 0
			for _, guard := range guards {
				if guard.Pos.X == c && guard.Pos.Y == r {
					numGuards++
				}
			}
			if numGuards == 0 {
				str += "."
			} else {
				str += "#"
			}
		}
		_, err := file.WriteString(str + "\n")
		utils.PanicErr(err)
	}
}
