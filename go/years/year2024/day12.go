package year2024

import (
	"slices"
	"strconv"

	"github.com/zyedidia/generic/queue"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 1452678
//
// Part 2: 873584
type Day12 struct {
	Input []string
}

func (d Day12) Part1() string {
	answer := 0

	visited := map[string]bool{}
	for row, line := range d.Input {
		for col, char := range line {
			initialPos := utils.Pos{X: col, Y: row}
			if visited[initialPos.String(false)] {
				continue
			}
			ch := byte(char)
			perimeter := 0
			area := 0
			q := queue.New[utils.Pos]()
			q.Enqueue(initialPos)
			for !q.Empty() {
				pos := q.Dequeue()
				if visited[pos.String(false)] {
					continue
				}
				visited[pos.String(false)] = true

				area++
				neighbors := 0
				if pos.Y > 0 && d.Input[pos.Y-1][pos.X] == ch {
					neighbors++
					q.Enqueue(utils.Pos{X: pos.X, Y: pos.Y - 1})
				}

				if pos.Y < len(d.Input)-1 && d.Input[pos.Y+1][pos.X] == ch {
					neighbors++
					q.Enqueue(utils.Pos{X: pos.X, Y: pos.Y + 1})
				}

				if pos.X > 0 && d.Input[pos.Y][pos.X-1] == ch {
					neighbors++
					q.Enqueue(utils.Pos{X: pos.X - 1, Y: pos.Y})
				}

				if pos.X < len(line)-1 && d.Input[pos.Y][pos.X+1] == ch {
					neighbors++
					q.Enqueue(utils.Pos{X: pos.X + 1, Y: pos.Y})
				}
				perimeter += 4 - neighbors
			}

			answer += area * perimeter
		}
	}

	return strconv.Itoa(answer)
}

func (d Day12) Part2() string {
	answer := 0

	visited := map[string]bool{}
	for row, line := range d.Input {
		for col, char := range line {
			initialPos := utils.Pos{X: col, Y: row}
			if visited[initialPos.String(false)] {
				continue
			}
			ch := byte(char)
			area := 0
			q := queue.New[utils.Pos]()
			q.Enqueue(initialPos)

			type Face struct {
				face  rune
				axis  int
				index int
			}

			faces := make([]Face, 0)
			for !q.Empty() {
				pos := q.Dequeue()
				if visited[pos.String(false)] {
					continue
				}
				visited[pos.String(false)] = true

				area++
				neighbors := 0
				if pos.Y > 0 && d.Input[pos.Y-1][pos.X] == ch {
					neighbors++
					q.Enqueue(utils.Pos{X: pos.X, Y: pos.Y - 1})
				} else {
					faces = append(faces, Face{face: 'u', axis: pos.Y, index: pos.X})
				}

				if pos.Y < len(d.Input)-1 && d.Input[pos.Y+1][pos.X] == ch {
					neighbors++
					q.Enqueue(utils.Pos{X: pos.X, Y: pos.Y + 1})
				} else {
					faces = append(faces, Face{face: 'd', axis: pos.Y, index: pos.X})
				}

				if pos.X > 0 && d.Input[pos.Y][pos.X-1] == ch {
					neighbors++
					q.Enqueue(utils.Pos{X: pos.X - 1, Y: pos.Y})
				} else {
					faces = append(faces, Face{face: 'l', axis: pos.X, index: pos.Y})
				}

				if pos.X < len(line)-1 && d.Input[pos.Y][pos.X+1] == ch {
					neighbors++
					q.Enqueue(utils.Pos{X: pos.X + 1, Y: pos.Y})
				} else {
					faces = append(faces, Face{face: 'r', axis: pos.X, index: pos.Y})
				}
			}

			numSides := 0

			for _, dir := range []rune{'u', 'd', 'l', 'r'} {
				sameFace := make([]Face, 0)
				for _, face := range faces {
					if dir == face.face {
						sameFace = append(sameFace, face)
					}
				}

				slices.SortFunc(sameFace, func(a Face, b Face) int {
					if a.axis < b.axis {
						return -1
					}
					if a.axis > b.axis {
						return 1
					}
					if a.index < b.index {
						return -1
					}
					if a.index > b.index {
						return 1
					}
					return 0
				})

				for i, face := range sameFace {
					if i == 0 || sameFace[i-1].axis != face.axis || sameFace[i-1].index != face.index-1 {
						numSides++
					}
				}
			}

			answer += area * numSides
		}
	}

	return strconv.Itoa(answer)
}
