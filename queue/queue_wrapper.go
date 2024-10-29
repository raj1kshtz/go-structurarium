package queue

type WrapperQueue[T any] struct {
	queue *GenericQueue[T]
}

// NewWrapperQueue creates a new instance of WrapperStack
func NewWrapperQueue[T any]() *WrapperQueue[T] {
	return &WrapperQueue[T]{
		queue: NewGenericQueue[T](),
	}
}

// AddFirst adds element at first of the queue
func (q *WrapperQueue[T]) AddFirst(value T) error {
	return q.queue.AddFirst(value)
}

// AddLast adds element at last of the queues
func (q *WrapperQueue[T]) AddLast(value T) error {
	return q.queue.AddLast(value)
}

// Offer adds element at last of the queue
func (q *WrapperQueue[T]) Offer(value T) error {
	return q.queue.Offer(value)
}

// Poll removes element from start of the queue
func (q *WrapperQueue[T]) Poll() (T, error) {
	return q.queue.Poll()
}

// PollFirst removes element from start of the queue
func (q *WrapperQueue[T]) PollFirst() (T, error) {
	return q.queue.PollFirst()
}

// PollLast removes element from last of the queue
func (q *WrapperQueue[T]) PollLast() (T, error) {
	return q.queue.PollLast()
}

// Display prints the element of the queue
func (q *WrapperQueue[T]) Display() {
	q.queue.Display()
}

// Size returns size of the queue
func (q *WrapperQueue[T]) Size() int {
	return q.queue.Size()
}

// Peek returns the first of the queue to check
func (q *WrapperQueue[T]) Peek() (T, error) {
	return q.queue.Peek()
}

// Clear empties the queue and return error
func (q *WrapperQueue[T]) Clear() error {
	return q.queue.Clear()
}

func (q *WrapperQueue[T]) IsEmpty() bool {
	return q.queue.IsEmpty()
}
