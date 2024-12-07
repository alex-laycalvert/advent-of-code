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
// Part 2: Not Implemented
type Day7 struct {
	Input []string
}

// avg ~69.6ms
func (d Day7) Part1() string {
	answer := 0

	for _, line := range d.Input {
		parts := strings.Split(line, ":")
		expectedResult, err := strconv.Atoi(parts[0])
		utils.PanicErr(err)
		nums := parseNumbers(parts[1])

		// iterate through 2^(n-1) times
		for i := range int(math.Pow(2, float64(len(nums))-1)) {
			result := nums[0]
			ops := generateOps(i, len(nums)-1)
			onlyAddition := true
			for j, op := range ops {
				switch op {
				case '+':
					result += nums[j+1]
					break
				case '*':
					result *= nums[j+1]
					onlyAddition = false
					break
				default:
					break
				}

				if result > expectedResult {
					result = nums[0]
					if onlyAddition {
						break
					}
					continue
				}
			}
			if result == expectedResult {
				answer += expectedResult
                break
			}
		}
	}

	return strconv.Itoa(answer)
}

func (d Day7) Part2() string {
	return "Not Implemented"
}

func generateOps(perm int, length int) []rune {
	var result []rune
	for i := perm; i > 0; i >>= 1 {
		if i&1 == 1 {
			result = append([]rune{'*'}, result...)
		} else {
			result = append([]rune{'+'}, result...)
		}
	}

	for len(result) < length {
		result = append([]rune{'+'}, result...)
	}

	return result
}

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
