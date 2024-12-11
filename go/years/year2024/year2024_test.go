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

func TestDay7(t *testing.T) {
	day7 := year2024.Day7{Input: []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}}

	expectedPart1 := "3749"
	part1 := day7.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "11387"
	part2 := day7.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay8(t *testing.T) {
	day8 := year2024.Day8{Input: []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}}

	expectedPart1 := "14"
	part1 := day8.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "34"
	part2 := day8.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay9(t *testing.T) {
	day9 := year2024.Day9{Input: []string{
		"2333133121414131402",
	}}

	expectedPart1 := "1928"
	part1 := day9.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "2858"
	part2 := day9.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay10(t *testing.T) {
	day10 := year2024.Day10{Input: []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}}

	expectedPart1 := "36"
	part1 := day10.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "81"
	part2 := day10.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

// 125 17

//    v
// 0: (125, 0) (17, 0)
// 1: (253000, 0), (17, 0)
// 2: (253, 0), (0, 2), (17, 0)
// 3: (512072, 0), (0, 2), (17, 0)
// 4: (512, 0) (72, 4), (0, 2), (17, 0)
// 5: (1036288, 0) (72, 4), (0, 2), (17, 0)
// 6: (2097446912, 0) (72, 4), (0, 2), (17, 0)

//                      v
// 0: (2097446912, 0) (72, 4), (0, 2), (17, 0)
// 1: (2097446912, 0) (7, 4) (2, 5), (0, 2), (17, 0)
// 2: (2097446912, 0) (14168, 4) (2, 5), (0, 2), (17, 0)

//                                 v
// 0: (2097446912, 0) (14168, 4) (2, 5), (0, 2), (17, 0)
// 1: (2097446912, 0) (14168, 4) (4048, 5), (0, 2), (17, 0)

//                                             v
// 0: (2097446912, 0) (14168, 4) (4048, 5), (0, 2), (17, 0)
// 1: (2097446912, 0) (14168, 4) (4048, 5), (1, 2), (17, 0)
// 2: (2097446912, 0) (14168, 4) (4048, 5), (2024, 2), (17, 0)
// 3: (2097446912, 0) (14168, 4) (4048, 5), (20, 2) (24, 5), (17, 0)
// 4: (2097446912, 0) (14168, 4) (4048, 5), (2, 2) (0, 6) (24, 5), (17, 0)

//                                                           v
// 0: (2097446912, 0) (14168, 4) (4048, 5), (2, 2) (0, 6) (24, 5), (17, 0)
// 1: (2097446912, 0) (14168, 4) (4048, 5), (2, 2) (0, 6) (2, 3) (4, 4), (17, 0)

//                                                                          v
// 0: (2097446912, 0) (14168, 4) (4048, 5), (2, 2) (0, 6) (40, 3) (48, 6) (4, 4), (17, 0)
// 1: (2097446912, 0) (14168, 4) (4048, 5), (2, 2) (0, 6) (40, 3) (48, 6) (8096, 4), (17, 0)
// 2: (2097446912, 0) (14168, 4) (4048, 5), (2, 2) (0, 6) (40, 3) (48, 6) (80, 4) (96, 6), (17, 0)

//                                                                                           v
// 0: (2097446912, 0) (14168, 4) (4048, 5), (2, 2) (0, 6) (40, 3) (48, 6) (80, 4) (96, 6), (17, 0)
// 1: (2097446912, 0) (14168, 4) (4048, 5), (2, 2) (0, 6) (40, 3) (48, 6) (80, 4) (96, 6), (1, 0) (7, 1)

func TestDay11(t *testing.T) {
	day11 := year2024.Day11{Input: []string{
		"125 17",
	}}

	expectedPart1 := "55312"
	part1 := day11.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "81"
	part2 := day11.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}
