package vector

import "fmt"

type vectorRequest[T any] struct {
	action    string
	index     int
	value     T
	replyChan chan interface{}
}

type GenericVector[T any] struct {
	vectorChan chan vectorRequest[T]
	data       []T
	capacity   int
}

func NewGenericVector[T any](initialCapacity ...int) *GenericVector[T] {
	var capacity int
	if len(initialCapacity) > 0 {
		capacity = initialCapacity[0]
	} else {
		capacity = 0
	}
	v := &GenericVector[T]{
		vectorChan: make(chan vectorRequest[T]),
		data:       make([]T, 0, capacity),
		capacity:   0,
	}
	go v.manageVector()
	return v
}

func (v *GenericVector[T]) manageVector() {
	for req := range v.vectorChan {
		switch req.action {
		case "add":
			v.data = append(v.data, req.value)
			v.capacity++
			req.replyChan <- nil
		case "addAt":
			if req.index < 0 || req.index > v.capacity {
				req.replyChan <- fmt.Errorf("index out of bounds")
			} else {
				v.data = append(v.data[:req.index], append([]T{req.value}, v.data[req.index:]...)...)
				v.capacity++
				req.replyChan <- nil
			}
		case "removeAt":
			if req.index < 0 || req.index >= v.capacity {
				req.replyChan <- fmt.Errorf("index out of bounds")
			} else {
				v.data = append(v.data[:req.index], v.data[req.index+1:]...)
				v.capacity--
				req.replyChan <- nil
			}
		case "get":
			if req.index < 0 || req.index >= v.capacity {
				req.replyChan <- fmt.Errorf("index out of bounds")
			} else {
				req.replyChan <- v.data[req.index]
			}
		case "set":
			if req.index < 0 || req.index >= v.capacity {
				req.replyChan <- fmt.Errorf("index out of bounds")
			} else {
				v.data[req.index] = req.value
				req.replyChan <- nil
			}
		case "capacity":
			req.replyChan <- v.capacity
		case "clear":
			v.data = make([]T, 0, cap(v.data)) // Reset data while keeping capacity
			v.capacity = 0
			req.replyChan <- nil
		case "isEmpty":
			req.replyChan <- v.capacity == 0
		case "ensureCapacity":
			req.replyChan <- nil
		case "trimToSize":
			v.data = v.data[:v.capacity]
			req.replyChan <- nil
		case "toArray":
			req.replyChan <- append([]T{}, v.data...)
		}
	}
}

func (v *GenericVector[T]) add(value T) error {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "add", value: value, replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (v *GenericVector[T]) addAt(index int, value T) error {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "addAt", index: index, value: value, replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (v *GenericVector[T]) removeAt(index int) error {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "removeAt", index: index, replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (v *GenericVector[T]) get(index int) (T, error) {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "get", index: index, replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

func (v *GenericVector[T]) set(index int, value T) error {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "set", index: index, value: value, replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (v *GenericVector[T]) size() int {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "capacity", replyChan: replyChan}
	return (<-replyChan).(int)
}

func (v *GenericVector[T]) clear() error {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "clear", replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (v *GenericVector[T]) isEmpty() bool {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "isEmpty", replyChan: replyChan}
	return (<-replyChan).(bool)
}

func replyChanReceive(replyChan chan interface{}) error {
	result := <-replyChan
	if err, ok := result.(error); ok {
		return err
	}
	return nil
}

func (v *GenericVector[T]) ensureCapacity(minCapacity int) {
	if minCapacity > cap(v.data) {
		replyChan := make(chan interface{})
		v.vectorChan <- vectorRequest[T]{action: "ensureCapacity", replyChan: replyChan}
		<-replyChan
	}
}

func (v *GenericVector[T]) trimToSize() error {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "trimToSize", replyChan: replyChan}
	return replyChanReceive(replyChan)
}

func (v *GenericVector[T]) toArray() []T {
	replyChan := make(chan interface{})
	v.vectorChan <- vectorRequest[T]{action: "toArray", replyChan: replyChan}
	return (<-replyChan).([]T)
}

func (v *GenericVector[T]) Add(value T) error {
	return v.add(value)
}

func (v *GenericVector[T]) AddAt(index int, value T) error {
	return v.addAt(index, value)
}

func (v *GenericVector[T]) RemoveAt(index int) error {
	return v.removeAt(index)
}

func (v *GenericVector[T]) Set(index int, value T) error {
	return v.set(index, value)
}

func (v *GenericVector[T]) Get(index int) (T, error) {
	return v.get(index)
}

func (v *GenericVector[T]) Size() int {
	return v.size()
}

func (v *GenericVector[T]) IsEmpty() bool {
	return v.isEmpty()
}

func (v *GenericVector[T]) Clear() error {
	return v.clear()
}

func (v *GenericVector[T]) ToArray() []T {
	return v.toArray()
}

func (v *GenericVector[T]) EnsureCapacity(minCapacity int) {
	v.ensureCapacity(minCapacity)
}

func (v *GenericVector[T]) TrimToSize() error {
	return v.trimToSize()
}
