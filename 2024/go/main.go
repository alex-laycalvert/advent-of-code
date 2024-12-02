package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/2024/days"
	"github.com/alex-laycalvert/advent-of-code/2024/utils"
)

const TIMING_ITERATIONS = 10_000

func main() {
	if len(os.Args) < 2 {
		panic("Must provide a day number")
	}
	dayNum, err := strconv.Atoi(os.Args[1])
	utils.PanicErr(err)

	input := "../inputs/" + strconv.Itoa(dayNum) + ".txt"

	day := days.NewDay(dayNum, input)

	part1, ms := utils.TimeFunc(day.Part1, TIMING_ITERATIONS)
	fmt.Printf("Part 1: %s - Avg Time: %vms\n", part1, ms)

	part2, ms := utils.TimeFunc(day.Part2, TIMING_ITERATIONS)
	fmt.Printf("Part 2: %s - Avg Time: %vms\n", part2, ms)
}
