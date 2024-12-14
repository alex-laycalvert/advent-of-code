package year2024

import (
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 228
//
// Part 2: 766
type Day8 struct {
	Input []string
}

// avg ~0.04ms per iter over 10k iters
func (d Day8) Part1() string {
	answer := 0

	antennas := make([]utils.Pos, 0)
	width, height := len(d.Input[0]), len(d.Input)

	for r, line := range d.Input {
		for c, char := range line {
			if char == '.' {
				continue
			}
			antennas = append(antennas, utils.Pos{X: c, Y: r, Ch: char})
		}
	}

	antinodes := make(map[string]bool)
	for i, ant1 := range antennas {
		for j := i + 1; j < len(antennas); j++ {
			ant2 := antennas[j]
			if ant1.Ch != ant2.Ch {
				continue
			}
			dy := ant2.Y - ant1.Y
			dx := ant2.X - ant1.X

			if ant1.X-dx >= 0 && ant1.Y-dy >= 0 && ant1.X-dx < width && ant1.Y-dy < height {
				if !antinodes[utils.PosToString(ant1.X-dx, ant1.Y-dy)] {
					answer++
				}
				antinodes[utils.PosToString(ant1.X-dx, ant1.Y-dy)] = true
			}

			if ant2.X+dx >= 0 && ant2.Y+dy >= 0 && ant2.X+dx < width && ant2.Y+dy < height {
				if !antinodes[utils.PosToString(ant2.X+dx, ant2.Y+dy)] {
					answer++
				}
				antinodes[utils.PosToString(ant2.X+dx, ant2.Y+dy)] = true
			}
		}
	}

	return strconv.Itoa(answer)
}

// avg ~0.12ms per iter over 10k iters
func (d Day8) Part2() string {
	answer := 0

	antennas := make([]utils.Pos, 0)
	width, height := len(d.Input[0]), len(d.Input)

	for r, line := range d.Input {
		for c, char := range line {
			if char == '.' {
				continue
			}
			antennas = append(antennas, utils.Pos{X: c, Y: r, Ch: char})
		}
	}

	antinodes := make(map[string]bool)
	for i, ant1 := range antennas {
		for j := i + 1; j < len(antennas); j++ {
			ant2 := antennas[j]
			if ant1.Ch != ant2.Ch {
				continue
			}
			dy := ant2.Y - ant1.Y
			dx := ant2.X - ant1.X

			currX := ant1.X
			currY := ant1.Y
			for currX >= 0 && currX < width && currY >= 0 && currY < height {
				key := utils.PosToString(currX, currY)
				if !antinodes[key] {
					answer++
				}
				antinodes[key] = true
				currX -= dx
				currY -= dy
			}

			currX = ant2.X
			currY = ant2.Y
			for currX >= 0 && currX < width && currY >= 0 && currY < height {
				key := utils.PosToString(currX, currY)
				if !antinodes[key] {
					answer++
				}
				antinodes[key] = true
				currX += dx
				currY += dy
			}
		}
	}

	return strconv.Itoa(answer)
}
