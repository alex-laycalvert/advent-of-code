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

// func TestDay2(t *testing.T) {
// 	day2 := year2023.Day2{Input: []string{
// 		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
// 		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
// 		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
// 		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
// 		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
// 	}}
// 	expected := "8"
// 	if actual := day2.Part1(); actual != expected {
// 		t.Errorf("Expected %s but got %s", expected, actual)
// 	}
//
// 	expected = "2286"
// 	if actual := day2.Part2(); actual != expected {
// 		t.Errorf("Expected %s but got %s", expected, actual)
// 	}
// }
