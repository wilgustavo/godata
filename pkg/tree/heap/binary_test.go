package heap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilgustavo/godata/pkg/tree/heap"
)

func TestInsert(t *testing.T) {
	t.Run("Insert a item", func(t *testing.T) {

		h := heap.NewBinaryHeap()
		h.Insert(10)
		assert.Equal(t, []int{10}, h.List())

		h.Insert(5)
		assert.Equal(t, []int{5, 10}, h.List())

		h.Insert(3)
		assert.Equal(t, []int{3, 10, 5}, h.List())

	})
	t.Run("Should insert multiple items", func(t *testing.T) {

		h := heap.NewBinaryHeap()
		insertElements(h, 7, 1, 3, 8, 9, 5)
		assert.Equal(t, []int{1, 7, 3, 8, 9, 5}, h.List())

	})
}

func TestExtract(t *testing.T) {
	t.Run("Extract a item", func(t *testing.T) {
		h := heap.NewBinaryHeap()
		h.Insert(7)
		h.Insert(1)
		item, err := h.Extract()
		assert.Equal(t, 1, item)
		assert.NoError(t, err)

		item, err = h.Extract()
		assert.Equal(t, 7, item)
		assert.NoError(t, err)

		_, err = h.Extract()
		assert.Error(t, err)

	})
	t.Run("Should extract several elements", func(t *testing.T) {
		h := heap.NewBinaryHeap()
		insertElements(h, 7, 1, 3, 8, 9, 5)

		for _, expected := range []int{1, 3, 5, 7, 8, 9} {
			item, err := h.Extract()
			assert.NoError(t, err)
			assert.Equal(t, expected, item)
		}

		_, err := h.Extract()
		assert.Error(t, err)

	})

}

func insertElements(heap heap.BinaryHeap, items ...int) {
	for _, item := range items {
		heap.Insert(item)
	}
}
