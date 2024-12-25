package year2024

import (
	"strconv"
)

// ANSWERS:
//
// Part 1: 2586
//
// Part 2: N/A
type Day25 struct {
	Input []string
}

func (d Day25) Part1() string {
	locks := make([][]int, 0)
	keys := make([][]int, 0)

	keyOrLock := make([]int, 5)
	for i, line := range d.Input {
		if line == "" {
			if d.Input[i-1] == "#####" {
				keys = append(keys, keyOrLock)
			} else {
				locks = append(locks, keyOrLock)
			}
			keyOrLock = make([]int, 5)
			continue
		}

		for j, c := range line {
			if c == '#' {
				keyOrLock[j]++
			}
		}
	}

	if d.Input[len(d.Input)-1] == "#####" {
		keys = append(keys, keyOrLock)
	} else {
		locks = append(locks, keyOrLock)
	}

	answer := 0
	for _, lock := range locks {
		for _, key := range keys {
			works := true
			for i := range 5 {
				if key[i]+lock[i] > 7 {
					works = false
					break
				}
			}
			if works {
				answer++
			}
		}
	}

	return strconv.Itoa(answer)
}

func (d Day25) Part2() string {
	return "Not Implemented"
}
