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

func (p Pos) Equals(other Pos, withChar bool) bool {
	return p.X == other.X && p.Y == other.Y && (!withChar || p.Ch == other.Ch)
}

func (p Pos) TurnAround() Pos {
	newCh := p.Ch
	switch p.Ch {
	case '^':
		newCh = 'v'
		break
	case '>':
		newCh = '<'
		break
	case 'v':
		newCh = '^'
		break
	case '<':
		newCh = '>'
		break
	}
	return Pos{X: p.X, Y: p.Y, Ch: newCh}
}

func (p Pos) TurnRight() Pos {
	newCh := p.Ch
	switch p.Ch {
	case '^':
		newCh = '>'
		break
	case '>':
		newCh = 'v'
		break
	case 'v':
		newCh = '<'
		break
	case '<':
		newCh = '^'
		break
	}
	return Pos{X: p.X, Y: p.Y, Ch: newCh}
}

func (p Pos) TurnLeft() Pos {
	newCh := p.Ch
	switch p.Ch {
	case '^':
		newCh = '<'
		break
	case '>':
		newCh = '^'
		break
	case 'v':
		newCh = '>'
		break
	case '<':
		newCh = 'v'
		break
	}
	return Pos{X: p.X, Y: p.Y, Ch: newCh}
}

func (p Pos) MoveForward() Pos {
	newPos := Pos{X: p.X, Y: p.Y, Ch: p.Ch}
	switch p.Ch {
	case '^':
		newPos.Y--
		break
	case '>':
		newPos.X++
		break
	case 'v':
		newPos.Y++
		break
	case '<':
		newPos.X--
		break
	}
	return newPos
}

func PosToString(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

// given a string, parses out all numbers and returns them as a slice of ints.
func ParseNumbers(s string) []int {
	numbers := []int{}

	currentNum := ""
	isNegative := false
	for _, ch := range s {
		if ch == '-' && !isNegative {
			isNegative = true
			currentNum = string(ch) + currentNum
			continue
		}

		if ch >= '0' && ch <= '9' {
			currentNum += string(ch)
			continue
		}

		if currentNum != "" {
			num, _ := strconv.Atoi(currentNum)
			numbers = append(numbers, num)
			currentNum = ""
			isNegative = false
		}
	}

	if currentNum != "" {
		num, err := strconv.Atoi(currentNum)
		PanicErr(err)
		numbers = append(numbers, num)
	}
	return numbers
}

type Queue[T any] struct {
	queue []T
	size  int
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{queue: []T{}, size: 0}
}

func (q *Queue[T]) Enqueue(val T) {
	q.size++
	q.queue = append(q.queue, val)
}

func (q *Queue[T]) Dequeue() T {
	var val T
	if q.size == 0 {
		return val
	}

	val = q.queue[0]
	q.queue = q.queue[1:]
	q.size--
	return val
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

type Stack[T any] struct {
	stack []T
	size  int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{stack: []T{}, size: 0}
}

func (q *Stack[T]) Push(val T) {
	q.size++
	q.stack = append(q.stack, val)
}

func (q *Stack[T]) Pop() T {
	var val T
	if q.size == 0 {
		return val
	}

	val = q.stack[q.size-1]
	q.stack = q.stack[:q.size-1]
	q.size--
	return val
}

func (q *Stack[T]) IsEmpty() bool {
	return q.size == 0
}

type QueueItem[T any] interface {
	Less(other T) bool
}

// PriorityQueue represents a generic priority queue
type PriorityQueue[T QueueItem[T]] struct {
	items []T
}

// New creates a new priority queue
func NewPriorityQueue[T QueueItem[T]]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the queue
func (pq *PriorityQueue[T]) Push(item T) {
	pq.items = append(pq.items, item)
	pq.up(len(pq.items) - 1)
}

// Pop removes and returns the highest priority item
func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if len(pq.items) == 0 {
		var zero T
		return zero, false
	}

	item := pq.items[0]
	last := len(pq.items) - 1
	pq.items[0] = pq.items[last]
	pq.items = pq.items[:last]
	if len(pq.items) > 0 {
		pq.down(0)
	}
	return item, true
}

// Len returns the number of items in the queue
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.items)
}

// Internal heap operations
func (pq *PriorityQueue[T]) up(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if pq.items[index].Less(pq.items[parent]) {
			pq.items[index], pq.items[parent] = pq.items[parent], pq.items[index]
			index = parent
		} else {
			break
		}
	}
}

func (pq *PriorityQueue[T]) down(index int) {
	for {
		smallest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < len(pq.items) && pq.items[left].Less(pq.items[smallest]) {
			smallest = left
		}
		if right < len(pq.items) && pq.items[right].Less(pq.items[smallest]) {
			smallest = right
		}

		if smallest == index {
			break
		}

		pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
		index = smallest
	}
}
