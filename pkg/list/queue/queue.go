package queue

import (
	"fmt"

	"github.com/wilgustavo/godata/pkg/list/linked"
)

type queue struct {
	list linked.List
}

// Queue represents operations on a Queue
type Queue interface {
	Enqueue(item interface{})
	Dequeue() (interface{}, error)
	DequeueString() (string, error)
	IsEmpty() bool
}

// NewQueue create a new queue
func NewQueue() Queue {
	return &queue{list: linked.NewLinkedList()}
}

func (q *queue) Enqueue(item interface{}) {
	q.list.AddLast(item)
}

func (q *queue) Dequeue() (interface{}, error) {
	item, err := q.list.GetFirst()
	if err != nil {
		return nil, err
	}
	_, err = q.list.RemoveFirst()
	return item, err
}

func (q *queue) DequeueString() (string, error) {
	item, err := q.Dequeue()
	if err != nil {
		return "", err
	}
	str, ok := item.(string)
	if !ok {
		return "", fmt.Errorf("Item not a string")
	}
	return str, nil
}

func (q *queue) IsEmpty() bool {
	return q.list.Lenght() == 0
}
