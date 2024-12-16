package year2024

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	// "time"

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
	WALL      Point = '#'
	EMPTY     Point = '.'
	BOX       Point = 'O'
	BOX_RIGHT Point = ']'
	BOX_LEFT  Point = '['

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

	printGrid(grid, width, height, robot, os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// for _, move := range moves {
		var move Move
		switch scanner.Text() {
		case "w":
			move = UP
			break
		case "s":
			move = DOWN
			break
		case "a":
			move = LEFT
			break
		case "d":
			move = RIGHT
			break
		default:
			break
		}
		fmt.Fprintln(os.Stdout, string(move))
		switch move {
		case UP:
			if robot.Y == 0 {
				continue
			}

			firstEmptyY := -1
			left := robot.X
			right := robot.X + 1
			rowBounds := make([][]int, 0)
			for y := robot.Y - 1; y >= 0; y-- {
				// check for wall
				foundWall := false
				for i := left; i < right; i++ {
					key := utils.PosToString(left, y)
					if grid[key] == WALL {
						foundWall = true
						break
					}
				}
				if foundWall {
					break
				}

				newLeft := -1
				newRight := right
				for i := left; i < right; i++ {
					key := utils.PosToString(i, y)
					if i == left && grid[key] == BOX_RIGHT {
						newLeft = left - 1
						break
					}

					if grid[key] != EMPTY {
						newLeft = i
						break
					}
				}

				if newLeft == -1 {
					firstEmptyY = y
					break
				}

				for i := right - 1; i >= left; i++ {
					key := utils.PosToString(i, y)
					if i == right-1 && grid[key] == BOX_LEFT {
						newRight = right + 1
						break
					}

					if grid[key] != EMPTY {
						right = i
						break
					}
				}

				left = newLeft
				right = newRight
				rowBounds = append(rowBounds, []int{left, right})
			}

			if firstEmptyY == -1 {
				break
			}

			if len(rowBounds) == 0 {
				robot.Y--
				break
			}

			for y := firstEmptyY; y < robot.Y; y++ {
				if y == robot.Y-1 {
					row := rowBounds[0]
					left := row[0]
					right := row[1]
					for i := left; i < right; i++ {
						key := utils.PosToString(i, y)
						grid[key] = EMPTY
					}
					break
				}

				nextRow := rowBounds[len(rowBounds)-1-(y-firstEmptyY)]
				left := nextRow[0]
				right := nextRow[1]
				for i := left; i < right; i++ {
					key := utils.PosToString(i, y)
					nextKey := utils.PosToString(i, y+1)
					grid[key] = grid[nextKey]
					grid[nextKey] = EMPTY
				}
			}

			robot.Y--
			break
		case DOWN:
			if robot.Y == height-1 {
				break
			}

			firstEmptyY := -1
			left := robot.X
			right := robot.X + 1
			rowBounds := make([][]int, 0)
			for y := robot.Y + 1; y < height; y++ {
				// check for wall
				foundWall := false
				for i := left; i < right; i++ {
					key := utils.PosToString(left, y)
					if grid[key] == WALL {
						foundWall = true
						break
					}
				}
				if foundWall {
					break
				}

				newLeft := -1
				newRight := right
				for i := left; i < right; i++ {
					key := utils.PosToString(i, y)
					if i == left && grid[key] == BOX_RIGHT {
						newLeft = left - 1
						break
					}

					if grid[key] != EMPTY {
						newLeft = i
						break
					}
				}

				if newLeft == -1 {
					firstEmptyY = y
					break
				}

				for i := right - 1; i >= left; i++ {
					key := utils.PosToString(i, y)
					if i == right-1 && grid[key] == BOX_LEFT {
						newRight = right + 1
						break
					}

					if grid[key] != EMPTY {
						right = i
						break
					}
				}

				left = newLeft
				right = newRight
				rowBounds = append(rowBounds, []int{left, right})
			}

			if firstEmptyY == -1 {
				break
			}

			if len(rowBounds) == 0 {
				robot.Y++
				break
			}

			for y := firstEmptyY; y > robot.Y; y-- {
				if y == robot.Y+1 {
					row := rowBounds[0]
					left := row[0]
					right := row[1]
					for i := left; i < right; i++ {
						key := utils.PosToString(i, y)
						grid[key] = EMPTY
					}
					break
				}

				nextRow := rowBounds[len(rowBounds)-1-(firstEmptyY-y)]
				left := nextRow[0]
				right := nextRow[1]
				for i := left; i < right; i++ {
					key := utils.PosToString(i, y)
					nextKey := utils.PosToString(i, y-1)
					grid[key] = grid[nextKey]
					grid[nextKey] = EMPTY
				}
			}

			robot.Y++
			break
		case LEFT:
			if robot.X == 0 {
				break
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
				break
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
				break
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
				break
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
		printGrid(grid, width, height, robot, os.Stdout)
	}

	answer := 0
	for y := range height {
		for x := range width {
			if grid[utils.PosToString(x, y)] == BOX_LEFT {
				answer += 100*(y+1) + x + 2
			}
		}
	}

	return strconv.Itoa(answer)
}

func printGrid(grid map[string]Point, width int, height int, robot utils.Pos, file *os.File) {
	for y := range height {
		for x := range width {
			if x == robot.X && y == robot.Y {
				fmt.Fprint(file, "@")
				continue
			}
			fmt.Fprint(file, string(grid[utils.PosToString(x, y)]))
		}
		fmt.Fprintln(file)
	}
	fmt.Fprintln(file)
}
