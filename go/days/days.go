package days

import "github.com/alex-laycalvert/advent-of-code/years/year2024"

type Day interface {
	Part1() string
	Part2() string
}

func NewDay(year int, day int, input []string) Day {
	switch year {
	case 2024:
		return new2024Day(day, input)
	default:
		return nil
	}
}

func new2024Day(day int, input []string) Day {
	switch day {
	case 1:
		return year2024.Day1{Input: input}
	case 2:
		return year2024.Day2{Input: input}
	case 3:
		return year2024.Day3{Input: input}
	case 4:
		return year2024.Day4{Input: input}
	case 5:
		return year2024.Day5{Input: input}
	case 6:
		return year2024.Day6{Input: input}
	case 7:
		return year2024.Day7{Input: input}
	case 8:
		return year2024.Day8{Input: input}
	case 9:
		return year2024.Day9{Input: input}
	case 10:
		return year2024.Day10{Input: input}
	case 11:
		return year2024.Day11{Input: input}
	case 12:
		return year2024.Day12{Input: input}
	case 13:
		return year2024.Day13{Input: input}
	case 14:
		return year2024.Day14{Input: input}
	case 15:
		return year2024.Day15{Input: input}
	case 16:
		return year2024.Day16{Input: input}
	case 17:
		return year2024.Day17{Input: input}
	case 18:
		return year2024.Day18{Input: input}
	case 19:
		return year2024.Day19{Input: input}
	case 20:
		return year2024.Day20{Input: input}
	case 21:
		return year2024.Day21{Input: input}
	case 22:
		return year2024.Day22{Input: input}
	case 23:
		return year2024.Day23{Input: input}
	case 24:
		return year2024.Day24{Input: input}
	case 25:
		return year2024.Day25{Input: input}
	}
	return nil
}
