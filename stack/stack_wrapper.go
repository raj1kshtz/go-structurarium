package stack

type Wrapper[T any] struct {
	stack *GenericStack[T]
}

// NewStackWrapper creates a new instance of StackWrapper
func NewStackWrapper[T any]() *Wrapper[T] {
	return &Wrapper[T]{stack: NewGenericStack[T]()}
}

// Push adds a new element to the stack
func (sw *Wrapper[T]) Push(value T) error {
	return sw.stack.Push(value)
}

// Pop removes and returns the top element of the stack
func (sw *Wrapper[T]) Pop() (T, error) {
	return sw.stack.Pop()
}

// Top returns the top element of the stack without removing it
func (sw *Wrapper[T]) Top() (T, error) {
	return sw.stack.Top()
}

// Display shows all elements in the stack
func (sw *Wrapper[T]) Display() {
	sw.stack.Display()
}

// Clear empties the stack and return error
func (sw *Wrapper[T]) Clear() error {
	return sw.stack.Clear()
}
