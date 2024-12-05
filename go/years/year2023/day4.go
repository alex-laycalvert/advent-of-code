package year2023

import (
	"slices"
	"strconv"
)

// ANSWERS:
//
// Part 1: 32001
//
// Part 2: 5037841
type Day4 struct {
	Input []string
}

func (d Day4) Part1() string {
	return d.part1_sliceIncludes()
	// return d.part1_mapOnEachLine()
	// return d.part1_globalMap()
}

// on my laptop, avg time per iter is ~0.41ms over 100k iters
// Each slice isn't more than maybe 8 items or so so it's not
// very resource intensive to just iterate over the slice.
func (d *Day4) part1_sliceIncludes() string {
	sum := 0

	firstNumIndex := 0
	if len(d.Input) < 10 {
		firstNumIndex = 8
	} else {
		firstNumIndex = 10
	}

	for _, line := range d.Input {
		winningNums, lastIndex := parseNumbersToSlice(line, firstNumIndex)
		playingNums, _ := parseNumbersToSlice(line, lastIndex+1)
		gamePoints := 0
		for _, num := range playingNums {
			if !slices.Contains(winningNums, num) {
				continue
			}
			if gamePoints == 0 {
				gamePoints = 1
				continue
			}
			gamePoints *= 2
		}
		sum += gamePoints
	}

	return strconv.Itoa(sum)
}

// ~0.5ms / iter over 100k iters
// Probably marginally slower because a map is being intialized on each
// line which is slower than a slice.
func (d *Day4) part1_mapOnEachLine() string {
	sum := 0

	firstNumIndex := 0
	if len(d.Input) < 10 {
		firstNumIndex = 8
	} else {
		firstNumIndex = 10
	}

	for _, line := range d.Input {
		winningNums, lastIndex := parseNumbersToNewMap(line, firstNumIndex)
		playingNums, _ := parseNumbersToSlice(line, lastIndex+1)
		gamePoints := 0
		for _, num := range playingNums {
			if _, ok := winningNums[num]; !ok {
				continue
			}
			if gamePoints == 0 {
				gamePoints = 1
				continue
			}
			gamePoints *= 2
		}
		sum += gamePoints
	}

	return strconv.Itoa(sum)
}

// ~0.43ms / iter over 100k iters
// Clearing the map on each call is probably the most expensive part
func (d *Day4) part1_globalMap() string {
	sum := 0

	firstNumIndex := 0
	if len(d.Input) < 10 {
		firstNumIndex = 8
	} else {
		firstNumIndex = 10
	}

	winningNums := make(map[int]int)
	for _, line := range d.Input {
		lastIndex := parseNumbersToExistingMap(line, firstNumIndex, winningNums)
		playingNums, _ := parseNumbersToSlice(line, lastIndex+1)
		gamePoints := 0
		for _, num := range playingNums {
			if _, ok := winningNums[num]; !ok {
				continue
			}
			if gamePoints == 0 {
				gamePoints = 1
				continue
			}
			gamePoints *= 2
		}
		sum += gamePoints
	}

	return strconv.Itoa(sum)
}

func (d Day4) Part2() string {
	sum := 0

	firstNumIndex := 0
	if len(d.Input) < 10 {
		firstNumIndex = 8
	} else {
		firstNumIndex = 10
	}

	copies := make(map[int]int)

	for i, line := range d.Input {
		winningNums, lastIndex := parseNumbersToSlice(line, firstNumIndex)
		playingNums, _ := parseNumbersToSlice(line, lastIndex+1)

		gamePoints := 0
		copiesOfThisGame := copies[i]
		for _, num := range playingNums {
			if !slices.Contains(winningNums, num) {
				continue
			}
			gamePoints++
			copies[i+gamePoints] += copiesOfThisGame + 1
		}
		sum += copiesOfThisGame + 1
	}

	return strconv.Itoa(sum)
}

func parseNumbersToSlice(line string, start int) ([]int, int) {
	nums := []int{}
	currentNumStr := ""
	lastIndex := 0
	for i, ch := range line[start:] {
		if ch == '|' {
			lastIndex = start + i
			break
		}

		if ch == ' ' {
			if currentNumStr == "" {
				continue
			}
			num, err := strconv.Atoi(currentNumStr)
			if err == nil {
				nums = append(nums, num)
			}
			currentNumStr = ""
			continue
		}

		currentNumStr += string(ch)
	}
	num, err := strconv.Atoi(currentNumStr)
	if err == nil {
		nums = append(nums, num)
	}
	return nums, lastIndex
}

func parseNumbersToNewMap(line string, start int) (map[int]int, int) {
	nums := make(map[int]int)
	currentNumStr := ""
	lastIndex := 0
	for i, ch := range line[start:] {
		if ch == '|' {
			lastIndex = start + i
			break
		}

		if ch == ' ' {
			if currentNumStr == "" {
				continue
			}
			num, err := strconv.Atoi(currentNumStr)
			if err == nil {
				nums[num] = num
			}
			currentNumStr = ""
			continue
		}

		currentNumStr += string(ch)
	}
	num, err := strconv.Atoi(currentNumStr)
	if err == nil {
		nums[num] = num
	}
	return nums, lastIndex
}

func parseNumbersToExistingMap(line string, start int, nums map[int]int) int {
	clear(nums)
	currentNumStr := ""
	lastIndex := 0
	for i, ch := range line[start:] {
		if ch == '|' {
			lastIndex = start + i
			break
		}

		if ch == ' ' {
			if currentNumStr == "" {
				continue
			}
			num, err := strconv.Atoi(currentNumStr)
			if err == nil {
				nums[num] = num
			}
			currentNumStr = ""
			continue
		}

		currentNumStr += string(ch)
	}
	num, err := strconv.Atoi(currentNumStr)
	if err == nil {
		nums[num] = num
	}
	return lastIndex
}
