package year2024

import (
	"math"
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWER:
//
// Part 1: 945512582195
//
// Part 2: 271691107779347
type Day7 struct {
	Input []string
}

// avg ~42ms over 1k iterations
func (d Day7) Part1() string {
	answer := 0

	for _, line := range d.Input {
		parts := strings.Split(line, ":")
		expectedResult, err := strconv.Atoi(parts[0])
		utils.PanicErr(err)
		nums := parseNumbers(parts[1])

		// iterate through 2^(n-1) times
		for i := range int(math.Pow(2, float64(len(nums))-1)) {
			result := evaluateAtIteration(nums, i, 2)
			if result == expectedResult {
				answer += expectedResult
				break
			}
		}
	}

	return strconv.Itoa(answer)
}

// avg ~2.9s over 10 iterations. too long to test more.
func (d Day7) Part2() string {
	answer := 0

	for _, line := range d.Input {
		parts := strings.Split(line, ":")
		expectedResult, err := strconv.Atoi(parts[0])
		utils.PanicErr(err)
		nums := parseNumbers(parts[1])

		for i := range int(math.Pow(3, float64(len(nums))-1)) {
			result := evaluateAtIteration(nums, i, 3)
			if result == expectedResult {
				answer += expectedResult
				break
			}
		}
	}

	return strconv.Itoa(answer)
}

// given a number, generates a string of runes that represent that number
// at the given base.
//
// Example:
//
//	generateNumberAtBase(5, 4, 2)  // = "0101"
//	generateNumberAtBase(7, 10, 2) // = "0000000111"
//	generateNumberAtBase(15, 4, 3) // = "0112"
func generateNumberAtBase(num int, length int, base int) []rune {
	result := make([]rune, length)
	for i := range length {
		place := int(math.Pow(
			float64(base),
			float64(length-i-1),
		))

		for j := base - 1; j >= 0; j-- {
			if num < place*j {
				continue
			}
			result[i] = rune(48 + j)
			num -= place * j
			break
		}
	}
	return result
}

// given a list of nums, evalutes the expression based on the given iteration.
// The given iteration determines what order of operations to use.
func evaluateAtIteration(nums []int, iter int, base int) int {
	result := nums[0]
	ops := generateNumberAtBase(iter, len(nums)-1, base)
	for j, op := range ops {
		switch op {
		case '0':
			result += nums[j+1]
			break
		case '1':
			result *= nums[j+1]
			break
		case '2':
			r, err := strconv.Atoi(
				strconv.Itoa(result) + strconv.Itoa(nums[j+1]),
			)
			utils.PanicErr(err)
			result = r
		default:
			break
		}
	}
	return result
}

// given a string, parses out all numbers and returns them as a slice of ints.
func parseNumbers(s string) []int {
	numbers := []int{}

	currentNum := ""
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			currentNum += string(ch)
			continue
		}

		if currentNum != "" {
			num, _ := strconv.Atoi(currentNum)
			numbers = append(numbers, num)
			currentNum = ""
		}
	}

	if currentNum != "" {
		num, err := strconv.Atoi(currentNum)
		utils.PanicErr(err)
		numbers = append(numbers, num)
	}
	return numbers
}
