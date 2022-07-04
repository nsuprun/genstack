package main

import (
	"fmt"

	stack "github.com/nsuprun/genstack"
)

func main() {

	s := "(t)es(t)()()"

	if hasBalancedBrackets(s) == true {
		fmt.Printf("Brackets in %s are balanced\r\n", s)
	} else {
		fmt.Printf("Brackets in %s are not balanced\r\n", s)
	}
}

func hasBalancedBrackets(s string) bool {
	stack := stack.New[string]()

	for _, char := range s {
		symbol := fmt.Sprintf("%c", char)
		if symbol == "(" {
			stack.Push(symbol)
		} else if symbol == ")" {
			if stack.IsEmpty() {
				return false
			}

			stack.Pop()
		}
	}
	return stack.IsEmpty()
}
