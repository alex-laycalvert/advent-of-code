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

func TestDay3(t *testing.T) {
	day3 := year2024.Day3{Input: []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}}

	expectedPart1 := "161"
	part1 := day3.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "48"
	part2 := day3.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay4(t *testing.T) {
	day4 := year2024.Day4{Input: []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}}

	expectedPart1 := "18"
	part1 := day4.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "9"
	part2 := day4.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay5(t *testing.T) {
	day5 := year2024.Day5{Input: []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}}

	expectedPart1 := "143"
	part1 := day5.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "123"
	part2 := day5.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay6(t *testing.T) {
	day6 := year2024.Day6{Input: []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}}

	expectedPart1 := "41"
	part1 := day6.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "6"
	part2 := day6.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}
