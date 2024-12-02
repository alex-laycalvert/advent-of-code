package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/days"
	"github.com/alex-laycalvert/advent-of-code/utils"
)

const DEFAULT_TIMING_ITERATIONS = 10_000

func main() {
	if len(os.Args) < 3 {
		panic("Must provide a year and day")
	}
	yearNum, err := strconv.Atoi(os.Args[1])
	utils.PanicErr(err)
	dayNum, err := strconv.Atoi(os.Args[2])
	utils.PanicErr(err)

	timingIterations := DEFAULT_TIMING_ITERATIONS
	if len(os.Args) > 3 {
		timingIterations, err = strconv.Atoi(os.Args[3])
		utils.PanicErr(err)
	}

	inputPath := fmt.Sprintf("../inputs/%d/%d.txt", yearNum, dayNum)
	utils.PanicErr(err)
	input := utils.ReadFile(inputPath)

	day := days.NewDay(yearNum, dayNum, input)

	part1, ms := utils.TimeFunc(day.Part1, timingIterations)
	fmt.Printf("Part 1: %s - Avg Time: %vms\n", part1, ms)

	part2, ms := utils.TimeFunc(day.Part2, timingIterations)
	fmt.Printf("Part 2: %s - Avg Time: %vms\n", part2, ms)
}
