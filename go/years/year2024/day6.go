package year2024

import (
	"slices"
	"strconv"
	"strings"
)

// ANSWERS:
//
// Part 1: 4758
//
// Part 2:
type Day6 struct {
	Input []string
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

// avg ~0.66ms per iter over 10k iters
func (d Day6) Part1() string {
	pos := d.getInitialPosiiton()
	dir := Pos{0, -1}

	path := d.walk(pos, dir)

	return strconv.Itoa(len(path))
}

func (d Day6) Part2() string {
	// determine initial position
	pos := d.getInitialPosiiton()
	dir := Pos{0, -1}
	path := d.walk(pos, dir)

	sum := 0
	for _, blockPos := range path {
		row := blockPos.y
		col := blockPos.x
		char := d.Input[row][col]
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

func (d *Day6) move(pos *Pos, dir *Pos) {
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

		d.move(&pos, &dir)
	}

	return false
}

func (d *Day6) walk(pos Pos, dir Pos) []Pos {
	width, height := len(d.Input[0]), len(d.Input)
	visited := map[string]bool{}
	path := []Pos{}

	for pos.x >= 0 && pos.x < width && pos.y >= 0 && pos.y < height {
		if !visited[pos.String()] {
			visited[pos.String()] = true
			path = append(path, pos)
		}

		if pos.x+dir.x < 0 || pos.x+dir.x >= width || pos.y+dir.y < 0 || pos.y+dir.y >= height {
			break
		}

		d.move(&pos, &dir)
	}
	return path
}
