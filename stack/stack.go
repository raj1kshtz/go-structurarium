package stack

import "fmt"

type stackRequest[T any] struct {
	action    string
	value     T
	replyChan chan interface{}
}

type GenericStack[T any] struct {
	stackChan chan stackRequest[T]
}

func NewGenericStack[T any]() *GenericStack[T] {
	s := &GenericStack[T]{stackChan: make(chan stackRequest[T])}
	go s.manageStack()
	return s
}

func (s *GenericStack[T]) manageStack() {
	data := make([]T, 0)

	for req := range s.stackChan {
		switch req.action {
		case "push":
			data = append(data, req.value)
			req.replyChan <- nil
		case "pop":
			if len(data) == 0 {
				req.replyChan <- fmt.Errorf("stack underflow")
			} else {
				value := data[len(data)-1]
				data = data[:len(data)-1]
				req.replyChan <- value
			}
		case "top":
			if len(data) == 0 {
				req.replyChan <- fmt.Errorf("stack is empty")
			} else {
				req.replyChan <- data[len(data)-1]
			}
		case "display":
			req.replyChan <- fmt.Sprintf("%v", data)
		case "clear":
			data = make([]T, 0)
			req.replyChan <- nil
		}
	}
}

func (s *GenericStack[T]) push(value T) error {
	replyChan := make(chan interface{})
	s.stackChan <- stackRequest[T]{action: "push", value: value, replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (s *GenericStack[T]) pop() (T, error) {
	replyChan := make(chan interface{})
	s.stackChan <- stackRequest[T]{action: "pop", replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

func (s *GenericStack[T]) top() (T, error) {
	replyChan := make(chan interface{})
	s.stackChan <- stackRequest[T]{action: "top", replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

func (s *GenericStack[T]) display() string {
	replyChan := make(chan interface{})
	s.stackChan <- stackRequest[T]{action: "display", replyChan: replyChan}
	return (<-replyChan).(string)
}

func (s *GenericStack[T]) clear() error {
	replyChan := make(chan interface{})
	s.stackChan <- stackRequest[T]{action: "clear", replyChan: replyChan}
	return replyChanReceive(replyChan)
}

// Wrapper methods

func (s *GenericStack[T]) Push(value T) error {
	return s.push(value)
}

func (s *GenericStack[T]) Pop() (T, error) {
	return s.pop()
}

func (s *GenericStack[T]) Top() (T, error) {
	return s.top()
}

func (s *GenericStack[T]) Display() {
	fmt.Println(s.display())
}

func (s *GenericStack[T]) Clear() error {
	return s.clear()
}

// Helper function to receive from the reply channel and handle errors
func replyChanReceive(replyChan chan interface{}) error {
	result := <-replyChan
	if err, ok := result.(error); ok {
		return err
	}
	return nil
}
