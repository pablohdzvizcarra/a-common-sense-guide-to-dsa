package heap

import (
	"errors"
	"fmt"
)

type Heap struct {
	items []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Insert(key int) {
	h.items = append(h.items, key)
	h.heapifyUp(len(h.items) - 1)
}

func (h *Heap) ExtractMax() (int, error) {
	if len(h.items) == 0 {
		return 0, errors.New("heap is empty")
	}

	max := h.items[0]
	lastIndex := len(h.items) - 1
	h.items[0] = h.items[lastIndex]
	h.items = h.items[:lastIndex]

	if len(h.items) > 0 {
		h.heapifyDown(0)
	}

	return max, nil
}

func (h *Heap) heapifyDown(index int) {
	maxIndex := index
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2

		if leftChild < len(h.items) && h.items[leftChild] > h.items[maxIndex] {
			maxIndex = leftChild
		}

		if rightChild < len(h.items) && h.items[rightChild] > h.items[maxIndex] {
			maxIndex = rightChild
		}

		if maxIndex == index {
			break
		}

		h.swap(index, maxIndex)
		index = maxIndex
	}
}

func (h *Heap) heapifyUp(index int) {
	for h.hasParent(index) && h.parent(index) < h.items[index] {
		h.swap(h.parentIndex(index), index)
		index = h.parentIndex(index)
	}
}

func (h *Heap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *Heap) hasParent(index int) bool {
	return h.parentIndex(index) >= 0
}

func (h *Heap) parent(index int) int {
	return h.items[h.parentIndex(index)]
}

func (h *Heap) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}

func (h *Heap) PrintHeap() {
	fmt.Println(h.items)
}
