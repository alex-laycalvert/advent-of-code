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
// Part 2:
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
var numericKeypadNeighbors = map[rune][][]rune{
	'A': {{'<', '0'}, {'^', '3'}},
	'0': {{'>', 'A'}, {'^', '2'}},
	'1': {{'>', '2'}, {'^', '4'}},
	'2': {{'v', '0'}, {'<', '1'}, {'>', '3'}, {'^', '5'}},
	'3': {{'v', 'A'}, {'<', '2'}, {'^', '6'}},
	'4': {{'v', '1'}, {'>', '5'}, {'^', '7'}},
	'5': {{'v', '2'}, {'<', '4'}, {'>', '6'}, {'^', '8'}},
	'6': {{'v', '3'}, {'<', '5'}, {'^', '9'}},
	'7': {{'v', '4'}, {'>', '8'}},
	'8': {{'v', '5'}, {'<', '7'}, {'>', '9'}},
	'9': {{'v', '6'}, {'<', '8'}},
}

/*
*     +---+---+
*     | ^ | A |
* +---+---+---+
* | < | v | > |
* +---+---+---+
 */
var directionalKeypadNeighbors = map[rune][][]rune{
	'A': {{'<', '^'}, {'v', '>'}},
	'^': {{'>', 'A'}, {'v', 'v'}},
	'v': {{'^', '^'}, {'<', '<'}, {'>', '>'}},
	'<': {{'>', 'v'}},
	'>': {{'<', 'v'}, {'^', 'A'}},
}

func (d Day21) Part1() string {
	// shortestNumericPaths := getAllKeyboardNeighborPaths(numericKeypadNeighbors)
	// shortestDirectionalPaths := getAllKeyboardNeighborPaths(directionalKeypadNeighbors)

	answer := 0
	input := d.Input[0]

	_ = solve(input)

	// numericPart, err := strconv.Atoi(strings.TrimRight(strings.TrimLeft(input, "0"), "A"))
	// utils.PanicErr(err)
	// answer += numericPart * len(minPath)
	return strconv.Itoa(answer + 1)
}

func (d Day21) Part2() string {
	return "Not Implemented"
}

type Path struct {
	value byte
	index int
	cost  int
}

func (p Path) Less(other Path) bool {
	return p.cost > other.cost
}

func solve(input string) string {
	pq := utils.NewPriorityQueue[Path]()
	pq.Push(Path{'A', 0, 0})
	visited := make(map[string]string)
	for current, ok := pq.Pop(); ok; current, ok = pq.Pop() {
		key := string(current.value) + ":" + strconv.Itoa(current.index)

		if _, ok := visited[key]; ok {
			continue
		}
		visited[key] = key

		if current.index == len(input) {
			return ""
		}

		// Pressing action
		if current.value == input[current.index] {
			pq.Push(Path{
				value: current.value,
				index: current.index + 1,
			})
		}

		// Moving
		for _, neighbor := range numericKeypadNeighbors[rune(current.value)] {
			pq.Push(Path{
				value: byte(neighbor[1]),
				index: current.index,
			})
		}
	}
	return ""
}

func calculateCost(target rune, current rune, depth int) int {
	if depth == 0 {
		return 1
	}

	pq := utils.NewPriorityQueue[Path]()
	pq.Push(Path{
		value: byte(current),
		cost:  0,
		index: 0,
	})

	for current, ok := pq.Pop(); ok; current, ok = pq.Pop() {
		if current.value == byte(target) {
			return current.cost
		}

		for _, neighbor := range directionalKeypadNeighbors[rune(current.value)] {
			pq.Push(Path{
				value: byte(neighbor[1]),
				cost:  current.cost + 1,
			})
		}
	}

	return 0
}

func getAllKeyboardNeighborPaths(neighbors map[rune][][]rune) map[rune]map[rune][]string {
	shortestPaths := make(map[rune]map[rune][]string)
	for start := range neighbors {
		shortestPaths[start] = make(map[rune][]string)
		for end := range neighbors {
			if start == end {
				continue
			}
			shortestPaths[start][end] = getKeyboardNeighborPaths(start, end, neighbors)
		}
	}
	return shortestPaths
}

func getKeyboardNeighborPaths(start rune, end rune, neighbors map[rune][][]rune) []string {
	q := utils.NewPriorityQueue[KeypadPath]()
	q.Push(KeypadPath{[]rune{'0', start}, "", ""})
	paths := make([]string, 0)
	bestPath := 999
	for !q.IsEmpty() {
		current, _ := q.Pop()
		if strings.ContainsRune(current.points, current.value[1]) {
			continue
		}
		if current.value[0] != '0' {
			current.path += string(current.value[0])
		}
		current.points += string(current.value[1])
		if current.value[1] == end {
			if len(current.path) < bestPath {
				bestPath = len(current.path)
				paths = []string{current.path}
			} else if len(current.path) == bestPath {
				paths = append(paths, current.path)
			}
			continue
		}
		for _, neighbor := range neighbors[current.value[1]] {
			q.Push(KeypadPath{neighbor, current.points, current.path})
		}
	}

	return paths
}

type KeypadPath struct {
	value  []rune
	points string
	path   string
}

func (p KeypadPath) Less(other KeypadPath) bool {
	return len(p.path) > len(other.path)
}
