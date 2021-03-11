package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilgustavo/godata/pkg/list/queue"
)

func TestNewQueue(t *testing.T) {
	t.Run("Should create a new empty queue", func(t *testing.T) {
		q := queue.NewQueue()
		assert.True(t, q.IsEmpty())
	})
	t.Run("Should get error when Enqueue an empty queue", func(t *testing.T) {
		q := queue.NewQueue()
		_, err := q.Dequeue()
		assert.Error(t, err)
	})
}

func TestEnqueue(t *testing.T) {
	t.Run("Should dont get an empty queue when Enqueue 2 elements", func(t *testing.T) {
		q := queue.NewQueue()
		q.Enqueue("Hello")
		q.Enqueue("World")
		assert.False(t, q.IsEmpty())
	})
}

func TestDequeue(t *testing.T) {
	t.Run("Should dequeue an element when enqueue an element", func(t *testing.T) {
		q := queue.NewQueue()
		q.Enqueue("Hola")
		q.Enqueue("Mundo")
		q.Dequeue()
		item, err := q.DequeueString()
		assert.NoError(t, err)
		assert.Equal(t, "Mundo", item)

	})
}
