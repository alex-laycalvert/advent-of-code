package year2024

import (
	"strconv"
)

// ANSWERS:
//
// Part 1: 174561379
//
// Part 2: 106921067
type Day3 struct {
	Input []string
}

func (d Day3) Part1() string {
	sum := 0

	for _, line := range d.Input {
		sum += parseLine(line)
	}

	return strconv.Itoa(sum)
}

func (d Day3) Part2() string {
	sum := 0
	isEnabled := true

	for _, line := range d.Input {
		product, newIsEnabled := parseLineWithConditionals(line, isEnabled)
		isEnabled = newIsEnabled
		sum += product
	}

	return strconv.Itoa(sum)
}

func parseLine(line string) int {
	i := 0
	sum := 0
	for {
		if i >= len(line) {
			break
		}
		char := line[i]
		switch char {
		// start of possible `mul` instruction
		case 'm':
			product, newI, ok := parseMul(line, i)
			if ok {
				sum += product
			}
			i = newI
			break
		default:
			i++
			break
		}
	}
	return sum
}

func parseLineWithConditionals(line string, enabled bool) (int, bool) {
	i := 0
	sum := 0
	for {
		if i >= len(line) {
			break
		}
		char := line[i]
		switch char {
		// start of possible conditional `do` or `don't`
		case 'd':
			shouldEnable, newI, ok := parseConditional(line, i)
			if ok {
				enabled = shouldEnable
			}
			i = newI
			break
		// start of possible `mul` instruction
		case 'm':
			if !enabled {
				i++
				break
			}
			product, newI, ok := parseMul(line, i)
			if ok {
				sum += product
			}
			i = newI
			break
		default:
			i++
			break
		}
	}
	return sum, enabled
}

// parseMul attempts to parse a `mul` instruction in the `line`
// starting at the `start` index. Assumes that `line[start] == 'm'`.
//
// If a `mul` instruction is found, the function returns the product
// of the two numbers, the index of the character after the closing
// parenthesis, and true.
//
// Otherwise, the function returns 0, the index of the failed token (new
// start index), and false.
func parseMul(line string, start int) (int, int, bool) {
	ok, start := checkWord(line, start, "mul(")
	if !ok {
		return 0, start, false
	}

	// parse first number
	num1, start, ok := parseNumber(line, start)
	if !ok {
		return 0, start, false
	}

	// parse comma
	if !checkChar(line, start, ',') {
		return 0, start, false
	}
	start++

	// parse second number
	num2, start, ok := parseNumber(line, start)
	if !ok {
		return 0, start, false
	}

	// parse closing parenthesis
	if !checkChar(line, start, ')') {
		return 0, start, false
	}

	return num1 * num2, start + 1, true
}

func checkChar(line string, i int, char byte) bool {
	return i < len(line) && line[i] == char
}

func checkWord(line string, start int, word string) (bool, int) {
	for i, char := range word {
		if !checkChar(line, start+i, byte(char)) {
			return false, start + i
		}
	}
	return true, start + len(word)
}

func parseNumber(line string, start int) (int, int, bool) {
	numStr := ""
	for {
		if start >= len(line) {
			break
		}
		char := line[start]
		if char < '0' || char > '9' {
			break
		}
		numStr += string(char)
		start++
	}
	num, err := strconv.Atoi(numStr)
	return num, start, err == nil
}

func parseConditional(line string, start int) (bool, int, bool) {
	ok, doStart := checkWord(line, start, "do()")
	if ok {
		return true, doStart, true
	}

	ok, dontStart := checkWord(line, start, "don't()")
	if ok {
		return false, dontStart, true
	}

	return false, start, false
}
