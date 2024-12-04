package year2024

import "strconv"

// ANSWERS:
//
// Part 1: 2551
//
// Part 2: 1985
type Day4 struct {
	Input []string
}

func (d Day4) Part1() string {
	return d.part1_naive()
}

func (d *Day4) part1_naive() string {
	count := 0
	for row, line := range d.Input {
		for col, char := range line {
			if char != 'X' {
				continue
			}

			if row-3 >= 0 {
				if d.Input[row-1][col] == 'M' && d.Input[row-2][col] == 'A' && d.Input[row-3][col] == 'S' {
					count++
				}
			}

			if row+3 < len(d.Input) {
				if d.Input[row+1][col] == 'M' && d.Input[row+2][col] == 'A' && d.Input[row+3][col] == 'S' {
					count++
				}
			}

			if col-3 >= 0 {
				if line[col-1] == 'M' && line[col-2] == 'A' && line[col-3] == 'S' {
					count++
				}
			}

			if col+3 < len(line) {
				if line[col+1] == 'M' && line[col+2] == 'A' && line[col+3] == 'S' {
					count++
				}
			}

			if row-3 >= 0 && col-3 >= 0 {
				if d.Input[row-1][col-1] == 'M' && d.Input[row-2][col-2] == 'A' && d.Input[row-3][col-3] == 'S' {
					count++
				}
			}

			if row-3 >= 0 && col+3 < len(line) {
				if d.Input[row-1][col+1] == 'M' && d.Input[row-2][col+2] == 'A' && d.Input[row-3][col+3] == 'S' {
					count++
				}
			}

			if row+3 < len(d.Input) && col-3 >= 0 {
				if d.Input[row+1][col-1] == 'M' && d.Input[row+2][col-2] == 'A' && d.Input[row+3][col-3] == 'S' {
					count++
				}
			}

			if row+3 < len(d.Input) && col+3 < len(line) {
				if d.Input[row+1][col+1] == 'M' && d.Input[row+2][col+2] == 'A' && d.Input[row+3][col+3] == 'S' {
					count++
				}
			}
		}
	}
	return strconv.Itoa(count)
}

func (d Day4) Part2() string {
	return d.part2_naive()
}

func (d *Day4) part2_naive() string {
	count := 0
	for row, line := range d.Input {
		if row+2 >= len(d.Input) {
			break
		}
		for col, char := range line {
			if col+2 >= len(line) {
				break
			}
			if d.Input[row+1][col+1] != 'A' {
				continue
			}

			if !((char == 'M' && d.Input[row+2][col+2] == 'S') ||
				(char == 'S' && d.Input[row+2][col+2] == 'M')) {
				continue
			}

			if !((line[col+2] == 'M' && d.Input[row+2][col] == 'S') ||
				(line[col+2] == 'S' && d.Input[row+2][col] == 'M')) {
				continue
			}

            count++
			// ...
			// .A.
			// ...
		}
	}
	return strconv.Itoa(count)
}
