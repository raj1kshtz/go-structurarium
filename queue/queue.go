package queue

import "fmt"

type queueRequest[T any] struct {
	action    string
	value     T
	replyChan chan interface{}
}

type GenericQueue[T any] struct {
	queueChan chan queueRequest[T]
}

func NewGenericQueue[T any]() *GenericQueue[T] {
	s := &GenericQueue[T]{queueChan: make(chan queueRequest[T])}
	go s.manageQueue()
	return s
}

func (q *GenericQueue[T]) manageQueue() {
	data := make([]T, 0)

	for req := range q.queueChan {
		switch req.action {
		case "addFirst":
			data = append([]T{req.value}, data...)
			req.replyChan <- nil
		case "addLast", "offer":
			data = append(data, req.value)
			req.replyChan <- nil
		case "peek":
			if len(data) == 0 {
				req.replyChan <- fmt.Errorf("queue is empty")
			} else {
				req.replyChan <- data[0]
			}
		case "pollFirst", "poll":
			if len(data) == 0 {
				req.replyChan <- fmt.Errorf("queue is empty")
			} else {
				value := data[0]
				data = data[1:]
				req.replyChan <- value
			}
		case "pollLast":
			if len(data) == 0 {
				req.replyChan <- fmt.Errorf("queue is empty")
			} else {
				value := data[len(data)-1]
				data = data[:len(data)-1]
				req.replyChan <- value
			}
		case "display":
			req.replyChan <- fmt.Sprintf("%v", data)
		case "size":
			req.replyChan <- len(data)
		case "clear":
			data = make([]T, 0)
			req.replyChan <- nil
		case "isEmpty":
			req.replyChan <- len(data) == 0
		}
	}
}

func (q *GenericQueue[T]) addFirst(value T) error {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "addFirst", value: value, replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (q *GenericQueue[T]) addLast(value T) error {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "addLast", value: value, replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (q *GenericQueue[T]) offer(value T) error {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "offer", value: value, replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (q *GenericQueue[T]) peek() (T, error) {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "peek", replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

func (q *GenericQueue[T]) poll() (T, error) {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "pollFirst", replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

func (q *GenericQueue[T]) pollFirst() (T, error) {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "pollFirst", replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

func (q *GenericQueue[T]) pollLast() (T, error) {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "pollLast", replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

func (q *GenericQueue[T]) display() string {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "display", replyChan: replyChan}
	return (<-replyChan).(string)
}

func (q *GenericQueue[T]) size() int {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "size", replyChan: replyChan}
	return (<-replyChan).(int)
}

func replyChanReceive(replyChan chan interface{}) error {
	result := <-replyChan
	if err, ok := result.(error); ok {
		return err
	}
	return nil
}

func (q *GenericQueue[T]) clear() error {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "clear", replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (q *GenericQueue[T]) isEmpty() bool {
	replyChan := make(chan interface{})
	q.queueChan <- queueRequest[T]{action: "isEmpty", replyChan: replyChan}
	return (<-replyChan).(bool)
}

// Wrapper method
func (q *GenericQueue[T]) AddFirst(value T) error {
	return q.addFirst(value)
}

func (q *GenericQueue[T]) AddLast(value T) error {
	return q.addLast(value)
}

func (q *GenericQueue[T]) Offer(value T) error {
	return q.offer(value)
}

func (q *GenericQueue[T]) PollFirst() (T, error) {
	return q.pollFirst()
}

func (q *GenericQueue[T]) PollLast() (T, error) {
	return q.pollLast()
}

func (q *GenericQueue[T]) Poll() (T, error) {
	return q.poll()
}

func (q *GenericQueue[T]) Display() {
	fmt.Println(q.display())
}

func (q *GenericQueue[T]) Size() int {
	return q.size()
}

func (q *GenericQueue[T]) Peek() (T, error) {
	return q.peek()
}

func (q *GenericQueue[T]) Clear() error {
	return q.clear()
}

func (q *GenericQueue[T]) IsEmpty() bool {
	return q.isEmpty()
}
