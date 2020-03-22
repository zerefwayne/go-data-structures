package stack

import (
	"errors"
	"fmt"
)

var (
	// ErrStackEmpty : When the stack is empty
	ErrStackEmpty = errors.New("error: Stack is empty")
)

// Generic : A generic data type to hold any type of data
type Generic interface{}

// Stack holds the items in a LIFO Structure
type Stack struct {
	items []Generic // items of generic type in the stack
	size  int       // size of stack
}

// Len : Returns the size of Stack
func (stack *Stack) Len() int {
	return stack.size
}

// Print : Prints the stack in oldest -> latest order
func (stack *Stack) Print() {
	for _, val := range stack.items {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

// Push : Pushes an item onto Stack, returns the updated size of array
func (stack *Stack) Push(item Generic) (int, error) {
	stack.items = append(stack.items, item)
	stack.size = stack.size + 1
	return stack.size, nil
}

// Top : Returns the top element of the stack
func (stack *Stack) Top() (Generic, error) {
	if stack.size > 0 {
		return stack.items[stack.size-1], nil
	}

	return nil, ErrStackEmpty
}

// Pop : Pops off the top element and returns the updated size
func (stack *Stack) Pop() (int, error) {
	if stack.size > 0 {
		stack.size = stack.size - 1
		stack.items = stack.items[:stack.size]
		return stack.size, nil
	}

	return 0, ErrStackEmpty
}

// Empty : Returns true if the stack is empty
func (stack *Stack) Empty() bool {
	return stack.size == 0
}