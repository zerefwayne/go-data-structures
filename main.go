package main

import (
	"fmt"
	"github.com/zerefwayne/go-data-structures/stack"
	)

func main() {
	newStack := new(stack.Stack)

	_, _ = newStack.Push(34)

	_, _ = newStack.Pop()

	_, err := newStack.Pop()

	if err != nil {
		fmt.Println(err.Error())
	}

	_, _ = newStack.Push(24)
	_, _ = newStack.Push(45)

	for !newStack.Empty() {
		top, _ := newStack.Top()
		_, _ = newStack.Pop()
		fmt.Println(top)
	}


}
