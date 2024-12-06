package year2024

import (
	"strconv"
	"strings"
	"time"
)

// ANSWERS:
//
// Part 1: 4758
//
// Part 2:
type Day6 struct {
	Input          []string
	ShowSimulation bool
}

type Pos struct {
	x int
	y int
}

func (p Pos) String() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func (p *Pos) TurnRight() {
	p.x, p.y = -p.y, p.x
}

func (p *Pos) TurnLeft() {
	p.x, p.y = p.y, -p.x
}

// avg ~0.66ms per iter over 10k iters
func (d Day6) Part1() string {
	pos := d.getInitialPosiiton()
	dir := Pos{0, -1}
	width, height := len(d.Input[0]), len(d.Input)
	visitedPositions := map[string]bool{}
	path := 0
	for pos.x >= 0 && pos.x < width && pos.y >= 0 && pos.y < height {
		if d.ShowSimulation {
			d.Input[pos.y] = d.Input[pos.y][:pos.x] + "X" + d.Input[pos.y][pos.x+1:]
		}
		if !visitedPositions[pos.String()] {
			visitedPositions[pos.String()] = true
			path++
		}

		if pos.x+dir.x < 0 || pos.x+dir.x >= width || pos.y+dir.y < 0 || pos.y+dir.y >= height {
			break
		}

		// turn right
		if d.Input[pos.y+dir.y][pos.x+dir.x] == '#' {
			dir.TurnRight()
		}
		pos.y += dir.y
		pos.x += dir.x

		if d.ShowSimulation {
			for _, line := range d.Input {
				println(line)
			}
			time.Sleep(1 * time.Second)
		}
	}

	return strconv.Itoa(path)
}

func (d Day6) Part2() string {
	// determine initial position
	pos := d.getInitialPosiiton()
	dir := Pos{0, -1}
	width, height := len(d.Input[0]), len(d.Input)

	visitedPositions := map[string]string{}
	sum := 0
	for pos.x >= 0 && pos.x < width && pos.y >= 0 && pos.y < height {
		dir.TurnRight()
		pos.y += dir.y
		pos.x += dir.x
		visitedDir, ok := visitedPositions[pos.String()]
		if ok && dir.String() == visitedDir {
			sum++
		}
		pos.y -= dir.y
		pos.x -= dir.x
		dir.TurnLeft()

		visitedPositions[pos.String()] = dir.String()

		if pos.x+dir.x < 0 || pos.x+dir.x >= width || pos.y+dir.y < 0 || pos.y+dir.y >= height {
			break
		}

		// turn right
		if d.Input[pos.y+dir.y][pos.x+dir.x] == '#' {
			dir.TurnRight()
		}
		pos.y += dir.y
		pos.x += dir.x

		if d.ShowSimulation {
			for _, line := range d.Input {
				println(line)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

	return strconv.Itoa(sum)
}

func (d *Day6) getInitialPosiiton() Pos {
	pos := Pos{}
	for row, line := range d.Input {
		if i := strings.Index(line, "^"); i != -1 {
			pos.x = i
			pos.y = row
			break
		}
	}
	return pos
}
