package year2024

import (
	"strconv"
)

// ANSWERS:
//
// Part 1: 324
//
// Part 2: 575227823167869
type Day19 struct {
	Input []string
}

func (d Day19) Part1() string {
	answer := 0
	towelMap := parseTowelMap(d.Input[0])
	for i, line := range d.Input {
		if i < 2 {
			continue
		}
		if isTowelPossible(line, towelMap, make(map[string]int)) > 0 {
			answer++
		}
	}
	return strconv.Itoa(answer)
}

func (d Day19) Part2() string {
	answer := 0
	towelMap := parseTowelMap(d.Input[0])
	for i, line := range d.Input {
		if i < 2 {
			continue
		}
		answer += isTowelPossible(line, towelMap, make(map[string]int))
	}
	return strconv.Itoa(answer)
}

func parseTowelMap(line string) map[string]bool {
	towelMap := make(map[string]bool)
	curr := ""
	for _, ch := range line {
		if ch == ',' || ch == ' ' {
			towelMap[curr] = true
			curr = ""
			continue
		}
		curr += string(ch)
	}
	towelMap[curr] = true
	return towelMap
}

func isTowelPossible(towel string, towelMap map[string]bool, cache map[string]int) int {
	if ways, ok := cache[towel]; ok {
		return ways
	}
	if len(towel) == 0 {
		return 1
	}
	numWays := 0
	for i := 0; i < len(towel); i++ {
		if !towelMap[towel[:i+1]] {
			continue
		}
		numWays += isTowelPossible(towel[i+1:], towelMap, cache)
	}
	cache[towel] = numWays
	return numWays
}
