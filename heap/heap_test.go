package heap

import (
	"testing"
)

func TestHeapCreation(t *testing.T) {
	heap := NewHeap()

	lastNode := heap.LastNode()

	if lastNode != nil {
		t.Errorf("have %v must %v", lastNode, nil)
	}
}

func TestInsertSixElements(t *testing.T) {
	heap := NewHeap()

	heap.Insert(100)
	heap.Insert(90)
	heap.Insert(60)
	heap.Insert(1)
	heap.Insert(10)
	heap.Insert(92)
	heap.Insert(31)
	heap.Insert(98)

	if len(heap.data) != 8 {
		t.Errorf("have %d must %d", len(heap.data), 5)
	}
}

func TestPopFirstIndex(t *testing.T) {
	heap := NewHeap()

	heap.Insert(100)
	heap.Insert(90)
	heap.Insert(60)
	heap.Insert(1)
	heap.Insert(10)
	heap.Insert(92)
	heap.Insert(31)
	heap.Insert(98)

	t.Log(heap.data)

	if value, ok := heap.Pop(); ok {
		if value != 100 {
			t.Errorf("must %d have %d", 100, value)
		}
	}
}
