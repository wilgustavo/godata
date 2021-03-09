package brackets

import (
	"fmt"
)

// Validate return error if brackets is not valid
func Validate(value string) error {

	for _, r := range value {

		char := string(r)
		fmt.Println(char)
		switch char {
		case "[", "{", "(":
			fmt.Println("POP!")
		case "]", "}", ")":
			fmt.Println("PUSH!")
		}

	}

	return nil

}
