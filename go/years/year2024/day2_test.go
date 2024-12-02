package year2024_test

import (
	"testing"

	"github.com/alex-laycalvert/advent-of-code/years/year2024"
)

func TestDay2(t *testing.T) {
	day2 := year2024.Day2{Input: []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}}

	expectedPart1 := "2"
	part1 := day2.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "4"
	part2 := day2.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}
