package year2024

import (
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

type Day10 struct {
	Input []string
}

func (d Day10) Part1() string {
	topographicMap := make([][]int, len(d.Input))
	for i, line := range d.Input {
		topographicMap[i] = make([]int, len(line))
		for j, char := range line {
			num, err := strconv.Atoi(string(char))
			utils.PanicErr(err)
			topographicMap[i][j] = num
		}
	}

	answer := 0
	for row := range topographicMap {
		for col := range topographicMap[row] {
			if topographicMap[row][col] != 0 {
				continue
			}
			ans := searchTrailsUnique(topographicMap, row, col, make(map[string]bool))
			answer += ans
		}
	}

	return strconv.Itoa(answer)
}

func (d Day10) Part2() string {
	topographicMap := make([][]int, len(d.Input))
	for i, line := range d.Input {
		topographicMap[i] = make([]int, len(line))
		for j, char := range line {
			num, err := strconv.Atoi(string(char))
			utils.PanicErr(err)
			topographicMap[i][j] = num
		}
	}

	answer := 0
	for row := range topographicMap {
		for col := range topographicMap[row] {
			if topographicMap[row][col] != 0 {
				continue
			}
			ans := searchTrails(topographicMap, row, col)
			answer += ans
		}
	}

	return strconv.Itoa(answer)
}

func searchTrailsUnique(topographicMap [][]int, row int, col int, visited map[string]bool) int {
	if visited[utils.PosToString(row, col)] {
		return 0
	}
	height := topographicMap[row][col]
	if height == 9 {
		visited[utils.PosToString(row, col)] = true
		return 1
	}

	numTrails := 0

	if row > 0 && topographicMap[row-1][col] == height+1 {
		numTrails += searchTrailsUnique(topographicMap, row-1, col, visited)
	}

	if row < len(topographicMap)-1 && topographicMap[row+1][col] == height+1 {
		numTrails += searchTrailsUnique(topographicMap, row+1, col, visited)
	}

	if col > 0 && topographicMap[row][col-1] == height+1 {
		numTrails += searchTrailsUnique(topographicMap, row, col-1, visited)
	}

	if col < len(topographicMap[row])-1 && topographicMap[row][col+1] == height+1 {
		numTrails += searchTrailsUnique(topographicMap, row, col+1, visited)
	}

	visited[utils.PosToString(row, col)] = true
	return numTrails
}

func searchTrails(topographicMap [][]int, row int, col int) int {
	height := topographicMap[row][col]
	if height == 9 {
		return 1
	}

	numTrails := 0

	if row > 0 && topographicMap[row-1][col] == height+1 {
		numTrails += searchTrails(topographicMap, row-1, col)
	}

	if row < len(topographicMap)-1 && topographicMap[row+1][col] == height+1 {
		numTrails += searchTrails(topographicMap, row+1, col)
	}

	if col > 0 && topographicMap[row][col-1] == height+1 {
		numTrails += searchTrails(topographicMap, row, col-1)
	}

	if col < len(topographicMap[row])-1 && topographicMap[row][col+1] == height+1 {
		numTrails += searchTrails(topographicMap, row, col+1)
	}

	return numTrails
}
