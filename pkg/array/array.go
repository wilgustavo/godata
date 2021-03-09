package array

import "errors"

const defaultCapacity = 10

type arrayHead struct {
	len int
	cap int
	arr []int
}

// Array operations over an array
type Array interface {
	Length() int
	Add(value int)
	Get(position int) (int, error)
}

// NewArray create an array
func NewArray() Array {
	return &arrayHead{len: 0, cap: defaultCapacity, arr: make([]int, defaultCapacity)}
}

func (head *arrayHead) Length() int {
	return head.len
}

func (head *arrayHead) Add(value int) {
	head.arr[head.len] = value
	head.len++
}

func (head *arrayHead) Get(positon int) (int, error) {
	if head.len > positon {
		return 0, errors.New("array index out of bounds")
	}
	return head.arr[positon], nil
}
