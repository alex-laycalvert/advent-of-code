package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const INPUT = "../input.txt"

func main() {
	// PART 1 ANSWER: 765748
	// PART 2 ANSWER: 27732508

	part1Answer := part1()
	fmt.Println("Part 1:", part1Answer)

	part2Answer := part2()
	fmt.Println("Part 2:", part2Answer)
}

func part1() int {
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for line := range readLines(INPUT) {
		parts := strings.Split(line, " ")
		leftNum, err := strconv.Atoi(parts[0])
		panicErr(err)
		leftList = append(leftList, leftNum)

		rightNum, err := strconv.Atoi(parts[len(parts)-1])
		panicErr(err)
		rightList = append(rightList, rightNum)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	sum := 0
	for i := 0; i < len(leftList); i++ {
		sum += intAbs(leftList[i] - rightList[i])
	}

	return sum
}

func part2() int {
	leftList := make([]int, 0)
	rightFreq := make(map[int]int)

	for line := range readLines(INPUT) {
		parts := strings.Split(line, " ")
		leftNum, err := strconv.Atoi(parts[0])
		panicErr(err)
		leftList = append(leftList, leftNum)

		rightNum, err := strconv.Atoi(parts[len(parts)-1])
		panicErr(err)
		if _, ok := rightFreq[rightNum]; ok {
			rightFreq[rightNum]++
		} else {
			rightFreq[rightNum] = 1
		}
		fmt.Println(rightFreq[rightNum])
	}

	sum := 0
	for _, num := range leftList {
		freq, ok := rightFreq[num]
		if !ok {
			continue
		}
		sum += num * freq
	}

	return sum
}

func panicErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func intAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func readLines(filename string) <-chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open(filename)
		panicErr(err)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch
}
