package year2024

import (
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
	d rune
}

var turnRightMap = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

func (p Pos) String(withDir bool) string {
	key := strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
	if withDir {
		key += string(p.d)
	}
	return key
}

func (p *Pos) getDirOffset() (x int, y int) {
	switch p.d {
	case '^':
		return 0, -1
	case 'v':
		return 0, 1
	case '>':
		return 1, 0
	case '<':
		return -1, 0
	}
	return 0, 0
}

func (p *Pos) NextPos() Pos {
	x, y := p.getDirOffset()
	return Pos{
		x: p.x + x,
		y: p.y + y,
		d: p.d,
	}
}

// avg ~0.66ms per iter over 10k iters
func (d Day6) Part1() string {
	pos := d.getInitialPosiiton()
	path, _ := d.walk(pos, false)
	return strconv.Itoa(len(path))
}

func (d Day6) Part2() string {
	return d.part2_bruteForce()
}

// avg ~2.4s
func (d *Day6) part2_bruteForce() string {
	// determine initial position
	pos := d.getInitialPosiiton()
	path, _ := d.walk(pos, false)

	sum := 0
	// we only need to iterate over the covered path, not every single
	// position in the grid
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
		if _, foundLoop := d.walk(pos, true); foundLoop {
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
			pos.d = '^'
			break
		}
	}
	return pos
}

func (d *Day6) move(pos *Pos) {
	nextPos := pos.NextPos()

	// normal movement procedure
	if d.Input[nextPos.y][nextPos.x] == '#' {
		pos.d = turnRightMap[pos.d]
		return
	}

	pos.y = nextPos.y
	pos.x = nextPos.x
}

func (d *Day6) walk(pos Pos, checkForLoop bool) (path []Pos, foundLoop bool) {
	width, height := len(d.Input[0]), len(d.Input)
	visited := map[string]bool{}
	path = []Pos{}
	foundLoop = false

	for pos.x >= 0 && pos.x < width && pos.y >= 0 && pos.y < height {
		key := pos.String(checkForLoop)
		if visited[key] {
			if checkForLoop {
				return path, true
			}
		} else {
			visited[key] = true
			path = append(path, pos)
		}
		nextPos := pos.NextPos()
		if nextPos.x < 0 || nextPos.x >= width || nextPos.y < 0 || nextPos.y >= height {
			break
		}
		d.move(&pos)
	}
	return path, foundLoop
}
