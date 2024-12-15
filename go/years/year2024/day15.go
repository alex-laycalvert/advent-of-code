package year2024

import (
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 1487337
//
// Part 2:
type Day15 struct {
	Input []string
}

type (
	Point rune
	Move  rune
)

const (
	WALL  Point = '#'
	EMPTY Point = '.'
	BOX   Point = 'O'

	UP    Move = '^'
	DOWN  Move = 'v'
	LEFT  Move = '<'
	RIGHT Move = '>'
)

func (d Day15) Part1() string {
	grid := make(map[string]Point)
	moves := make([]Move, 0)

	var width, height int
	var robot utils.Pos
	parsingMoves := false

	for row, line := range d.Input {
		if parsingMoves {
			for _, c := range line {
				moves = append(moves, Move(c))
			}
			continue
		}

		if line == "" {
			parsingMoves = true
			height = row - 2
			continue
		}

		if row == 0 {
			width = len(line) - 1
			continue
		}

		for col, c := range line {
			if col == 0 || col == width {
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

	// for y := range height {
	// 	for x := range width {
	// 		if x == robot.X && y == robot.Y {
	// 			print("@")
	// 			continue
	// 		}
	// 		print(string(grid[utils.PosToString(x, y)]))
	// 	}
	// 	println()
	// }
	// println()

	for _, move := range moves {
		// println(string(move))
		switch move {
		case UP:
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
		case DOWN:
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
		case LEFT:
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
		case RIGHT:
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

		// for y := range height {
		// 	for x := range width {
		// 		if x == robot.X && y == robot.Y {
		// 			print("@")
		// 			continue
		// 		}
		// 		print(string(grid[utils.PosToString(x, y)]))
		// 	}
		// 	println()
		// }
		// println()
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
	return "Not Implemented"
}
