package heap

type Heap struct {
	data []int
}

func NewHeap() Heap {
	return Heap{
		data: make([]int, 0),
	}
}

func (h *Heap) RootNode() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	return h.data[0], true
}

func (h *Heap) LastNode() interface{} {
	if len(h.data) == 0 {
		return nil
	}

	return h.data[len(h.data)-1]
}

func (h *Heap) leftChildIndex(index int) int {
	return (index * 2) + 1
}

func (h *Heap) rightChildIndex(index int) int {
	return (index * 2) + 2
}

func (h *Heap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *Heap) Insert(value int) {
	// we insert the value in the last node, remember the last node is the last element in the array
	h.data = append(h.data, value)

	// create the index for the new value inserted
	newNodeIndex := len(h.data) - 1

	for newNodeIndex > 0 && h.data[newNodeIndex] > h.data[h.parentIndex(newNodeIndex)] {
		parentIndex := h.parentIndex(newNodeIndex)

		tempParentValue := h.data[parentIndex]
		h.data[parentIndex] = h.data[newNodeIndex]
		h.data[newNodeIndex] = tempParentValue

		newNodeIndex = parentIndex
	}
}

func (h *Heap) Pop() (int, bool) {
	if valueToDelete, ok := h.RootNode(); ok {
		h.data[0] = h.data[len(h.data)-1]
		trickleNodeIndex := 0

		for h.hasGreaterChild(trickleNodeIndex) {
			largerChildIndex := h.findLargerChildIndex(trickleNodeIndex)

			tempTricklyValue := h.data[trickleNodeIndex]
			h.data[trickleNodeIndex] = h.data[largerChildIndex]
			h.data[largerChildIndex] = tempTricklyValue

			trickleNodeIndex = largerChildIndex
		}

		return valueToDelete, true
	}

	return 0, false
}

func (h *Heap) hasGreaterChild(index int) bool {
	firstCondition := h.leftChildIndex(index) <= len(h.data) && h.data[h.leftChildIndex(index)] > h.data[index]
	secondCondition := h.rightChildIndex(index) <= len(h.data) && h.data[h.rightChildIndex(index)] > h.data[index]

	return firstCondition || secondCondition
}

func (h *Heap) findLargerChildIndex(index int) int {
	// if h.data[h.RightChildIndex(index)]
	if h.rightChildIndex(index) > len(h.data) {
		return h.leftChildIndex(index)
	}

	if h.data[h.rightChildIndex(index)] > h.data[h.leftChildIndex(index)] {
		return h.rightChildIndex(index)
	} else {
		return h.leftChildIndex(index)
	}
}
