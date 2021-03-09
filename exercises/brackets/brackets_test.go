package brackets_test

import (
	"testing"

	"github.com/wilgustavo/godata/exercises/brackets"
)

func TestValidateBrackets(t *testing.T) {
	t.Run("Prueba", func(t *testing.T) {
		brackets.Validate("Hola{[]} mundo!")

	})
}
