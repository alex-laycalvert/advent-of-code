package days

type Day interface {
	Part1() string
	Part2() string
}

func NewDay(day int, input string) Day {
	switch day {
	case 1:
		return Day1{input}
	case 2:
		return Day2{input}
	case 3:
		return Day3{input}
	case 4:
		return Day4{input}
	case 5:
		return Day5{input}
	case 6:
		return Day6{input}
	case 7:
		return Day7{input}
	case 8:
		return Day8{input}
	case 9:
		return Day9{input}
	case 10:
		return Day10{input}
	case 11:
		return Day11{input}
	case 12:
		return Day12{input}
	case 13:
		return Day13{input}
	case 14:
		return Day14{input}
	case 15:
		return Day15{input}
	case 16:
		return Day16{input}
	case 17:
		return Day17{input}
	case 18:
		return Day18{input}
	case 19:
		return Day19{input}
	case 20:
		return Day20{input}
	case 21:
		return Day21{input}
	case 22:
		return Day22{input}
	case 23:
		return Day23{input}
	case 24:
		return Day24{input}
	case 25:
		return Day25{input}
	}
	return nil
}
