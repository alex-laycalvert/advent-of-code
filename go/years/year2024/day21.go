package year2024

import (
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 94426
//
// Part 2: 118392478819140
type Day21 struct {
	Input []string
}

/*
* +---+---+---+
* | 7 | 8 | 9 |
* +---+---+---+
* | 4 | 5 | 6 |
* +---+---+---+
* | 1 | 2 | 3 |
* +---+---+---+
*     | 0 | A |
*     +---+---+
 */
var numpad = map[string]rune{
	"0,0": '7',
	"0,1": '8',
	"0,2": '9',
	"1,0": '4',
	"1,1": '5',
	"1,2": '6',
	"2,0": '1',
	"2,1": '2',
	"2,2": '3',
	"3,1": '0',
	"3,2": 'A',
}

/*
*     +---+---+
*     | ^ | A |
* +---+---+---+
* | < | v | > |
* +---+---+---+
 */
var dirpad = map[string]rune{
	"0,1": '^',
	"0,2": 'A',
	"1,0": '<',
	"1,1": 'v',
	"1,2": '>',
}

var instructions = []rune{'^', 'A', '<', 'v', '>'}

func (d Day21) Part1() string {
	sum := 0
	for _, line := range d.Input {
		num, err := strconv.Atoi(strings.TrimRight(line, "A"))
		utils.PanicErr(err)
		length := solve(line, 2)
		sum += num * length
	}

	return strconv.Itoa(sum)
}

func (d Day21) Part2() string {
	sum := 0
	for _, line := range d.Input {
		num, err := strconv.Atoi(strings.TrimRight(line, "A"))
		utils.PanicErr(err)
		length := solve(line, 25)
		sum += num * length
	}

	return strconv.Itoa(sum)
}

type D21Point struct {
	row, col int
}

func ToD21Point(key string) D21Point {
	coords := strings.Split(key, ",")
	row, _ := strconv.Atoi(coords[0])
	col, _ := strconv.Atoi(coords[1])
	return D21Point{row, col}
}

func (p D21Point) String() string {
	return strconv.Itoa(p.row) + "," + strconv.Itoa(p.col)
}

func (p *D21Point) Execute(instruction rune, pad map[string]rune) (D21Point, rune, bool) {
	switch instruction {
	case '^':
		return D21Point{p.row - 1, p.col}, '0', false
	case 'v':
		return D21Point{p.row + 1, p.col}, '0', false
	case '<':
		return D21Point{p.row, p.col - 1}, '0', false
	case '>':
		return D21Point{p.row, p.col + 1}, '0', false
	case 'A':
		return D21Point{p.row, p.col}, pad[p.String()], true
	}
	panic("Invalid instruction")
}

type NumpadNode struct {
	cost        int
	pos         D21Point
	instruction rune
	length      int
}

func (n NumpadNode) Less(other NumpadNode) bool {
	return n.cost < other.cost
}

type NumpadPathNode struct {
	current rune
	cost    int
	length  int
}

func (n NumpadPathNode) Less(other NumpadPathNode) bool {
	return n.cost < other.cost
}

func solve(code string, depth int) int {
	pq := utils.NewPriorityQueue[NumpadNode]()
	costmap := make(map[string]int)

	pq.Push(NumpadNode{
		cost:        0,
		pos:         D21Point{row: 3, col: 2},
		instruction: 'A',
		length:      0,
	})

	cache := make(map[string]int)
	for !pq.IsEmpty() {
		node, _ := pq.Pop()
		if node.length == len(code) {
			return node.cost
		}

		key := posInstructionLengthKey(node.pos, node.instruction, node.length)
		if _, ok := costmap[key]; ok {
			continue
		}
		costmap[key] = node.cost

		for _, newInstruction := range instructions {
			newPosition, output, hasOutput := node.pos.Execute(newInstruction, numpad)
			if _, ok := numpad[newPosition.String()]; !ok {
				continue
			}

			newLength := node.length
			if hasOutput {
				if output != rune(code[newLength]) {
					continue
				}
				newLength++
			}

			newCost := node.cost + calculateCost(newInstruction, node.instruction, depth, cache)
			pq.Push(NumpadNode{
				cost:        newCost,
				pos:         newPosition,
				instruction: newInstruction,
				length:      newLength,
			})
		}
	}

	panic("No solution found")
}

type DirpadNode struct {
	cost        int
	pos         D21Point
	instruction rune
	output      rune
	hasOutput   bool
}

func (n DirpadNode) Less(other DirpadNode) bool {
	return n.cost < other.cost
}

func calculateCost(target rune, previousInstruction rune, depth int, cache map[string]int) int {
	key := targetPrevDepthKey(target, previousInstruction, depth)
	if val, ok := cache[key]; ok {
		return val
	}

	if depth == 0 {
		cache[key] = 1
		return 1
	}

	dirpadKeysToPoints := make(map[rune]string)
	for k, v := range dirpad {
		dirpadKeysToPoints[v] = k
	}

	pq := utils.NewPriorityQueue[DirpadNode]()
	pq.Push(DirpadNode{
		cost:        0,
		pos:         ToD21Point(dirpadKeysToPoints[previousInstruction]),
		instruction: 'A',
		output:      '0',
		hasOutput:   false,
	})

	for !pq.IsEmpty() {
		node, _ := pq.Pop()
		if node.hasOutput && node.output == target {
			cache[key] = node.cost
			return node.cost
		}

		for _, newInstruction := range instructions {
			newPosition, newOutput, hasOutput := node.pos.Execute(newInstruction, dirpad)
			if _, ok := dirpad[newPosition.String()]; !ok {
				continue
			}

			if hasOutput && newOutput != target {
				continue
			}

			newCost := node.cost + calculateCost(newInstruction, node.instruction, depth-1, cache)
			pq.Push(DirpadNode{
				cost:        newCost,
				pos:         newPosition,
				instruction: newInstruction,
				output:      newOutput,
				hasOutput:   hasOutput,
			})
		}
	}

	panic("No solution found")
}

func posInstructionLengthKey(pos D21Point, instruction rune, length int) string {
	return pos.String() + "," + string(instruction) + "," + strconv.Itoa(length)
}

func targetPrevDepthKey(target rune, previousInstruction rune, depth int) string {
	return string(target) + "," + string(previousInstruction) + "," + strconv.Itoa(depth)
}
