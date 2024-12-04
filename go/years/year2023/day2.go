package year2023

import (
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

type Day2 struct {
	Input []string
}

func (d Day2) Part1() string {
	sum := 0
	for _, line := range d.Input {
		parts := strings.Split(line, ":")
		_, err := strconv.Atoi(parts[0][5:])
		utils.PanicErr(err)

		parts = strings.Split(parts[1], ";")
		for _, part := range parts {
			_ = strings.Split(strings.TrimSpace(part), ",")
		}
	}
	return strconv.Itoa(sum)
}

func (d Day2) Part2() string {
	return "Not Implemented"
}
