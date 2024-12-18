package year2024

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 7,1,5,2,4,0,7,6,1
//
// Part 2: 37222273957364
type Day17 struct {
	Input []string
}

func (d Day17) Part1() string {
	computer := NewComputer(d.Input)
	output := computer.Execute()
	return output
}

func (d Day17) Part2() string {
	computer := NewComputer(d.Input)
	possibleAs := computer.Backtrack()

	answer := 0
	for i, a := range possibleAs {
		if i == 0 || a < answer {
			answer = a
		}
	}

	return strconv.Itoa(answer)
}

const (
	ADV = 0
	BXL = 1
	BST = 2
	JNZ = 3
	BXC = 4
	OUT = 5
	BDV = 6
	CDV = 7
)

type Computer struct {
	RA int
	RB int
	RC int
	I  int

	Program []int
}

func NewComputer(input []string) Computer {
	computer := Computer{I: 0}
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
	if c.I >= len(c.Program) {
		return 0, false, true
	}
	switch c.Program[c.I] {
	case ADV:
		numerator := c.RA
		operand := c.comboOperand(c.Program[c.I+1])
		c.RA = numerator >> operand
		break
	case BXL:
		operand := c.Program[c.I+1]
		c.RB = operand ^ c.RB
		break
	case BST:
		operand := c.comboOperand(c.Program[c.I+1])
		c.RB = operand % 8
		break
	case JNZ:
		if c.RA != 0 {
			operand := c.Program[c.I+1]
			c.I = operand
			return 0, false, false
		}
		break
	case BXC:
		c.RB = c.RB ^ c.RC
		break
	case OUT:
		operand := c.comboOperand(c.Program[c.I+1])
		output = operand % 8
		hasOutput = true
		break
	case BDV:
		numerator := c.RA
		operand := c.comboOperand(c.Program[c.I+1])
		denominator := int(math.Pow(2, float64(operand)))
		c.RB = numerator / denominator
		break
	case CDV:
		numerator := c.RA
		operand := c.comboOperand(c.Program[c.I+1])
		denominator := int(math.Pow(2, float64(operand)))
		c.RC = numerator / denominator
		break
	default:
		break
	}
	c.I += 2
	return output, hasOutput, false
}

func (c Computer) Backtrack() []int {
	if c.Program[len(c.Program)-2] != JNZ {
		panic("Unable to backtrack")
	}

	revProgram := make([]int, len(c.Program))
	copy(revProgram, c.Program)
	slices.Reverse(revProgram)

	valid := []int{0}
	for _, value := range revProgram {
		currentValid := []int{}
		for _, validNextA := range valid {
			for n := range 8 {
				a := (validNextA << 3) | n
				computer := &Computer{
					RA:      a,
					RB:      0,
					RC:      0,
					I:       0,
					Program: c.Program,
				}

				for output, ok, done := computer.Step(); !done; output, ok, done = computer.Step() {
					if !ok {
						continue
					}
					if output == value {
						currentValid = append(currentValid, a)
					}
					break
				}
			}
		}
		valid = currentValid
	}

	return valid
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

func LambertW(z float64) float64 {
	if z < -1/math.E {
		panic("less than -1/e")
	}

	// Initial guess for W(z). Use log(z + 1) for large z, or 0 for small z.
	var w float64
	if z > 1 {
		w = math.Log(z)
	} else {
		w = z
	}

	// Iterate using Newton's method to refine the solution.
	for i := 0; i < 100; i++ { // Limit iterations to 100
		eW := math.Exp(w)
		wNext := w - (w*eW-z)/(eW*(w+1)-(w+2)*(w*eW-z)/(2*(w+1)))
		if math.Abs(wNext-w) < 1e-10 {
			return wNext
		}
		w = wNext
	}

	return math.NaN()
}
