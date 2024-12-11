package year2024

import (
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 217812
//
// Part 2: 259,112,729,857,522
type Day11 struct {
	Input []string
}

// avg ~0.5ms per iter over 10k iters
func (d Day11) Part1() string {
	input := parseDay11Input(d.Input[0])
	size := blink(input, 25)
	return strconv.Itoa(size)
}

// avg ~45-50ms per iter over 100 iters
func (d Day11) Part2() string {
	input := parseDay11Input(d.Input[0])
	size := blink(input, 75)
	return strconv.Itoa(size)
}

func parseDay11Input(line string) []int {
	input := strings.Split(line, " ")
	numInput := make([]int, len(input))
	for i, v := range input {
		val, err := strconv.Atoi(v)
		utils.PanicErr(err)
		numInput[i] = val
	}
	return numInput
}

func blink(input []int, times int) int {
	cache := make(map[string]int)
	size := 0
	for _, num := range input {
		size += blinkNumber(num, times, cache)
	}
	return size
}

func blinkNumber(number int, times int, cache map[string]int) int {
	key := strconv.Itoa(number) + ":" + strconv.Itoa(times)
	if val, ok := cache[key]; ok {
		return val
	}

	if times == 0 {
		cache[key] = 1
		return 1
	}

	if number == 0 {
		value := blinkNumber(1, times-1, cache)
		cache[key] = value
		return value
	}

	valAsStr := strconv.Itoa(number)
	if len(valAsStr)%2 != 0 {
		value := blinkNumber(number*2024, times-1, cache)
		cache[key] = value
		return value
	}

	left, err := strconv.Atoi(valAsStr[:len(valAsStr)/2])
	utils.PanicErr(err)
	right, err := strconv.Atoi(valAsStr[len(valAsStr)/2:])
	utils.PanicErr(err)

	value := blinkNumber(left, times-1, cache) + blinkNumber(right, times-1, cache)
	cache[key] = value
	return value
}
