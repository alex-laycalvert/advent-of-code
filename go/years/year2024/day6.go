package year2024

import (
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 4758
//
// Part 2: 1670
type Day6 struct {
	Input []string
}

var turnRightMap = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

func (d Day6) Part1() string {
	return d.part1_bruteForceStructured()
}

// avg ~0.99ms per iter over 10k iters
func (d *Day6) part1_bruteForceStructured() string {
	game := NewGame_day6(d.Input)
	path, _ := game.Walk(false)
	return strconv.Itoa(len(path))
}

func (d Day6) Part2() string {
	return d.part2_bruteForceStructured()
}

// avg ~2.6s
func (d *Day6) part2_bruteForceStructured() string {
	game := NewGame_day6(d.Input)
	path, _ := game.Walk(false)

	sum := 0
	// we only need to iterate over the covered path, not every single
	// position in the grid
	for _, blockPos := range path {
		if !game.SetObstacle(blockPos.X, blockPos.Y) {
			continue
		}
		if _, foundLoop := game.Walk(true); foundLoop {
			sum++
		}
		game.RemoveObstacle(blockPos.X, blockPos.Y)
		game.Reset()
	}

	return strconv.Itoa(sum)
}

type Node_day6 struct {
	IsObstacle bool
}

type Game_day6 struct {
	nodeMap  [][]Node_day6
	width    int
	height   int
	pos      utils.Pos
	startPos utils.Pos
}

func NewGame_day6(input []string) *Game_day6 {
	nodeMap := make([][]Node_day6, len(input))

	startingPos := utils.Pos{}
	for i, line := range input {
		nodeMap[i] = make([]Node_day6, len(line))
		for j, char := range line {
			if char == '#' {
				nodeMap[i][j].IsObstacle = true
				continue
			}

			if char == '^' {
				startingPos = utils.Pos{X: j, Y: i, Ch: char}
				continue
			}
		}
	}

	width, height := len(input[0]), len(input)

	return &Game_day6{nodeMap, width, height, startingPos, startingPos}
}

func (g *Game_day6) Walk(checkForLoop bool) (path []utils.Pos, foundLoop bool) {
	visited := map[string]bool{}
	path = []utils.Pos{}
	foundLoop = false

	for g.pos.X >= 0 && g.pos.X < g.width && g.pos.Y >= 0 && g.pos.Y < g.height {
		key := g.pos.String(checkForLoop)
		_, ok := visited[key]
		if ok && checkForLoop {
			foundLoop = true
			break
		}

		if !ok {
			path = append(path, g.pos)
			visited[key] = true
		}

		dx, dy := g.getDirectionOffsets()
		if g.pos.X+dx < 0 || g.pos.X+dx >= g.width || g.pos.Y+dy < 0 || g.pos.Y+dy >= g.height {
			break
		}

		if g.nodeMap[g.pos.Y+dy][g.pos.X+dx].IsObstacle {
			g.pos.Ch = turnRightMap[g.pos.Ch]
		} else {
			g.pos.X += dx
			g.pos.Y += dy
		}
	}

	return path, foundLoop
}

func (g *Game_day6) SetObstacle(x int, y int) bool {
	if x == g.startPos.X && y == g.startPos.Y {
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

func (g *Game_day6) RemoveObstacle(x int, y int) {
	if x == g.startPos.X && y == g.startPos.Y {
		return
	}

	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return
	}

	g.nodeMap[y][x].IsObstacle = false
}

func (g *Game_day6) Reset() {
	g.pos.X = g.startPos.X
	g.pos.Y = g.startPos.Y
	g.pos.Ch = g.startPos.Ch
}

func (g *Game_day6) getDirectionOffsets() (x int, y int) {
	switch g.pos.Ch {
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
