package year2023

import (
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 2076
//
// Part 2: 70950
type Day2 struct {
	Input []string
}

func (d Day2) Part1() string {
	totalColors := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for _, line := range d.Input {
		selections := strings.Split(line, ":")
		gameID, err := strconv.Atoi(selections[0][5:])
		utils.PanicErr(err)

		selections = strings.Split(selections[1], ";")
		isPossible := true
		for _, part := range selections {
			marbles := strings.Split(strings.TrimSpace(part), ",")
			for _, marble := range marbles {
				parts := strings.Split(strings.TrimSpace(marble), " ")
				count, err := strconv.Atoi(parts[0])
				utils.PanicErr(err)
				color := parts[1]
				if total, _ := totalColors[color]; count > total {
					isPossible = false
					break
				}
			}

			if !isPossible {
				break
			}
		}
		if isPossible {
			sum += gameID
		}
	}
	return strconv.Itoa(sum)
}

func (d Day2) Part2() string {
	minColors := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	sum := 0
	for _, line := range d.Input {
		selections := strings.Split(line, ":")
		selections = strings.Split(selections[1], ";")
		for _, part := range selections {
			marbles := strings.Split(strings.TrimSpace(part), ",")
			for _, marble := range marbles {
				parts := strings.Split(strings.TrimSpace(marble), " ")
				count, err := strconv.Atoi(parts[0])
				utils.PanicErr(err)
				color := parts[1]
				if currentMin, _ := minColors[color]; count > currentMin {
					minColors[color] = count
				}
			}
		}

		numRed, _ := minColors["red"]
		numGreen, _ := minColors["green"]
		numBlue, _ := minColors["blue"]
		sum += numRed * numGreen * numBlue

		minColors["red"] = 0
		minColors["green"] = 0
		minColors["blue"] = 0
	}
	return strconv.Itoa(sum)
}
