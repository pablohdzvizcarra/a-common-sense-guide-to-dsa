package heap

import (
	"testing"
)

func TestHeapCreation(t *testing.T) {
	heap := NewHeap()
	heap.Insert(100)
	heap.Insert(80)
	heap.Insert(101)
	heap.Insert(70)
	heap.Insert(60)
	heap.Insert(31)

	maxElem, _ := heap.ExtractMax()
	heap.PrintHeap()

	if maxElem != 101 {
		t.Errorf("must 101 but have %d", maxElem)
	}
}
