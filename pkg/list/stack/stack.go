package stack

import (
	"github.com/wilgustavo/godata/pkg/list/linked"
)

type stack struct {
	list linked.List
}

// Stack represents a stack
type Stack interface {
	Push(value int)
	Pop() (int, error)
	IsEmpty() bool
}

// NewStack returns a new stack
func NewStack() Stack {
	return &stack{list: linked.NewLinkedList()}
}

func (s *stack) Push(value int) {
	s.list.AddFirst(value)
}

func (s *stack) Pop() (int, error) {
	value, err := s.list.GetFirst()
	if err != nil {
		return 0, err
	}
	_, err = s.list.RemoveFirst()
	if err != nil {
		return 0, err
	}
	return value, err
}

func (s *stack) IsEmpty() bool {
	return s.list.Lenght() == 0
}
