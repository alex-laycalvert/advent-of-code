package year2024

import (
	"strconv"
)

// ANSWERS:
//
// Part 1: 36954
//
// Part 2: 79352015273424
type Day13 struct {
	Input []string
}

func (d Day13) Part1() string {
	answer := 0
	for i := 0; i < len(d.Input); i += 4 {
		clawMachine := NewClawMachine(d.Input[i:i+3], 0)
		tokens, canWin := clawMachine.TokensToWin()
		if canWin {
			answer += tokens
		}
	}

	return strconv.Itoa(answer)
}

func (d Day13) Part2() string {
	prizeOffset := 10_000_000_000_000

	answer := 0
	for i := 0; i < len(d.Input); i += 4 {
		clawMachine := NewClawMachine(d.Input[i:i+3], prizeOffset)
		tokens, canWin := clawMachine.TokensToWin()
		if canWin {
			answer += tokens
		}
	}

	return strconv.Itoa(answer)
}

type ClawMachine struct {
	AX     int
	AY     int
	BX     int
	BY     int
	PrizeX int
	PrizeY int
}

type ClawMachineResult struct {
	tokens int
	canWin bool
}

func NewClawMachine(input []string, prizeOffset int) ClawMachine {
	a := parseNumbers(input[0])
	b := parseNumbers(input[1])
	prize := parseNumbers(input[2])

	return ClawMachine{
		AX:     a[0],
		AY:     a[1],
		BX:     b[0],
		BY:     b[1],
		PrizeX: prize[0] + prizeOffset,
		PrizeY: prize[1] + prizeOffset,
	}
}

func (c *ClawMachine) TokensToWin() (tokens int, canWin bool) {
	a := ((c.PrizeX * c.BY) - (c.PrizeY * c.BX)) / ((c.AX * c.BY) - (c.AY * c.BX))
	b := (c.PrizeX - a*c.AX) / c.BX

	x := a*c.AX + b*c.BX
	y := a*c.AY + b*c.BY

	if x != c.PrizeX || y != c.PrizeY {
		return 0, false
	}

	return a*3 + b, true
}
