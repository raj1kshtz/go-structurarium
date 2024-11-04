package queue

type GenericQueueWrapper[T comparable] struct {
	queue *GenericQueue[T]
}

func NewGenericQueueWrapper[T comparable](initialCapacity ...int) *GenericQueueWrapper[T] {
	return &GenericQueueWrapper[T]{queue: NewGenericQueue[T](initialCapacity...)}
}

func (w *GenericQueueWrapper[T]) Enqueue(value T) bool {
	replyChan := make(chan interface{})
	w.queue.queueChan <- queueRequest[T]{action: "enQueue", value: value, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericQueueWrapper[T]) Dequeue() (T, bool) {
	replyChan := make(chan interface{})
	w.queue.queueChan <- queueRequest[T]{action: "deQueue", replyChan: replyChan}
	value := (<-replyChan).(T)
	success := (<-replyChan).(bool)
	return value, success
}

func (w *GenericQueueWrapper[T]) Size() int {
	replyChan := make(chan interface{})
	w.queue.queueChan <- queueRequest[T]{action: "size", replyChan: replyChan}
	return (<-replyChan).(int)
}

func (w *GenericQueueWrapper[T]) IsEmpty() bool {
	replyChan := make(chan interface{})
	w.queue.queueChan <- queueRequest[T]{action: "isEmpty", replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericQueueWrapper[T]) Clear() {
	replyChan := make(chan interface{})
	w.queue.queueChan <- queueRequest[T]{action: "clear", replyChan: replyChan}
}

func (w *GenericQueueWrapper[T]) Peek() int {
	replyChan := make(chan interface{})
	w.queue.queueChan <- queueRequest[T]{action: "peek", replyChan: replyChan}
	return (<-replyChan).(int)
}

func (w *GenericQueueWrapper[T]) ToArray() []T {
	replyChan := make(chan interface{})
	w.queue.queueChan <- queueRequest[T]{action: "toArray", replyChan: replyChan}
	return (<-replyChan).([]T)
}
