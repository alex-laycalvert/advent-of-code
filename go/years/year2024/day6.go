package year2024

import (
	"strconv"
)

// ANSWERS:
//
// Part 1: 4758
//
// Part 2: 1670
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

func (p Pos) String(withDirection bool) string {
	key := strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
	if withDirection {
		key += string(p.d)
	}
	return key
}

func (d Day6) Part1() string {
	return d.part1_bruteForceStructured()
}

// avg ~0.99ms per iter over 10k iters
func (d *Day6) part1_bruteForceStructured() string {
	game := NewGame(d.Input)
	path, _ := game.Walk(false)
	return strconv.Itoa(len(path))
}

func (d Day6) Part2() string {
	return d.part2_bruteForceStructured()
}

// avg ~2.6s
func (d *Day6) part2_bruteForceStructured() string {
	game := NewGame(d.Input)
	path, _ := game.Walk(false)

	sum := 0
	// we only need to iterate over the covered path, not every single
	// position in the grid
	for _, blockPos := range path {
		if !game.SetObstacle(blockPos.x, blockPos.y) {
			continue
		}
		if _, foundLoop := game.Walk(true); foundLoop {
			sum++
		}
		game.RemoveObstacle(blockPos.x, blockPos.y)
		game.Reset()
	}

	return strconv.Itoa(sum)
}

type Node struct {
	IsObstacle bool
}

type Game struct {
	nodeMap  [][]Node
	width    int
	height   int
	pos      Pos
	startPos Pos
}

func NewGame(input []string) *Game {
	nodeMap := make([][]Node, len(input))

	startingPos := Pos{}
	for i, line := range input {
		nodeMap[i] = make([]Node, len(line))
		for j, char := range line {
			if char == '#' {
				nodeMap[i][j].IsObstacle = true
				continue
			}

			if char == '^' {
				startingPos = Pos{j, i, char}
				continue
			}
		}
	}

	width, height := len(input[0]), len(input)

	return &Game{nodeMap, width, height, startingPos, startingPos}
}

func (g *Game) Walk(checkForLoop bool) (path []Pos, foundLoop bool) {
	visited := map[string]bool{}
	path = []Pos{}
	foundLoop = false

	for {
		_, ok := visited[g.pos.String(checkForLoop)]
		if ok && checkForLoop {
			foundLoop = true
			break
		}

		if !ok {
			path = append(path, g.pos)
			visited[g.pos.String(checkForLoop)] = true
		}

		if g.pos.x < 0 || g.pos.x >= g.width || g.pos.y < 0 || g.pos.y >= g.height {
			break
		}

		x, y := g.getDirectionOffsets()

		if g.pos.x+x < 0 || g.pos.x+x >= g.width || g.pos.y+y < 0 || g.pos.y+y >= g.height {
			break
		}

		if g.nodeMap[g.pos.y+y][g.pos.x+x].IsObstacle {
			g.pos.d = turnRightMap[g.pos.d]
		} else {
			g.pos.x += x
			g.pos.y += y
		}
	}

	return path, foundLoop
}

func (g *Game) SetObstacle(x int, y int) bool {
	if x == g.startPos.x && y == g.startPos.y {
		return false
	}

	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return false
	}

	if g.nodeMap[y][x].IsObstacle {
		return false
	}

	g.nodeMap[y][x].IsObstacle = true
	return true
}

func (g *Game) RemoveObstacle(x int, y int) {
	if x == g.startPos.x && y == g.startPos.y {
		return
	}

	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return
	}

	g.nodeMap[y][x].IsObstacle = false
}

func (g *Game) Reset() {
	g.pos.x = g.startPos.x
	g.pos.y = g.startPos.y
	g.pos.d = g.startPos.d
}

func (g *Game) getDirectionOffsets() (x int, y int) {
	switch g.pos.d {
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
