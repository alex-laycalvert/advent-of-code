package year2024

import (
	"slices"
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 765748
//
// Part 2: 27732508
type Day1 struct {
	Input []string
}

func (d Day1) Part1() string {
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for _, line := range d.Input {
		parts := strings.Split(line, " ")
		leftNum, err := strconv.Atoi(parts[0])
		utils.PanicErr(err)
		leftList = append(leftList, leftNum)

		rightNum, err := strconv.Atoi(parts[len(parts)-1])
		utils.PanicErr(err)
		rightList = append(rightList, rightNum)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	sum := 0
	for i := 0; i < len(leftList); i++ {
		sum += utils.IntAbs(leftList[i] - rightList[i])
	}

	return strconv.Itoa(sum)
}

func (d Day1) Part2() string {
	leftList := make([]int, 0)
	rightFreq := make(map[int]int)

	for _, line := range d.Input {
		parts := strings.Split(line, " ")
		leftNum, err := strconv.Atoi(parts[0])
		utils.PanicErr(err)
		leftList = append(leftList, leftNum)

		rightNum, err := strconv.Atoi(parts[len(parts)-1])
		utils.PanicErr(err)
		if _, ok := rightFreq[rightNum]; ok {
			rightFreq[rightNum]++
		} else {
			rightFreq[rightNum] = 1
		}
	}

	sum := 0
	for _, num := range leftList {
		freq, ok := rightFreq[num]
		if !ok {
			continue
		}
		sum += num * freq
	}

	return strconv.Itoa(sum)
}
