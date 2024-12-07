package year2024

import (
	"slices"
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

type PosDirPair struct {
	pos string
	dir string
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

		d.movePlayer(&pos, &dir)

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
	return d.part2_naive()
}

func (d *Day6) part2_naive() string {
	// determine initial position
	pos := d.getInitialPosiiton()
	dir := Pos{0, -1}

	sum := 0
	for row, line := range d.Input {
		for col, char := range line {
			// already a block or original starting position here
			if char == '#' || char == '^' {
				continue
			}

			// place a blockade, run simulation
			d.Input[row] = d.Input[row][:col] + "#" + d.Input[row][col+1:]
			if d.doesFindLoop(pos, dir) {
				sum++
			}
			d.Input[row] = d.Input[row][:col] + "." + d.Input[row][col+1:]
		}
	}

	return strconv.Itoa(sum)
}

func (d *Day6) getInitialPosiiton() Pos {
	pos := Pos{}
	for row, line := range d.Input {
		if col := strings.Index(line, "^"); col != -1 {
			pos.x = col
			pos.y = row
			break
		}
	}
	return pos
}

func (d *Day6) movePlayer(pos *Pos, dir *Pos) {
	// normal movement procedure
	if d.Input[pos.y+dir.y][pos.x+dir.x] == '#' {
		dir.TurnRight()
		return
	}

	pos.y += dir.y
	pos.x += dir.x
}

func (d *Day6) doesFindLoop(pos Pos, dir Pos) bool {
	width, height := len(d.Input[0]), len(d.Input)
	visitedPosDirs := make(map[string][]string)

	for {
		// No loop can be created since we are leaving the bounds of the map
		if pos.x+dir.x < 0 || pos.x+dir.x >= width || pos.y+dir.y < 0 || pos.y+dir.y >= height {
			break
		}

		// check if we've been here before at this direction
		dirs, ok := visitedPosDirs[pos.String()]
		if ok {
			if slices.Contains(dirs, dir.String()) {
				return true
			}
		}

		visitedPosDirs[pos.String()] = append(visitedPosDirs[pos.String()], dir.String())

		d.movePlayer(&pos, &dir)
	}

	return false
}
