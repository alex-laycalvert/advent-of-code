package year2024

import (
	"math"
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 7,1,5,2,4,0,7,6,1
//
// Part 2:
type Day17 struct {
	Input []string
}

func (d Day17) Part1() string {
	computer := NewComputer(d.Input)
	output := computer.Execute()
	return output
}

func (d Day17) Part2() string {
	initial := NewComputer(d.Input)
    programStr := ""
    for i, p := range initial.Program {
        if i > 0 {
            programStr += ","
        }
        programStr += strconv.Itoa(p)
    }
    answer := 0
	for i := initial.RA; ; i++ {
		computer := NewComputer(d.Input)
		computer.RA = i
		output := computer.Execute()
        if output == programStr {
            answer = i
            break
        }
	}

	return strconv.Itoa(answer)
}

type Computer struct {
	RA     int
	RB     int
	RC     int
	i      int
	output []int

	Program []int
}

func NewComputer(input []string) Computer {
	computer := Computer{i: 0}
	parsingRegisters := true
	for _, line := range input {
		if line == "" {
			parsingRegisters = false
			continue
		}

		if !parsingRegisters {
			computer.Program = utils.ParseNumbers(line)
			continue
		}

		parts := strings.Split(line, ":")
		value, err := strconv.Atoi(parts[1][1:])
		utils.PanicErr(err)
		switch parts[0][len(parts[0])-1:] {
		case "A":
			computer.RA = value
			break
		case "B":
			computer.RB = value
			break
		case "C":
			computer.RC = value
			break
		default:
			break
		}
	}
	return computer
}

func (c *Computer) Execute() string {
	output := []int{}
	for out, hasOutput, done := c.Step(); !done; out, hasOutput, done = c.Step() {
		if hasOutput {
			output = append(output, out)
		}
	}

	outputStr := ""
	for i, o := range output {
		if i > 0 {
			outputStr += ","
		}
		outputStr += strconv.Itoa(o)
	}

	return outputStr
}

func (c *Computer) Step() (output int, hasOutput bool, done bool) {
	if c.i >= len(c.Program) {
		return 0, false, true
	}
	switch c.Program[c.i] {
	case 0:
		// adv
		numerator := c.RA
		operand := c.comboOperand(c.Program[c.i+1])
		denominator := int(math.Pow(2, float64(operand)))
		c.RA = numerator / denominator
		break
	case 1:
		// bxl
		operand := c.Program[c.i+1]
		c.RB = operand ^ c.RB
		break
	case 2:
		// bst
		operand := c.comboOperand(c.Program[c.i+1])
		c.RB = operand % 8
		break
	case 3:
		// jnz
		if c.RA != 0 {
			operand := c.Program[c.i+1]
			c.i = operand
			return 0, false, false
		}
		break
	case 4:
		// bxc
		c.RB = c.RB ^ c.RC
		break
	case 5:
		// out
		operand := c.comboOperand(c.Program[c.i+1])
		output = operand % 8
		hasOutput = true
		break
	case 6:
		// bdv
		numerator := c.RA
		operand := c.comboOperand(c.Program[c.i+1])
		denominator := int(math.Pow(2, float64(operand)))
		c.RB = numerator / denominator
		break
	case 7:
		// cdv
		numerator := c.RA
		operand := c.comboOperand(c.Program[c.i+1])
		denominator := int(math.Pow(2, float64(operand)))
		c.RC = numerator / denominator
		break
	default:
		break
	}
	c.i += 2
	return output, hasOutput, false
}

func (c Computer) comboOperand(operand int) int {
	if operand <= 3 {
		return operand
	}

	switch operand {
	case 4:
		return c.RA
	case 5:
		return c.RB
	case 6:
		return c.RC
	default:
		return 0
	}
}
