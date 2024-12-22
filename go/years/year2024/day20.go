package year2024

import (
	"slices"
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils/search"
)

// ANSWERS:
//
// Part 1: 1378
//
// Part 2: 975379
type Day20 struct {
	Input []string
}

func (d Day20) Part1() string {
	maze := search.NewMaze(d.Input)
	searcher := search.NewAStar()
	path := make(search.Points, 0)
	for iter := range searcher.Search(*maze) {
		path = iter.CurrentPath
	}
	slices.Reverse(path)
	cheatSavings := 100
	cheatDistance := 2
	answer := findCheats(path, cheatSavings, cheatDistance)
	return strconv.Itoa(answer)
}

func (d Day20) Part2() string {
	maze := search.NewMaze(d.Input)
	searcher := search.NewAStar()
	path := make(search.Points, 0)
	for iter := range searcher.Search(*maze) {
		path = iter.CurrentPath
	}
	slices.Reverse(path)
	cheatSavings := 100
	cheatDistance := 20
	answer := findCheats(path, cheatSavings, cheatDistance)
	return strconv.Itoa(answer)
}

func findCheats(path search.Points, cheatSavings int, cheatDistance int) int {
	cheats := 0
	for i := 0; i < len(path); i++ {
		for j := i + cheatSavings + 1; j < len(path); j++ {
			distance := path[i].DistanceTo(path[j])
			if distance > cheatDistance {
				continue
			}
			if j-i-distance+1 < cheatSavings {
				continue
			}
			cheats++
		}
	}
	return cheats
}
