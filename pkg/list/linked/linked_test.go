package linked_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilgustavo/godata/pkg/list/linked"
)

func TestNewLinkedList(t *testing.T) {
	t.Run("When create a linked list should be empty", func(t *testing.T) {
		list := linked.NewLinkedList()
		assert.Equal(t, uint(0), list.Lenght(), "An new List must have a zero length")
	})
}

func TestAddFirst(t *testing.T) {
	t.Run("When add a value to empty list should be a length of 1", func(t *testing.T) {
		list := linked.NewLinkedList()
		len := list.AddFirst(7)
		assert.Equal(t, uint(1), len)
		assert.Equal(t, uint(1), list.Lenght())
	})

	t.Run("When append two values to empty list should be a length of 2", func(t *testing.T) {
		list := linked.NewLinkedList()
		len := list.AddFirst(7)
		len = list.AddFirst(5)
		assert.Equal(t, uint(2), len)
		assert.Equal(t, uint(2), list.Lenght())

		value, _ := list.GetAt(0)
		assert.Equal(t, 5, value)
		value, _ = list.GetAt(1)
		assert.Equal(t, 7, value)
	})

}

func TestGetFirst(t *testing.T) {
	t.Run("Should give a firt elemet", func(t *testing.T) {
		list := linked.NewLinkedList()
		list.AddFirst("Mundo")
		list.RemoveFirst()
		list.AddLast("Hola")
		list.AddLast("Mundo")
		item, err := list.GetFirst()
		assert.NoError(t, err)
		assert.Equal(t, "Hola", item)
	})
}

func TestAddLast(t *testing.T) {
	t.Run("When append a value to empty list should be a length of 1", func(t *testing.T) {
		list := linked.NewLinkedList()
		len := list.AddLast(5)
		assert.Equal(t, uint(1), len, "Append must return 1")
		assert.Equal(t, uint(1), list.Lenght(), "An new List must have a length of 1")
	})

	t.Run("When append two values to empty list should be a length of 2", func(t *testing.T) {
		list := linked.NewLinkedList()
		len := list.AddLast(5)
		len = list.AddLast(8)
		assert.Equal(t, uint(2), len, "Append must return 2")
		assert.Equal(t, uint(2), list.Lenght(), "The List must have a length of 2")
	})
}

func TestGetAt(t *testing.T) {
	t.Run("Should get the value at the given position", func(t *testing.T) {
		list := linked.NewLinkedList()
		list.AddLast(3)
		list.AddLast(5)
		result, err := list.GetAt(0)

		assert.Equal(t, 3, result)
		assert.Nil(t, err)

		result, err = list.GetAt(1)

		assert.Equal(t, 5, result)
		assert.Nil(t, err)

		list.AddLast(8)
		result, err = list.GetAt(2)
		assert.Equal(t, 8, result)
		assert.Nil(t, err)

		result, err = list.GetAt(0)
		assert.Equal(t, 3, result)
		assert.Nil(t, err)

	})

	t.Run("Should give an error if the position is greater than the length", func(t *testing.T) {
		list := linked.NewLinkedList()
		_, err := list.GetAt(0)
		assert.NotNil(t, err)

		list.AddFirst(1)
		_, err = list.GetAt(1)
		assert.NotNil(t, err)

		_, err = list.GetAt(0)
		assert.Nil(t, err)

		_, err = list.GetAt(99999999999)
		assert.NotNil(t, err)
	})
}

func TestRemoveFirst(t *testing.T) {
	t.Run("Should get error when remove on an empty list", func(t *testing.T) {
		list := linked.NewLinkedList()
		_, err := list.RemoveFirst()
		assert.NotNil(t, err)
	})
	t.Run("Should remove the first element", func(t *testing.T) {
		list := linked.NewLinkedList()
		len := list.AddFirst(17)
		len, err := list.RemoveFirst()
		assert.Equal(t, uint(0), len)
		assert.Nil(t, err)
		assert.Equal(t, uint(0), list.Lenght())
		_, err = list.GetAt(0)
		assert.NotNil(t, err)

		list.AddLast(23)
		list.AddLast(29)
		len, err = list.RemoveFirst()
		assert.Nil(t, err)
		assert.Equal(t, uint(1), len)
		assert.Equal(t, uint(1), list.Lenght())
		value, err := list.GetAt(0)
		assert.Nil(t, err)
		assert.Equal(t, 29, value)
	})
}

func TestToSlice(t *testing.T) {
	t.Run("Deberia retornar los elementos de la lista", func(t *testing.T) {
		list := linked.NewLinkedList()
		list.AddLast(1)

		assert.Equal(t, []interface{}{1}, list.ToSlice())

		list.AddLast(5)
		list.AddLast(13)
		list.AddFirst(8)
		assert.Equal(t, []interface{}{8, 1, 5, 13}, list.ToSlice())
	})

	t.Run("Deberia retornar vacio si no hay elementos en la lista", func(t *testing.T) {
		list := linked.NewLinkedList()
		assert.Equal(t, []interface{}{}, list.ToSlice())
	})
}

func TestRemoveLast(t *testing.T) {
	t.Run("Deberia borrar el ultimo elemento de la lista", func(t *testing.T) {
		list := linked.NewLinkedList()
		list.AddLast(1)
		list.AddLast(5)

		len, err := list.RemoveLast()
		assert.Equal(t, uint(1), len)
		assert.NoError(t, err)
		assert.Equal(t, []interface{}{1}, list.ToSlice())

		len, err = list.RemoveLast()
		assert.Equal(t, uint(0), len)
		assert.NoError(t, err)
		assert.Equal(t, []interface{}{}, list.ToSlice())
	})
}

func TestRevert(t *testing.T) {
	t.Run("Deberia revertir la lista", func(t *testing.T) {
		list := linked.NewLinkedList()
		list.AddLast(5)
		list.Revert()
		assert.Equal(t, []interface{}{5}, list.ToSlice())

		list.AddLast(77)
		list.AddLast(999)
		assert.Equal(t, []interface{}{5, 77, 999}, list.ToSlice())
		list.Revert()
		assert.Equal(t, []interface{}{999, 77, 5}, list.ToSlice())
		list.Revert()
		assert.Equal(t, []interface{}{5, 77, 999}, list.ToSlice())
	})
	t.Run("No deberia revertir una lista vacia", func(t *testing.T) {
		list := linked.NewLinkedList()
		list.Revert()
		assert.Equal(t, []interface{}{}, list.ToSlice())
	})
}
