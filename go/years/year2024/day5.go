package year2024

import (
	"slices"
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 5374
//
// Part 2: 4260
type Day5 struct {
	Input []string
}

func (d Day5) Part1() string {
	return d.part1_mapOfBeforeNums()
}

// avg time of ~1.19ms per iter over 10k iters
func (d *Day5) part1_mapOfBeforeNums() string {
	rules, startIndex := parseRules(d.Input)

	sum := 0
	for _, line := range d.Input[startIndex:] {
		nums := strings.Split(line, ",")
		if middleNum, reportIsGood := isLineOrdered(nums, rules); reportIsGood {
			num, err := strconv.Atoi(middleNum)
			utils.PanicErr(err)
			sum += num
		}

	}

	return strconv.Itoa(sum)
}

func (d Day5) Part2() string {
	// return d.part2_permutations()
	return d.part2_reordering()
}

// avg time of ~1.57ms per iter over 10k iters
//
// My original solution would iterate through every permmutation of
// incorrect lines until it found a correct one.
//
// I never ran the permutation version to completion because in the
// time it took for me to think about and implement this solution, it
// still never finished (probably about 10mins or so).
func (d *Day5) part2_reordering() string {
	rules, startIndex := parseRules(d.Input)

	sum := 0
	for _, line := range d.Input[startIndex:] {
		nums := strings.Split(line, ",")
		_, initialIsGood := isLineOrdered(nums, rules)
		if initialIsGood {
			continue
		}

		correctlyOrderedNums := make([]string, len(nums))
		for i, num := range nums {
			correctlyOrderedNums[i] = num

			for j := i; j > 0; j-- {
				beforeNums := rules[correctlyOrderedNums[j-1]]
				if !slices.Contains(beforeNums, correctlyOrderedNums[j]) {
					break
				}
				// swap
				correctlyOrderedNums[j-1], correctlyOrderedNums[j] = correctlyOrderedNums[j], correctlyOrderedNums[j-1]
			}
		}

		num, err := strconv.Atoi(correctlyOrderedNums[len(correctlyOrderedNums)/2])
		utils.PanicErr(err)
		sum += num
	}

	return strconv.Itoa(sum)
}

func parseRules(input []string) (map[string][]string, int) {
	rules := map[string][]string{}
	for i, line := range input {
		if strings.TrimSpace(line) == "" {
			return rules, i + 1
		}

		rule := strings.Split(line, "|")
		before := rule[0]
		after := rule[1]
		rules[after] = append(rules[after], before)
	}

	return rules, 0
}

func isLineOrdered(line []string, rules map[string][]string) (string, bool) {
	errorNums := []string{}
	middleNum := line[len(line)/2]
	for _, n := range line {
		if slices.Contains(errorNums, n) {
			return "", false
		}
		beforeline, ok := rules[n]
		if !ok {
			continue
		}
		errorNums = append(errorNums, beforeline...)
	}
	return middleNum, true
}
