package days

import (
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/2024/utils"
)

// ANSWERS:
//
// Part 1: 516
//
// Part 2: 561
type Day2 struct {
	input string
}

func (d Day2) Part1() string {
	sum := 0
	for line := range utils.ReadLines(d.input) {
		parts := strings.Split(line, " ")
		isSafe := day2CheckSafety(parts)

		if isSafe {
			sum++
		}
	}

	return strconv.Itoa(sum)
}

func (d Day2) Part2() string {
	return d.part2_bounding()
}

func (d *Day2) part2_naive() string {
	sum := 0
	for line := range utils.ReadLines(d.input) {
		parts := strings.Split(line, " ")
		isSafe := day2CheckSafety(parts)

		if !isSafe {
			for i := 0; i < len(parts); i++ {
				if day2CheckSafety(utils.RemoveIndex(parts, i)) {
					isSafe = true
					break
				}
			}
		}

		if isSafe {
			sum++
		}
	}

	return strconv.Itoa(sum)
}

// About a 34% speedup over the naive approach
func (d *Day2) part2_bounding() string {
	sum := 0
	for line := range utils.ReadLines(d.input) {
		parts := strings.Split(line, " ")
		isSafe := false

		for i := -1; i < len(parts); i++ {
			if day2CheckSafetyGeneric(parts, i) {
				isSafe = true
				break
			}
		}

		if isSafe {
			sum++
		}
	}

	return strconv.Itoa(sum)
}

func day2CheckSafety(row []string) bool {
	isIncreasing := false
	for i := 0; i < len(row)-1; i++ {
		first, err := strconv.Atoi(row[i])
		utils.PanicErr(err)
		second, err := strconv.Atoi(row[i+1])
		utils.PanicErr(err)
		diff := first - second
		absDiff := utils.IntAbs(diff)
		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if i == 0 {
			isIncreasing = diff < 0
			continue
		}

		if (isIncreasing && diff > 0) || (!isIncreasing && diff < 0) {
			return false
		}
	}

	return true
}

func day2CheckSafetyGeneric(row []string, middle int) bool {
	isIncreasing := false
	end := len(row) - 1
	for i := 0; i < end; i++ {
		if i == middle {
			continue
		}
		if i == middle-1 && middle == end {
			continue
		}
		first, err := strconv.Atoi(row[i])
		utils.PanicErr(err)

		second, err := strconv.Atoi(row[i+1])
		utils.PanicErr(err)
		if i == middle-1 {
			second, err = strconv.Atoi(row[i+2])
			utils.PanicErr(err)
		}

		diff := first - second
		absDiff := utils.IntAbs(diff)
		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if i == 0 || (middle == 0 && i == 1) {
			isIncreasing = diff < 0
			continue
		}

		if (isIncreasing && diff > 0) || (!isIncreasing && diff < 0) {
			return false
		}
	}

	return true
}
