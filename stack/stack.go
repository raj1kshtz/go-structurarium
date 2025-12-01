package stack

import (
	"fmt"

	"github.com/raj1kshtz/go-structurarium/vector"
)

type stackRequest[T any] struct {
	action    string
	value     T
	replyChan chan interface{}
}

type GenericStack[T any] struct {
	stackChan     chan stackRequest[T]
	vectorWrapper *vector.WrapperVector[T]
}

func NewGenericStack[T any](initialCapacity ...int) *GenericStack[T] {
	var capacity int
	if len(initialCapacity) > 0 {
		capacity = initialCapacity[0]
	} else {
		capacity = 0
	}
	s := &GenericStack[T]{
		stackChan:     make(chan stackRequest[T]),
		vectorWrapper: vector.NewWrapperVector[T](capacity),
	}
	go s.manageStack()
	return s
}

func (s *GenericStack[T]) manageStack() {
	for req := range s.stackChan {
		switch req.action {
		case "push":
			if err := s.push(req.value); err != nil {
				req.replyChan <- err
			} else {
				req.replyChan <- nil
			}
		case "pop":
			value, err := s.pop()
			req.replyChan <- value
			if err != nil {
				req.replyChan <- err
			}
		case "peek":
			value, err := s.peek()
			req.replyChan <- value
			if err != nil {
				req.replyChan <- err
			}
		case "clear":
			if err := s.clear(); err != nil {
				req.replyChan <- err
			} else {
				req.replyChan <- nil
			}
		case "isEmpty":
			req.replyChan <- s.isEmpty()
		case "size":
			req.replyChan <- s.size()
		}
	}
}

func (s *GenericStack[T]) push(value T) error {
	return s.vectorWrapper.Add(value)
}

func (s *GenericStack[T]) pop() (T, error) {
	if s.vectorWrapper.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack underflow")
	}
	value, err := s.vectorWrapper.Get(s.vectorWrapper.Size() - 1)
	if err != nil {
		var zero T
		return zero, err
	}
	_ = s.vectorWrapper.RemoveAt(s.vectorWrapper.Size() - 1)
	return value, nil
}

func (s *GenericStack[T]) peek() (T, error) {
	if s.vectorWrapper.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.vectorWrapper.Get(s.vectorWrapper.Size() - 1)
}

func (s *GenericStack[T]) clear() error {
	return s.vectorWrapper.Clear()
}

func (s *GenericStack[T]) isEmpty() bool {
	return s.vectorWrapper.IsEmpty()
}

func (s *GenericStack[T]) size() int {
	return s.vectorWrapper.Size()
}

func replyChanReceive(replyChan chan interface{}) error {
	result := <-replyChan
	if err, ok := result.(error); ok {
		return err
	}
	return nil
}
