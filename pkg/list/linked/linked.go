package linked

import (
	"errors"
)

// Item is element of list
type Item struct{}

type header struct {
	lenght uint
	first  *node
	last   *node
}

type node struct {
	value int
	next  *node
}

// List is a Linked list opetations
type List interface {
	AddAt(pos uint, value int) (uint, error)
	AddFirst(value int) uint
	AddLast(value int) uint
	GetAt(pos uint) (int, error)
	GetFirst() (int, error)
	Lenght() uint
	RemoveFirst() (uint, error)
	RemoveLast() (uint, error)
	Revert()
	ToSlice() []int
}

// NewLinkedList create a linked list
func NewLinkedList() List {
	return &header{}
}

func (list *header) Lenght() uint {
	return list.lenght
}

func (list *header) AddAt(pos uint, value int) (uint, error) {
	if pos > list.lenght {
		return list.lenght, errors.New("index out of bound")
	}

	if pos == 0 {
		return list.AddFirst(value), nil
	}

	if pos == list.lenght {
		return list.AddLast(value), nil
	}

	var prev, item *node
	trav := list.first
	for i := uint(0); i < pos; i++ {
		prev = trav
		trav = trav.next
	}
	item.next = trav
	item.value = value
	prev.next = item
	list.lenght++
	return list.lenght, nil
}

func (list *header) AddFirst(value int) uint {
	list.lenght++
	newNode := &node{value: value}
	if list.last == nil {
		list.last = newNode
	}
	if list.first != nil {
		newNode.next = list.first
	}
	list.first = newNode
	return list.lenght
}

func (list *header) AddLast(value int) uint {
	list.lenght++
	newNode := &node{value: value}
	if list.first == nil {
		list.first = newNode
	}
	if list.last != nil {
		list.last.next = newNode
	}
	list.last = newNode
	return list.lenght
}

func (list *header) GetFirst() (int, error) {
	if list.lenght == 0 {
		return 0, errors.New("list is empty")
	}
	return list.first.value, nil
}

func (list *header) GetAt(pos uint) (int, error) {
	if pos >= uint(list.lenght) {
		return 0, errors.New("index out of bounds")
	}
	trav := list.getNodeAt(pos)
	return trav.value, nil
}

func (list *header) ToSlice() []int {
	result := []int{}
	for node := list.first; node != nil; node = node.next {
		result = append(result, node.value)
	}
	return result
}

func (list *header) RemoveFirst() (uint, error) {
	if list.lenght == 0 {
		return 0, errors.New("List is empty")
	}

	temp := list.first
	list.first = temp.next
	temp.next = nil

	list.lenght--
	return list.lenght, nil
}

func (list *header) RemoveLast() (uint, error) {
	if list.lenght == 0 {
		return 0, errors.New("List is empty")
	}

	temp := list.first
	var trav *node = nil
	for temp.next != nil {
		trav = temp
		temp = temp.next
	}
	if trav != nil {
		trav.next = nil
	} else {
		list.first = trav
	}
	list.last = trav

	list.lenght--
	return list.lenght, nil
}

func (list *header) Revert() {
	trav := list.first
	if trav == nil {
		return
	}

	temp := trav.next
	for temp != nil {
		aux := temp.next
		temp.next = trav
		trav = temp
		temp = aux
	}
	list.first.next = nil
	list.last = list.first
	list.first = trav
}

func (list *header) getNodeAt(pos uint) *node {
	trav := list.first
	for i := uint(0); i < pos; i++ {
		trav = trav.next
	}
	return trav
}
