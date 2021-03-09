package array_test

import (
	"testing"

	"github.com/wilgustavo/godata/pkg/array"
)

func TestNewArray(t *testing.T) {

	arr := array.NewArray()
	if arr.Length() != 0 {
		t.Error("El tamaño inicial debería dar cero")
	}

}

func TestAddValue(t *testing.T) {
	arr := array.NewArray()
	arr.Add(1787)
	if arr.Length() != 1 {
		t.Error("El tamaño al agregar un elemento deberia ser 1")
	}

	arr.Add(8001)
	if arr.Length() != 2 {
		t.Error("El tamaño al agregar un segundo elemento deberia ser 2")
	}

}
