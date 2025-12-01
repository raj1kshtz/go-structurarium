package stack

type WrapperStack[T any] struct {
	stack *GenericStack[T]
}

// NewWrapperStack creates a new instance of WrapperStack
func NewWrapperStack[T any](initialCapacity ...int) *WrapperStack[T] {
	return &WrapperStack[T]{stack: NewGenericStack[T](initialCapacity...)}
}

// Push adds a new element to the stack
func (sw *WrapperStack[T]) Push(value T) error {
	return sw.stack.push(value)
}

// Pop removes and returns the top element of the stack
func (sw *WrapperStack[T]) Pop() (T, error) {
	return sw.stack.pop()
}

// Peek returns the top element of the stack without removing it
func (sw *WrapperStack[T]) Peek() (T, error) {
	return sw.stack.peek()
}

// Clear empties the stack and return error
func (sw *WrapperStack[T]) Clear() error {
	return sw.stack.clear()
}

// Size returns size of stack
func (sw *WrapperStack[T]) Size() int {
	return sw.stack.size()
}

// IsEmpty returns true if stack is empty
func (sw *WrapperStack[T]) IsEmpty() bool {
	return sw.stack.isEmpty()
}
