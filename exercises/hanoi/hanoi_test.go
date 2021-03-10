package hanoi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilgustavo/godata/exercises/hanoi"
	"github.com/wilgustavo/godata/pkg/list/stack"
)

func TestHanoi(t *testing.T) {
	t.Run("Hanoi with stacks", func(t *testing.T) {
		from := stack.NewStack()
		aux := stack.NewStack()
		to := stack.NewStack()
		from.Push(8)
		from.Push(5)
		from.Push(3)

		assert.Equal(t, []interface{}{3, 5, 8}, from.ToSlice())

		hanoi.Hanoi(3, from, aux, to)
		assert.Equal(t, []interface{}{3, 5, 8}, to.ToSlice())
		assert.True(t, aux.IsEmpty())
		assert.True(t, from.IsEmpty())
	})
}
