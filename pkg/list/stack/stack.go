package stack

import (
	"fmt"

	"github.com/wilgustavo/godata/pkg/list/linked"
)

type stack struct {
	list linked.List
}

// Stack represents a stack
type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	PopString() (string, error)
	IsEmpty() bool
	ToSlice() []interface{}
}

// NewStack returns a new stack
func NewStack() Stack {
	return &stack{list: linked.NewLinkedList()}
}

func (s *stack) Push(value interface{}) {
	s.list.AddFirst(value)
}

func (s *stack) Pop() (interface{}, error) {
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

func (s *stack) PopString() (string, error) {
	str, err := s.Pop()
	if err != nil {
		return "", err
	}
	token, ok := str.(string)
	if !ok {
		return "", fmt.Errorf("Elemento no es un string")
	}
	return token, nil
}

func (s *stack) ToSlice() []interface{} {
	return s.list.ToSlice()
}
