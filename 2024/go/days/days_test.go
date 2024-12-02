package days_test

import (
	"os"
	"testing"

	"github.com/alex-laycalvert/advent-of-code/2024/days"
	"github.com/alex-laycalvert/advent-of-code/2024/utils"
)

func TestDay1(t *testing.T) {
	inpputString := `3   4
4   3
2   5
1   3
3   9
3   3
`
	file, err := os.CreateTemp("/tmp", "input_1_test")
	utils.PanicErr(err)
	defer file.Close()
	defer os.Remove(file.Name())

	_, err = file.WriteString(inpputString)
	utils.PanicErr(err)

	day := days.NewDay(1, file.Name())

	expectedPart1 := "11"
	part1 := day.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "31"
	part2 := day.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay2(t *testing.T) {
	inpputString := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`
	file, err := os.CreateTemp("/tmp", "input_2_test")
	utils.PanicErr(err)
	defer file.Close()
	defer os.Remove(file.Name())

	_, err = file.WriteString(inpputString)
	utils.PanicErr(err)

	day := days.NewDay(2, file.Name())

	expectedPart1 := "2"
	part1 := day.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "4"
	part2 := day.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}
