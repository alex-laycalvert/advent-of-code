package days

import (
	"errors"

	"github.com/alex-laycalvert/advent-of-code/years/year2023"
	"github.com/alex-laycalvert/advent-of-code/years/year2024"
)

type Day interface {
	Part1() string
	Part2() string
}

func NewDay(year int, day int, input []string) (Day, error) {
	switch year {
	case 2023:
		return new2023Day(day, input)
	case 2024:
		return new2024Day(day, input)
	default:
		return nil, errors.New("Not implemented")
	}
}

func new2023Day(day int, input []string) (Day, error) {
	switch day {
	case 1:
		return year2023.Day1{Input: input}, nil
	case 2:
		return year2023.Day2{Input: input}, nil
	case 3:
		return year2023.Day3{Input: input}, nil
	case 4:
		return year2023.Day4{Input: input}, nil
	case 5:
		return year2023.Day5{Input: input}, nil
	case 6:
		return year2023.Day6{Input: input}, nil
	case 7:
		return year2023.Day7{Input: input}, nil
	case 8:
		return year2023.Day8{Input: input}, nil
	case 9:
		return year2023.Day9{Input: input}, nil
	case 10:
		return year2023.Day10{Input: input}, nil
	case 11:
		return year2023.Day11{Input: input}, nil
	case 12:
		return year2023.Day12{Input: input}, nil
	case 13:
		return year2023.Day13{Input: input}, nil
	case 14:
		return year2023.Day14{Input: input}, nil
	case 15:
		return year2023.Day15{Input: input}, nil
	case 16:
		return year2023.Day16{Input: input}, nil
	case 17:
		return year2023.Day17{Input: input}, nil
	case 18:
		return year2023.Day18{Input: input}, nil
	case 19:
		return year2023.Day19{Input: input}, nil
	case 20:
		return year2023.Day20{Input: input}, nil
	case 21:
		return year2023.Day21{Input: input}, nil
	case 22:
		return year2023.Day22{Input: input}, nil
	case 23:
		return year2023.Day23{Input: input}, nil
	case 24:
		return year2023.Day24{Input: input}, nil
	case 25:
		return year2023.Day25{Input: input}, nil
	}
	return nil, errors.New("Not implemented")
}

func new2024Day(day int, input []string) (Day, error) {
	switch day {
	case 1:
		return year2024.Day1{Input: input}, nil
	case 2:
		return year2024.Day2{Input: input}, nil
	case 3:
		return year2024.Day3{Input: input}, nil
	case 4:
		return year2024.Day4{Input: input}, nil
	case 5:
		return year2024.Day5{Input: input}, nil
	case 6:
		return year2024.Day6{Input: input}, nil
	case 7:
		return year2024.Day7{Input: input}, nil
	case 8:
		return year2024.Day8{Input: input}, nil
	case 9:
		return year2024.Day9{Input: input}, nil
	case 10:
		return year2024.Day10{Input: input}, nil
	case 11:
		return year2024.Day11{Input: input}, nil
	case 12:
		return year2024.Day12{Input: input}, nil
	case 13:
		return year2024.Day13{Input: input}, nil
	case 14:
		return year2024.Day14{Input: input}, nil
	case 15:
		return year2024.Day15{Input: input}, nil
	case 16:
		return year2024.Day16{Input: input}, nil
	case 17:
		return year2024.Day17{Input: input}, nil
	case 18:
		return year2024.Day18{Input: input}, nil
	case 19:
		return year2024.Day19{Input: input}, nil
	case 20:
		return year2024.Day20{Input: input}, nil
	case 21:
		return year2024.Day21{Input: input}, nil
	case 22:
		return year2024.Day22{Input: input}, nil
	case 23:
		return year2024.Day23{Input: input}, nil
	case 24:
		return year2024.Day24{Input: input}, nil
	case 25:
		return year2024.Day25{Input: input}, nil
	}
	return nil, errors.New("Not implemented")
}
