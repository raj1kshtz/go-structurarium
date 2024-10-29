package stack

type WrapperStack[T any] struct {
	stack *GenericStack[T]
}

// NewWrapperStack creates a new instance of WrapperStack
func NewWrapperStack[T any]() *WrapperStack[T] {
	return &WrapperStack[T]{stack: NewGenericStack[T]()}
}

// Push adds a new element to the stack
func (sw *WrapperStack[T]) Push(value T) error {
	return sw.stack.Push(value)
}

// Pop removes and returns the top element of the stack
func (sw *WrapperStack[T]) Pop() (T, error) {
	return sw.stack.Pop()
}

// Top returns the top element of the stack without removing it
func (sw *WrapperStack[T]) Top() (T, error) {
	return sw.stack.Top()
}

// Display shows all elements in the stack
func (sw *WrapperStack[T]) Display() {
	sw.stack.Display()
}

// Clear empties the stack and return error
func (sw *WrapperStack[T]) Clear() error {
	return sw.stack.Clear()
}

func (sw *WrapperStack[T]) IsEmpty() bool {
	return sw.stack.IsEmpty()
}
