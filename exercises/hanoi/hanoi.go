package hanoi

import (
	"github.com/wilgustavo/godata/pkg/list/stack"
)

// MoveTo prints the move from one pile to another
func MoveTo(from, to stack.Stack) {
	item, err := from.Pop()
	if err == nil {
		to.Push(item)
	}
}

// Hanoi print the moves in a Hanoi tower
func Hanoi(n uint, from, aux, to stack.Stack) {
	if n != 0 {
		Hanoi(n-1, from, to, aux)
		MoveTo(from, to)
		Hanoi(n-1, aux, from, to)
	}
}
