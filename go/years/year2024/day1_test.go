package year2024_test

import (
	"testing"

	"github.com/alex-laycalvert/advent-of-code/years/year2024"
)

func TestDay1(t *testing.T) {
	day1 := year2024.Day1{Input: []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}}

	expectedPart1 := "11"
	part1 := day1.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "31"
	part2 := day1.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}
