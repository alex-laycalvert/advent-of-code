package year2023_test

import (
	"testing"

	"github.com/alex-laycalvert/advent-of-code/years/year2023"
)

func TestDay1(t *testing.T) {
	day1 := year2023.Day1{Input: []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}}
	expected := "142"
	if actual := day1.Part1(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}

	day1 = year2023.Day1{Input: []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
		"one",
		"oneone",
		"oneone1",
		"ooooooooneone1",
	}}

	expected = "325"
	if actual := day1.Part2(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestDay2(t *testing.T) {
	day2 := year2023.Day2{Input: []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}}

	expected := "8"
	if actual := day2.Part1(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}

	expected = "2286"
	if actual := day2.Part2(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestDay3(t *testing.T) {
	day3 := year2023.Day3{Input: []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}}
	expected := "4361"
	if actual := day3.Part1(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}

	// expected = "467835"
	expected = "Not Implemented"
	if actual := day3.Part2(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestDay4(t *testing.T) {
	day4 := year2023.Day4{Input: []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}}

	expected := "13"
	if actual := day4.Part1(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}

	expected = "30"
	if actual := day4.Part2(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
