package stack

import (
	"errors"
)

// Stack definition
type Stack []interface{}

// Len stack'length
func (stack Stack) Len() int {
	return len(stack)
}

// IsEmpty tell the stack is empty
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

// Cap stack'cap
func (stack Stack) Cap() int {
	return cap(stack)
}

// Pop stack pop
func (stack *Stack) Pop() interface{} {
	newStack := *stack
	if len(newStack) == 0 {
		return nil
	}
	value := newStack[len(newStack)-1]
	*stack = newStack[:len(newStack)-1]
	return value
}

// Push stack push
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

// Top stack'top
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}
