package utils_test

import (
	"testing"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

type Item struct {
	value int
}

func (i Item) Less(j Item) bool {
	return i.value < j.value
}

func TestPriorityQueue(t *testing.T) {
	pq := utils.NewPriorityQueue[Item]()
	pq.Push(Item{2})
	pq.Push(Item{6})
	pq.Push(Item{1})
	pq.Push(Item{5})
	pq.Push(Item{3})
	pq.Push(Item{4})

	for i := range 6 {
		if pq.Len() != 6-i {
			t.Fatalf("expected 6, got %d", pq.Len())
		}

		item, ok := pq.Pop()
		if !ok {
			t.Fatal("expected item to exist")
		}

		if item.value != i+1 {
			t.Fatalf("expected 1, got %d", item.value)
		}
	}

	_, ok := pq.Pop()
	if ok {
		t.Fatal("expected no more items")
	}
}
