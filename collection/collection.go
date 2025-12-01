package collection

type collectionRequest[T any] struct {
	action    string
	values    []T
	replyChan chan interface{}
}

type GenericCollection[T comparable] struct {
	collectionChan chan collectionRequest[T]
	elements       []T
}

func NewGenericCollection[T comparable](initialCapacity ...int) *GenericCollection[T] {
	var capacity int
	if len(initialCapacity) > 0 {
		capacity = initialCapacity[0]
	} else {
		capacity = 0
	}
	c := &GenericCollection[T]{
		collectionChan: make(chan collectionRequest[T]),
		elements:       make([]T, 0, capacity),
	}
	go c.manageCollection()
	return c
}

func (c *GenericCollection[T]) manageCollection() {
	for req := range c.collectionChan {
		switch req.action {
		case "add":
			c.add(req.values[0])
			req.replyChan <- true
		case "addAll":
			c.addAll(req.values)
			req.replyChan <- true
		case "remove":
			req.replyChan <- c.remove(req.values[0])
		case "removeAll":
			req.replyChan <- c.removeAll(req.values)
		case "retainAll":
			req.replyChan <- c.retainAll(req.values)
		case "contains":
			req.replyChan <- c.contains(req.values[0])
		case "containsAll":
			req.replyChan <- c.containsAll(req.values)
		case "size":
			req.replyChan <- c.size()
		case "isEmpty":
			req.replyChan <- c.isEmpty()
		case "clear":
			c.clear()
			req.replyChan <- true
		case "toArray":
			req.replyChan <- c.toArray()
		}
	}
}

func (c *GenericCollection[T]) add(value T) {
	c.elements = append(c.elements, value)
}

func (c *GenericCollection[T]) addAll(values []T) {
	c.elements = append(c.elements, values...)
}

func (c *GenericCollection[T]) remove(value T) bool {
	for i, elem := range c.elements {
		if elem == value { // Now valid since T is comparable
			c.elements = append(c.elements[:i], c.elements[i+1:]...)
			return true
		}
	}
	return false
}

func (c *GenericCollection[T]) removeAll(values []T) bool {
	removed := false
	for _, val := range values {
		if c.remove(val) {
			removed = true
		}
	}
	return removed
}

func (c *GenericCollection[T]) retainAll(values []T) bool {
	valueMap := make(map[T]struct{})
	for _, val := range values {
		valueMap[val] = struct{}{}
	}

	newElements := []T{}
	for _, elem := range c.elements {
		if _, found := valueMap[elem]; found {
			newElements = append(newElements, elem)
		}
	}
	changed := len(newElements) != len(c.elements)
	c.elements = newElements
	return changed
}

func (c *GenericCollection[T]) contains(value T) bool {
	for _, elem := range c.elements {
		if elem == value {
			return true
		}
	}
	return false
}

func (c *GenericCollection[T]) containsAll(values []T) bool {
	for _, val := range values {
		if !c.contains(val) {
			return false
		}
	}
	return true
}

func (c *GenericCollection[T]) size() int {
	return len(c.elements)
}

func (c *GenericCollection[T]) isEmpty() bool {
	return len(c.elements) == 0
}

func (c *GenericCollection[T]) clear() {
	c.elements = []T{}
}

func (c *GenericCollection[T]) toArray() []T {
	return append([]T{}, c.elements...)
}
