package main

import (
	"fmt"

	stack "github.com/nsuprun/genstack"
)

func main() {

	s := "(), (()), ()(), (((())))"

	if hasBalancedBrackets(s) == true {
		fmt.Printf("String %s has balanced brackets\r\n", s)
	} else {
		fmt.Printf("String %s is not balanced\r\n", s)
	}
}

func hasBalancedBrackets(s string) bool {
	stack := stack.New[string]()

	for _, char := range s {
		symbol := fmt.Sprintf("%c", char)
		if symbol == "(" {
			stack.Push(symbol)
		} else if symbol == ")" {
			stack.Pop()
		}
	}

	return stack.IsEmpty()
}
