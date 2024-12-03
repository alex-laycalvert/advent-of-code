package year2023_test

import (
	"testing"

	"github.com/alex-laycalvert/advent-of-code/years/year2023"
)

func TestDay1(t *testing.T) {
	d := year2023.Day1{Input: []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}}
	expected := "142"
	if actual := d.Part1(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}

	d = year2023.Day1{Input: []string{
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
	if actual := d.Part2(); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
