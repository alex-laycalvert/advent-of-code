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
	shortestNumericPaths := getAllKeyboardNeighborPaths(numericKeypadNeighbors)
	shortestDirectionalPaths := getAllKeyboardNeighborPaths(directionalKeypadNeighbors)

	answer := 0
	for _, line := range d.Input {
		minPath := getShortestPathForRobots(line, 2, shortestNumericPaths, shortestDirectionalPaths)
		numericPart, err := strconv.Atoi(strings.TrimRight(strings.TrimLeft(line, "0"), "A"))
		utils.PanicErr(err)
		answer += len(minPath) * numericPart
	}

	return strconv.Itoa(answer)
}

func (d Day21) Part2() string {
	shortestNumericPaths := getAllKeyboardNeighborPaths(numericKeypadNeighbors)
	shortestDirectionalPaths := getAllKeyboardNeighborPaths(directionalKeypadNeighbors)

	answer := 0
	for _, line := range d.Input {
		minPath := getShortestPathForRobots(line, 25, shortestNumericPaths, shortestDirectionalPaths)
		numericPart, err := strconv.Atoi(strings.TrimRight(strings.TrimLeft(line, "0"), "A"))
		utils.PanicErr(err)
		answer += len(minPath) * numericPart
	}

	return strconv.Itoa(answer)
}

type RobotArmPath struct {
	path  string
	count int
}

func (p RobotArmPath) Less(other RobotArmPath) bool {
	return len(p.path) > len(other.path)
}

func getShortestPathForRobots(initial string, numberOfRobots int, numericPaths map[rune]map[rune][]string, directionalPaths map[rune]map[rune][]string) string {
	q := utils.NewPriorityQueue[RobotArmPath]()
	q.Push(RobotArmPath{
		path:  initial,
		count: 0,
	})

	visited := make(map[string]string)
	minPath := ""
	cache := make(map[string][]string)
	for !q.IsEmpty() {
		current, _ := q.Pop()
		if minPath != "" && len(current.path) > len(minPath) {
			continue
		}

		if _, ok := visited[current.path]; ok {
			continue
		}

		if current.count == numberOfRobots+1 {
			minPath = current.path
			continue
		}
		visited[current.path] = current.path

		var neighbors []string
		if current.count == 0 {
			neighbors = getShortestPaths(current.path, numericPaths, cache)
		} else {
			neighbors = getShortestPaths(current.path, directionalPaths, cache)
		}
		for _, neighbor := range neighbors {
			q.Push(RobotArmPath{
				path:  neighbor,
				count: current.count + 1,
			})
		}
	}

	return minPath
}

func getShortestPaths(input string, shortestKeypadPaths map[rune]map[rune][]string, cache map[string][]string) []string {
	paths := getShortestPathsInternal('A', input, shortestKeypadPaths, cache)
	return paths
}

func getShortestPathsInternal(start rune, input string, shortestKeypadPaths map[rune]map[rune][]string, cache map[string][]string) []string {
	key := string(start) + input
	if cached, ok := cache[key]; ok {
		return cached
	}

	if len(input) == 0 {
		cache[key] = []string{""}
		return []string{""}
	}

	end := rune(input[0])
	shortestPaths := shortestKeypadPaths[start][end]
	if len(shortestPaths) == 0 {
		shortestPaths = []string{""}
	}
	paths := getShortestPathsInternal(rune(input[0]), input[1:], shortestKeypadPaths, cache)

	allPaths := make([]string, 0, len(paths)*len(shortestPaths))
	for _, shortestPath := range shortestPaths {
		for _, path := range paths {
			newPath := shortestPath + "A" + path
			allPaths = append(allPaths, newPath)
		}
	}

	cache[key] = allPaths
	return allPaths
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
	q := utils.NewPriorityQueue[Path]()
	q.Push(Path{[]rune{'0', start}, "", ""})
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
			q.Push(Path{neighbor, current.points, current.path})
		}
	}

	return paths
}

type Path struct {
	value  []rune
	points string
	path   string
}

func (p Path) Less(other Path) bool {
	return len(p.path) > len(other.path)
}
