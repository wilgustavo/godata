package heap

import "fmt"

type heap struct {
	list   []int
	length int
}

// BinaryHeap is operations on a binary heap
type BinaryHeap interface {
	Insert(item int)
	Extract() (int, error)
	List() []int
}

// NewBinaryHeap returns a new binary heap
func NewBinaryHeap() BinaryHeap {
	return &heap{
		list:   []int{},
		length: 0,
	}
}

func (h *heap) Insert(item int) {
	h.list = append(h.list, item)
	h.bubbleup(h.length)
	h.length++
}

func (h *heap) Extract() (int, error) {
	if h.length == 0 {
		return 0, fmt.Errorf("Heap is empty")
	}
	first := 0
	last := h.length - 1
	item := h.list[first]
	h.swap(first, last)
	h.list = h.list[:last]
	h.length = last
	h.bubbledown(first)
	return item, nil
}

func (h *heap) List() []int {
	return h.list
}

func (h *heap) bubbleup(pos int) {
	cur := pos
	for cur > 0 {
		parent := parentPosition(cur)
		if h.list[cur] < h.list[parent] {
			h.swap(cur, parent)
		} else {
			break
		}
		cur = parent
	}
}

func (h *heap) bubbledown(pos int) {
	for pos < h.length {
		left := pos*2 + 1
		right := pos*2 + 2
		minorPos := pos

		if left < h.length && h.list[left] < h.list[minorPos] {
			minorPos = left
		}
		if right < h.length && h.list[right] < h.list[minorPos] {
			minorPos = right
		}
		if minorPos != pos {
			h.list[pos], h.list[minorPos] = h.list[minorPos], h.list[pos]
		} else {
			break
		}
		pos = minorPos
	}
}

func (h *heap) swap(a, b int) {
	h.list[a], h.list[b] = h.list[b], h.list[a]
}

func parentPosition(current int) int {
	return (current - 1) / 2
}
