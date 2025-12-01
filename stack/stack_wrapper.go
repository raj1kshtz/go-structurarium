package stack

type WrapperStack[T any] struct {
	stack *GenericStack[T]
}

func NewWrapperStack[T any](initialCapacity ...int) *WrapperStack[T] {
	return &WrapperStack[T]{stack: NewGenericStack[T](initialCapacity...)}
}

func (sw *WrapperStack[T]) Push(value T) error {
	return sw.stack.push(value)
}

func (sw *WrapperStack[T]) Pop() (T, error) {
	return sw.stack.pop()
}

func (sw *WrapperStack[T]) Peek() (T, error) {
	return sw.stack.peek()
}

func (sw *WrapperStack[T]) Clear() error {
	return sw.stack.clear()
}

func (sw *WrapperStack[T]) Size() int {
	return sw.stack.size()
}

func (sw *WrapperStack[T]) IsEmpty() bool {
	return sw.stack.isEmpty()
}
