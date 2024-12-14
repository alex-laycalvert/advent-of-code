package utils

import (
	"bufio"
	"os"
	"strconv"
	"time"
)

func ReadFile(filename string) []string {
	lines := []string{}
	for line := range ReadLines(filename) {
		lines = append(lines, line)
	}
	return lines
}

func ReadLines(filename string) <-chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open(filename)
		PanicErr(err)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch
}

func IntAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func PanicErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func RemoveIndex[S any](slice []S, index int) []S {
	if index < 0 || index >= len(slice) {
		return slice
	}

	newSlice := make([]S, len(slice)-1)
	copy(newSlice, slice[:index])
	copy(newSlice[index:], slice[index+1:])
	return newSlice
}

// TimeFunc returns result of the fn and the avg time taken to
// run a function over `iterations` times in milliseconds
func TimeFunc(fn func() string, iterations int) (string, float64) {
	var ans string

	var sum int64 = 0
	for i := range iterations {
		now := time.Now()
		newAns := fn()
		if i != 0 && newAns != ans {
			panic("Function returned different results")
		}
		ans = newAns
		sum += time.Since(now).Nanoseconds()
	}

	return ans, float64(sum) / float64(iterations) / 1_000_000
}

type Pos struct {
	X  int
	Y  int
	Ch rune
}

func NewPos(x, y int, ch rune) Pos {
	return Pos{x, y, ch}
}

func (p Pos) String(withChar bool) string {
	key := PosToString(p.X, p.Y)
	if withChar {
		key += "," + string(p.Ch)
	}
	return key
}

func PosToString(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}
