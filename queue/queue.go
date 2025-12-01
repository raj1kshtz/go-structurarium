package queue

import (
	"github.com/raj1kshtz/go-structurarium/collection"
)

type queueRequest[T comparable] struct {
	action    string
	value     T
	replyChan chan interface{}
}

type GenericQueue[T comparable] struct {
	queueChan         chan queueRequest[T]
	collectionWrapper *collection.GenericCollectionWrapper[T]
}

func NewGenericQueue[T comparable](initialCapacity ...int) *GenericQueue[T] {
	var capacity int
	if len(initialCapacity) > 0 {
		capacity = initialCapacity[0]
	} else {
		capacity = 0
	}
	s := &GenericQueue[T]{
		queueChan:         make(chan queueRequest[T]),
		collectionWrapper: collection.NewGenericCollectionWrapper[T](capacity),
	}

	go s.manageQueue()
	return s
}

func (q *GenericQueue[T]) manageQueue() {
	for req := range q.queueChan {
		switch req.action {
		case "enQueue":
			if ok := q.enQueue(req.value); !ok {
				req.replyChan <- "Unable to push data to queue"
			} else {
				req.replyChan <- true
			}
		case "deQueue":
			val, ok := q.deQueue()
			req.replyChan <- val
			req.replyChan <- ok
		case "size":
			req.replyChan <- q.size()
		case "isEmpty":
			req.replyChan <- q.isEmpty()
		case "peek":
			req.replyChan <- q.peek()
		case "clear":
			q.clear()
		case "toArray":
			req.replyChan <- q.toArray()

		}
	}
}

func (q *GenericQueue[T]) enQueue(value T) bool {
	return q.collectionWrapper.Add(value)
}

func (q *GenericQueue[T]) deQueue() (T, bool) {
	var zero T
	if q.isEmpty() {
		return zero, false
	}
	frontElement := q.peek()
	if frontElement == zero {
		return zero, false
	}
	success := q.collectionWrapper.Remove(frontElement)
	if success {
		return frontElement, true
	}
	return zero, false
}

func (q *GenericQueue[T]) size() int {
	return q.collectionWrapper.Size()
}

func (q *GenericQueue[T]) isEmpty() bool {
	return q.collectionWrapper.IsEmpty()
}

func (q *GenericQueue[T]) peek() T {
	if q.isEmpty() {
		var zeroValue T
		return zeroValue
	}
	return q.collectionWrapper.ToArray()[0]
}

func (q *GenericQueue[T]) clear() {
	q.collectionWrapper.Clear()
}

func (q *GenericQueue[T]) toArray() []T {
	return q.collectionWrapper.ToArray()
}

func replyChanReceive(replyChan chan interface{}) error {
	result := <-replyChan
	if err, ok := result.(error); ok {
		return err
	}
	return nil
}
