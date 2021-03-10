package brackets_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilgustavo/godata/exercises/brackets"
)

func TestValidateBrackets(t *testing.T) {
	t.Run("Prueba", func(t *testing.T) {
		assert.NoError(t, brackets.Validate("Hola{[]} mundo!"))
		assert.Error(t, brackets.Validate("{})"))
	})
}
