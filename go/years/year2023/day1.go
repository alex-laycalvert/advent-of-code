package year2023

import (
	"slices"
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 54634
//
// Part 2: 53855
type Day1 struct {
	Input []string
}

func (d Day1) Part1() string {
	sum := 0
	for _, line := range d.Input {
		left := 0
		right := len(line) - 1
		leftDig := ""
		rightDig := ""
		for {
			if leftDig == "" {
				leftCh := line[left]
				if leftCh >= '0' && leftCh <= '9' {
					leftDig = string(leftCh)
				} else {
					left++
				}
				continue
			}

			if rightDig == "" {
				rightCh := line[right]
				if rightCh >= '0' && rightCh <= '9' {
					rightDig = string(rightCh)
					break
				} else {
					right--
				}
			}
		}

		num, err := strconv.Atoi(leftDig + rightDig)
		utils.PanicErr(err)
		sum += num
	}
	return strconv.Itoa(sum)
}

func (d Day1) Part2() string {
	validStartChars := []byte{'o', 't', 'f', 's', 'e', 'n', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	sum := 0
	for _, line := range d.Input {
		left := 0
		right := len(line) - 1
		leftDig := ""
		rightDig := ""
		for {
			if leftDig == "" {
				leftCh := line[left]
				if !slices.Contains(validStartChars, leftCh) {
					left++
					continue
				}
				if num, matches := substrMatchesNum(line, left); matches {
					leftDig = num
				}
				left++
				continue
			}

			if rightDig == "" {
				rightCh := line[right]
				if !slices.Contains(validStartChars, rightCh) {
					right--
					continue
				}
				if num, matches := substrMatchesNum(line, right); matches {
					rightDig = num
					break
				}
				right--
			}
		}

		num, err := strconv.Atoi(leftDig + rightDig)
		utils.PanicErr(err)
		sum += num
	}

	return strconv.Itoa(sum)
}

var wordToNum = map[string]string{
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func substrMatchesNum(line string, start int) (string, bool) {
	for i := 1; i <= 5; i++ {
		if i == 2 {
			continue
		}
		if len(line) < start+i {
			return "", false
		}

		if num, ok := wordToNum[line[start:start+i]]; ok {
			return num, true
		}

	}

	return "", false
}
