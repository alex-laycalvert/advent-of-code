package year2024

import (
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
	"github.com/alex-laycalvert/advent-of-code/utils/search"
)

// ANSWERS:
//
// Part 1: 506
//
// Part 2: 62,6
type Day18 struct {
	Input []string
}

func (d Day18) Part1() string {
	size := 71
	maxBytes := 1024
	input := make([]string, size)
	for i := range size {
		row := ""
		for range size {
			row += "."
		}
		input[i] = row
	}

	for i, line := range d.Input {
		if i >= maxBytes {
			break
		}
		parts := utils.ParseNumbers(line)
		x, y := parts[0], parts[1]
		input[y] = input[y][:x] + "#" + input[y][x+1:]
	}

	maze := search.NewMaze(input)
	maze.Start = search.Point{Row: 0, Col: 0}
	maze.End = search.Point{Row: size - 1, Col: size - 1}

	searcher := search.NewDijkstra()

	bestPath := make(search.Points, 0)
	for path := range searcher.Search(*maze) {
		bestPath = path.CurrentPath
	}

	return strconv.Itoa(len(bestPath) - 1)
}

func (d Day18) Part2() string {
	size := 71

	searcher := search.NewDijkstra()
	for maxBytes := len(d.Input) - 1; maxBytes >= 0; maxBytes-- {
		input := make([]string, size)
		for i := range size {
			row := ""
			for range size {
				row += "."
			}
			input[i] = row
		}

		for i, line := range d.Input {
			if i >= maxBytes {
				break
			}
			parts := utils.ParseNumbers(line)
			x, y := parts[0], parts[1]
			input[y] = input[y][:x] + "#" + input[y][x+1:]
		}

		maze := search.NewMaze(input)
		maze.Start = search.Point{Row: 0, Col: 0}
		maze.End = search.Point{Row: size - 1, Col: size - 1}
		bestPath := make(search.Points, 0)
		for path := range searcher.Search(*maze) {
			bestPath = path.CurrentPath
		}
		if bestPath[0].Equals(maze.End) {
			return d.Input[maxBytes]
		}
	}

	return strconv.Itoa(0)
}
