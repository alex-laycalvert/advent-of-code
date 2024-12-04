package year2023

import (
	"strconv"
)

// ANSWERS:
//
// Part 1: 536202
//
// Part 2:
type Day3 struct {
	Input []string
}

func (d Day3) Part1() string {
	sum := 0
	for row, line := range d.Input {
		for col, char := range line {
			// haven't reached a valid symbol
			if char == '.' || isNumeric(byte(char)) {
				continue
			}

			num, _ := parseNumber(line, col, -1)
			sum += num
			num, _ = parseNumber(line, col, 1)
			sum += num

			if row > 0 {
				num, aboveIsNumeric := parseNumber(d.Input[row-1], col, -1)
				if !aboveIsNumeric {
					num, _ := parseNumber(d.Input[row-1], col, 1)
					sum += num
				}
				sum += num
			}

			if row < len(line)-1 {
				num, aboveIsNumeric := parseNumber(d.Input[row+1], col, -1)
				if !aboveIsNumeric {
					num, _ := parseNumber(d.Input[row+1], col, 1)
					sum += num
				}
				sum += num
			}
		}
	}
	return strconv.Itoa(sum)
}

func (d Day3) Part2() string {
	return "Not Implemented"
}

func isNumeric(char byte) bool {
	return char >= '0' && char <= '9'
}

func parseNumber(line string, col int, dir int) (int, bool) {
	numStr := ""
	rootIsNumeric := isNumeric(line[col])
	if rootIsNumeric {
		numStr = string(line[col])
	}

	if rootIsNumeric || dir <= 0 {
		// get left digits
		for i := col - 1; i >= 0 && isNumeric(line[i]); i-- {
			numStr = string(line[i]) + numStr
		}
	}

	if rootIsNumeric || dir >= 0 {
		// get right digits
		for i := col + 1; i < len(line) && isNumeric(line[i]); i++ {
			numStr = numStr + string(line[i])
		}
	}

	num, _ := strconv.Atoi(numStr)
	return num, rootIsNumeric
}
