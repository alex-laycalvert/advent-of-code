package year2024

import (
	"fmt"
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 99488
//
// Part 2: 516
type Day16 struct {
	Input []string
}

func (d Day16) Part1() string {
	maze := NewMaze(d.Input)
	paths := maze.GetPaths(maze.Start, maze.End)

	bestPath := MazePath{Score: -1}
	for _, path := range paths {
		if bestPath.Score == -1 || path.Score < bestPath.Score {
			bestPath = path
		}
	}

	return strconv.Itoa(bestPath.Score)
}

func (d Day16) Part2() string {
	maze := NewMaze(d.Input)
	paths := maze.GetPaths(maze.Start, maze.End)

	bestPath := MazePath{Score: -1}
	for _, path := range paths {
		if bestPath.Score == -1 || path.Score < bestPath.Score {
			bestPath = path
		}
	}

	uniqueTiles := make([]utils.Pos, 0)
	for _, path := range paths {
		if path.Score > bestPath.Score {
			continue
		}
		for _, pos := range path.Path {
			if pathContains(uniqueTiles, pos) {
				continue
			}
			uniqueTiles = append(uniqueTiles, pos)
		}
	}

	return strconv.Itoa(len(uniqueTiles))
}

type Maze struct {
	maze [][]Point

	Start utils.Pos
	End   utils.Pos
}

type MazePath struct {
	pos utils.Pos

	Score int
	Path  []utils.Pos
}

func NewMaze(input []string) *Maze {
	maze := make([][]Point, len(input))

	var start, end utils.Pos
	for i, line := range input {
		maze[i] = make([]Point, len(line))
		for j, c := range line {
			if c == 'S' {
				start = utils.Pos{X: j, Y: i, Ch: '>'}
				maze[i][j] = EMPTY
				continue
			}

			if c == 'E' {
				end = utils.Pos{X: j, Y: i}
				maze[i][j] = EMPTY
				continue
			}

			maze[i][j] = Point(c)
		}
	}

	return &Maze{
		maze:  maze,
		Start: start,
		End:   end,
	}
}

func (m Maze) GetPaths(start utils.Pos, end utils.Pos) []MazePath {
	unvisitedSet := make(map[string]utils.Pos)
	for i, row := range m.maze {
		for j, pt := range row {
			if pt == EMPTY {
				unvisitedSet[utils.Pos{X: j, Y: i}.String(false)] = utils.Pos{X: j, Y: i}
			}
		}
	}

	pq := utils.NewPriorityQueue[MazePath]()
	pq.Push(MazePath{
		pos:   start,
		Score: 0,
		Path:  []utils.Pos{start},
	})

	bestScore := 0
	visited := make(map[string]int)
	paths := make([]MazePath, 0)
	for pq.Len() > 0 {
		current, _ := pq.Pop()

		if current.pos.Equals(end, false) {
			paths = append(paths, current)
			if bestScore == 0 || current.Score < bestScore {
				bestScore = current.Score
			}
			continue
		}

		key := current.pos.String(false)
		if s, ok := visited[key]; ok && current.Score > s+1000 {
			continue
		}
		visited[key] = current.Score

		for i := range 3 {
			nextPos := current.pos
			score := 1
			if i == 1 {
				nextPos = nextPos.TurnLeft()
				score += 1000
			} else if i == 2 {
				nextPos = nextPos.TurnRight()
				score += 1000
			}
			nextPos = nextPos.MoveForward()
			if m.maze[nextPos.Y][nextPos.X] != EMPTY {
				continue
			}
			newPath := make([]utils.Pos, len(current.Path))
			copy(newPath, current.Path)
			newPath = append(newPath, nextPos)
			pq.Push(MazePath{
				pos:   nextPos,
				Score: current.Score + score,
				Path:  newPath,
			})
		}
	}

	return paths
}

func (m Maze) Print(path []utils.Pos) {
	for i, row := range m.maze {
		for j, pt := range row {
			printed := false
			for _, pos := range path {
				if pos.X == j && pos.Y == i {
					fmt.Print(string(pos.Ch))
					printed = true
					break
				}
			}

			if printed {
				continue
			}

			if m.End.X == j && m.End.Y == i {
				fmt.Print("E")
				continue
			}

			fmt.Print(string(pt))
		}
		fmt.Println()
	}
}

func (p MazePath) Less(q MazePath) bool {
	return p.Score < q.Score
}

func (p MazePath) Print() {
	for i, pos := range p.Path {
		fmt.Print(pos.String(true))
		if i != len(p.Path)-1 {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func pathContains(path []utils.Pos, pos utils.Pos) bool {
	for _, p := range path {
		if p.Equals(pos, false) {
			return true
		}
	}

	return false
}
