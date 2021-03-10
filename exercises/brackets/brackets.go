package brackets

import (
	"fmt"

	"github.com/wilgustavo/godata/pkg/list/stack"
)

// https://youtu.be/RBSGKlAvoiM?t=4075

// Validate return error if brackets is not valid
func Validate(value string) error {

	pila := stack.NewStack()

	for _, r := range value {

		char := string(r)
		switch char {
		case "[", "{", "(":
			pila.Push(char)
		case "]", "}", ")":
			token, err := pila.PopString()
			if err != nil ||
				(token == "[" && char != "]") ||
				(token == "{" && char != "}") ||
				(token == "(" && char != ")") {
				return fmt.Errorf("No es valido")
			}
		}

	}

	if !pila.IsEmpty() {
		return fmt.Errorf("No es valido")
	}
	return nil

}
