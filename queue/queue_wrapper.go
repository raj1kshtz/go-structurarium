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
	return q.queue.addFirst(value)
}

// AddLast adds element at last of the queues
func (q *WrapperQueue[T]) AddLast(value T) error {
	return q.queue.addLast(value)
}

// Offer adds element at last of the queue
func (q *WrapperQueue[T]) Offer(value T) error {
	return q.queue.offer(value)
}

// Poll removes element from start of the queue
func (q *WrapperQueue[T]) Poll() (T, error) {
	return q.queue.poll()
}

// PollFirst removes element from start of the queue
func (q *WrapperQueue[T]) PollFirst() (T, error) {
	return q.queue.pollFirst()
}

// PollLast removes element from last of the queue
func (q *WrapperQueue[T]) PollLast() (T, error) {
	return q.queue.pollLast()
}

// Display prints the element of the queue
func (q *WrapperQueue[T]) Display() {
	q.queue.display()
}

// Size returns size of the queue
func (q *WrapperQueue[T]) Size() int {
	return q.queue.size()
}

// Peek returns the first of the queue to check
func (q *WrapperQueue[T]) Peek() (T, error) {
	return q.queue.peek()
}

// Clear empties the queue and return error
func (q *WrapperQueue[T]) Clear() error {
	return q.queue.clear()
}

func (q *WrapperQueue[T]) IsEmpty() bool {
	return q.queue.isEmpty()
}
