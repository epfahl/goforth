package main

import "fmt"

// Struct that holds stack values in a slice.
type Stack struct {
	data []int64
}

// Return an empty `Stack`.
func NewStack() Stack {
	return Stack{}
}

// Return the size of the stack.
func (stack *Stack) Size() int {
	return len(stack.data)
}

// Push a value onto the stack.
func (stack *Stack) Push(value int64) bool {
	stack.data = append(stack.data, value)
	return true
}

// Add two integer values at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) Add() bool {
	return stack.binaryOp(func(a, b int64) int64 { return a + b })
}

// Subtract two integer values at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) Subtract() bool {
	return stack.binaryOp(func(a, b int64) int64 { return a - b })
}

// Multiply two integer values at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) Multiply() bool {
	return stack.binaryOp(func(a, b int64) int64 { return a * b })
}

// Divide two integer values at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) Divide() bool {
	size := stack.Size()
	if size < 2 {
		return true
	}

	if stack.data[size-1] == 0 {
		return true
	}

	stack.data[size-2] = stack.data[size-2] / stack.data[size-1]
	stack.data = stack.data[:size-1]
	return false
}

// Compare two integer values using `>` at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) GreaterThan() bool {
	return stack.binaryOp(func(a, b int64) int64 {
		if a > b {
			return -1
		} else {
			return 0
		}
	})
}

// Compare two integer values using `<` at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) LessThan() bool {
	return stack.binaryOp(func(a, b int64) int64 {
		if a < b {
			return -1
		} else {
			return 0
		}
	})
}

// Compare two integer values using `==` at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) EqualTo() bool {
	return stack.binaryOp(func(a, b int64) int64 {
		if a == b {
			return -1
		} else {
			return 0
		}
	})
}

// Apply conjunction to the two values at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) And() bool {
	return stack.binaryOp(func(a, b int64) int64 {
		return a & b
	})
}

// Apply disjunction two integer values at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) Or() bool {
	return stack.binaryOp(func(a, b int64) int64 {
		return a | b
	})
}

// Pop and print the value at the head of the stack.
//
// This returns `true` as an error value if the operation fails.
func (stack *Stack) Dot() bool {
	size := stack.Size()
	if size < 1 {
		return true
	}
	value := stack.data[size-1]
	fmt.Println(value)
	stack.data = stack.data[:size-1]
	return false
}

// Duplicate the value at the head of the stack.
//
// This returns `true` as an error if the stack is empty.
func (stack *Stack) Duplicate() bool {
	size := stack.Size()
	if size < 1 {
		return true
	}
	value := stack.data[size-1]
	stack.data = append(stack.data, value)
	return false
}

// Swap the two values at the head of the stack.
//
// This return `true` as an error value if there are less than two value at the
// head
// of the stack.
func (stack *Stack) Swap() bool {
	size := stack.Size()
	if size < 2 {
		return true
	}
	stack.data[size-2], stack.data[size-1] = stack.data[size-1], stack.data[size-2]
	return false
}

// Apply an integer-valued binary operation `op` on the top two
// elements of the stack.
//
// This returns `true` as an error value if the stack has fewer than 2 values
// or if the operation fails.
func (stack *Stack) binaryOp(op func(int64, int64) int64) bool {
	size := stack.Size()
	if size < 2 {
		return true
	}

	stack.data[size-2] = op(stack.data[size-2], stack.data[size-1])
	stack.data = stack.data[:size-1]
	return false
}
