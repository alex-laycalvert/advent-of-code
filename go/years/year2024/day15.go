package year2024

import (
	"fmt"
	"io"
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 1487337
//
// Part 2: 1521952
type Day15 struct {
	Input []string
}

type (
	Point rune
)

const (
	WALL      Point = '#'
	EMPTY     Point = '.'
	BOX       Point = 'O'
	BOX_RIGHT Point = ']'
	BOX_LEFT  Point = '['
)

func (d Day15) Part1() string {
	grid := make(map[string]Point)
	moves := make([]utils.Move, 0)

	var width, height int
	var robot utils.Pos
	parsingMoves := false

	for row, line := range d.Input {
		if parsingMoves {
			for _, c := range line {
				moves = append(moves, utils.Move(c))
			}
			continue
		}

		if line == "" {
			parsingMoves = true
			height = row - 2
			continue
		}

		if row == 0 {
			width = len(line) - 2
			continue
		}

		for col, c := range line {
			if col == 0 || col > width {
				continue
			}
			pos := utils.PosToString(col-1, row-1)
			switch c {
			case '.':
				grid[pos] = EMPTY
				break
			case 'O':
				grid[pos] = BOX
				break
			case '@':
				grid[pos] = EMPTY
				robot = utils.Pos{X: col - 1, Y: row - 1}
				break
			case '#':
				grid[pos] = WALL
				break
			default:
				break
			}
		}
	}

	for _, move := range moves {
		switch move {
		case utils.UP:
			if robot.Y == 0 {
				continue
			}

			firstEmptyY := -1
			for y := robot.Y - 1; y >= 0; y-- {
				key := utils.PosToString(robot.X, y)
				if grid[key] == WALL {
					break
				}
				if grid[key] == EMPTY {
					firstEmptyY = y
					break
				}
			}
			if firstEmptyY == -1 {
				continue
			}

			currSpot := EMPTY
			for y := robot.Y - 1; y >= firstEmptyY; y-- {
				key := utils.PosToString(robot.X, y)
				tmp := currSpot
				currSpot = grid[key]
				grid[key] = tmp
			}

			robot.Y--
			break
		case utils.DOWN:
			if robot.Y == height-1 {
				continue
			}

			firstEmptyY := -1
			for y := robot.Y + 1; y < height; y++ {
				key := utils.PosToString(robot.X, y)
				if grid[key] == WALL {
					break
				}
				if grid[key] == EMPTY {
					firstEmptyY = y
					break
				}
			}
			if firstEmptyY == -1 {
				continue
			}

			currSpot := EMPTY
			for y := robot.Y + 1; y <= firstEmptyY; y++ {
				key := utils.PosToString(robot.X, y)
				tmp := currSpot
				currSpot = grid[key]
				grid[key] = tmp
			}

			robot.Y++
			break
		case utils.LEFT:
			if robot.X == 0 {
				continue
			}

			firstEmptyX := -1
			for x := robot.X - 1; x >= 0; x-- {
				key := utils.PosToString(x, robot.Y)
				if grid[key] == WALL {
					break
				}
				if grid[key] == EMPTY {
					firstEmptyX = x
					break
				}
			}
			if firstEmptyX == -1 {
				continue
			}

			currSpot := EMPTY
			for x := robot.X - 1; x >= firstEmptyX; x-- {
				key := utils.PosToString(x, robot.Y)
				tmp := currSpot
				currSpot = grid[key]
				grid[key] = tmp
			}

			robot.X--
			break
		case utils.RIGHT:
			if robot.X == width-1 {
				continue
			}

			firstEmptyX := -1
			for x := robot.X + 1; x < width; x++ {
				key := utils.PosToString(x, robot.Y)
				if grid[key] == WALL {
					break
				}
				if grid[key] == EMPTY {
					firstEmptyX = x
					break
				}
			}
			if firstEmptyX == -1 {
				continue
			}

			currSpot := EMPTY
			for x := robot.X + 1; x <= firstEmptyX; x++ {
				key := utils.PosToString(x, robot.Y)
				tmp := currSpot
				currSpot = grid[key]
				grid[key] = tmp
			}

			robot.X++
			break
		default:
			break
		}
	}

	answer := 0
	for y := range height {
		for x := range width {
			if grid[utils.PosToString(x, y)] == BOX {
				answer += 100*(y+1) + x + 1
			}
		}
	}

	return strconv.Itoa(answer)
}

func (d Day15) Part2() string {
	grid, moves, robot := NewGrid(d.Input)

	// grid.Print(robot, os.Stdout)
	for _, move := range moves {
		// fmt.Fprintln(os.Stdout, string(move))
		robot = grid.Move(robot, move)
		// grid.Print(robot, os.Stdout)
	}

	answer := 0
	for y := range grid.Height {
		for x := range grid.Width {
			if grid.Map[utils.PosToString(x, y)] == BOX_LEFT {
				answer += 100*(y+1) + x + 2
			}
		}
	}

	return strconv.Itoa(answer)
}

type Grid struct {
	Map    map[string]Point
	Width  int
	Height int
}

func (grid *Grid) Move(pos utils.Pos, move utils.Move) utils.Pos {
	if !grid.IsPossible(pos, move) {
		return pos
	}

	currentKey := pos.String(false)
	currentPt := grid.Map[currentKey]
	switch move {
	case utils.UP:
		nextKey := utils.PosToString(pos.X, pos.Y-1)
		pt := grid.Map[nextKey]

		if pt == BOX_LEFT {
			grid.Move(utils.Pos{X: pos.X, Y: pos.Y - 1}, utils.UP)
			grid.Move(utils.Pos{X: pos.X + 1, Y: pos.Y - 1}, utils.UP)
		}

		if pt == BOX_RIGHT {
			grid.Move(utils.Pos{X: pos.X, Y: pos.Y - 1}, utils.UP)
			grid.Move(utils.Pos{X: pos.X - 1, Y: pos.Y - 1}, utils.UP)
		}

		grid.Map[nextKey] = currentPt
		grid.Map[currentKey] = EMPTY
		return utils.Pos{X: pos.X, Y: pos.Y - 1}
	case utils.DOWN:
		nextKey := utils.PosToString(pos.X, pos.Y+1)
		pt := grid.Map[nextKey]

		if pt == BOX_LEFT {
			grid.Move(utils.Pos{X: pos.X, Y: pos.Y + 1}, utils.DOWN)
			grid.Move(utils.Pos{X: pos.X + 1, Y: pos.Y + 1}, utils.DOWN)
		}

		if pt == BOX_RIGHT {
			grid.Move(utils.Pos{X: pos.X, Y: pos.Y + 1}, utils.DOWN)
			grid.Move(utils.Pos{X: pos.X - 1, Y: pos.Y + 1}, utils.DOWN)
		}

		grid.Map[nextKey] = currentPt
		grid.Map[currentKey] = EMPTY
		return utils.Pos{X: pos.X, Y: pos.Y + 1}
	case utils.LEFT:
		nextKey := utils.PosToString(pos.X-1, pos.Y)
		pt := grid.Map[nextKey]

		if pt != EMPTY {
			grid.Move(utils.Pos{X: pos.X - 1, Y: pos.Y}, utils.LEFT)
		}

		grid.Map[nextKey] = currentPt
		grid.Map[currentKey] = EMPTY
		return utils.Pos{X: pos.X - 1, Y: pos.Y}
	case utils.RIGHT:
		nextKey := utils.PosToString(pos.X+1, pos.Y)
		pt := grid.Map[nextKey]

		if pt != EMPTY {
			grid.Move(utils.Pos{X: pos.X + 1, Y: pos.Y}, utils.RIGHT)
		}

		grid.Map[nextKey] = currentPt
		grid.Map[currentKey] = EMPTY
		return utils.Pos{X: pos.X + 1, Y: pos.Y}
	default:
		return pos
	}
}

func (grid Grid) IsPossible(pos utils.Pos, move utils.Move) bool {
	switch move {
	case utils.UP:
		if pos.Y == 0 {
			return false
		}
		nextKey := utils.PosToString(pos.X, pos.Y-1)
		pt := grid.Map[nextKey]

		if pt == EMPTY {
			return true
		}

		if pt == WALL {
			return false
		}

		if pt == BOX_LEFT {
			return grid.IsPossible(utils.Pos{X: pos.X, Y: pos.Y - 1}, utils.UP) && grid.IsPossible(utils.Pos{X: pos.X + 1, Y: pos.Y - 1}, utils.UP)
		}

		// BOX_RIGHT
		return grid.IsPossible(utils.Pos{X: pos.X, Y: pos.Y - 1}, utils.UP) && grid.IsPossible(utils.Pos{X: pos.X - 1, Y: pos.Y - 1}, utils.UP)
	case utils.DOWN:
		if pos.Y == grid.Height-1 {
			return false
		}
		nextKey := utils.PosToString(pos.X, pos.Y+1)
		pt := grid.Map[nextKey]

		if pt == EMPTY {
			return true
		}

		if pt == WALL {
			return false
		}

		if pt == BOX_LEFT {
			return grid.IsPossible(utils.Pos{X: pos.X, Y: pos.Y + 1}, utils.DOWN) && grid.IsPossible(utils.Pos{X: pos.X + 1, Y: pos.Y + 1}, utils.DOWN)
		}

		// BOX_RIGHT
		return grid.IsPossible(utils.Pos{X: pos.X, Y: pos.Y + 1}, utils.DOWN) && grid.IsPossible(utils.Pos{X: pos.X - 1, Y: pos.Y + 1}, utils.DOWN)
	case utils.LEFT:
		if pos.X == 0 {
			return false
		}
		nextKey := utils.PosToString(pos.X-1, pos.Y)
		pt := grid.Map[nextKey]

		if pt == EMPTY {
			return true
		}

		if pt == WALL {
			return false
		}

		return grid.IsPossible(utils.Pos{X: pos.X - 1, Y: pos.Y}, utils.LEFT)
	case utils.RIGHT:
		if pos.X == grid.Width-1 {
			return false
		}
		nextKey := utils.PosToString(pos.X+1, pos.Y)
		pt := grid.Map[nextKey]

		if pt == EMPTY {
			return true
		}

		if pt == WALL {
			return false
		}

		return grid.IsPossible(utils.Pos{X: pos.X + 1, Y: pos.Y}, utils.RIGHT)
	default:
		return false
	}
}

func (grid *Grid) Print(robot utils.Pos, w io.Writer) {
	for y := range grid.Height {
		for x := range grid.Width {
			if x == robot.X && y == robot.Y {
				fmt.Fprint(w, "@")
				continue
			}
			fmt.Fprint(w, string(grid.Map[utils.PosToString(x, y)]))
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w)
}

func NewGrid(input []string) (Grid, []utils.Move, utils.Pos) {
	grid := make(map[string]Point)
	moves := make([]utils.Move, 0)

	var width, height int
	var robot utils.Pos
	parsingMoves := false
	for row, line := range input {
		if parsingMoves {
			for _, c := range line {
				moves = append(moves, utils.Move(c))
			}
			continue
		}

		if line == "" {
			parsingMoves = true
			height = row - 2
			continue
		}

		if row == 0 {
			width = 2 * (len(line) - 2)
			continue
		}

		for col := 0; col < width; col += 2 {
			actualCol := (col + 2) / 2
			c := line[actualCol]
			pos := utils.PosToString(col, row-1)
			nextPos := utils.PosToString(col+1, row-1)

			switch c {
			case '.':
				grid[pos] = EMPTY
				grid[nextPos] = EMPTY
				break
			case 'O':
				grid[pos] = BOX_LEFT
				grid[nextPos] = BOX_RIGHT
				break
			case '@':
				grid[pos] = EMPTY
				grid[nextPos] = EMPTY
				robot = utils.Pos{X: col, Y: row - 1}
				break
			case '#':
				grid[pos] = WALL
				grid[nextPos] = WALL
				break
			default:
				break
			}
		}
	}

	return Grid{
		Map:    grid,
		Width:  width,
		Height: height,
	}, moves, robot
}
