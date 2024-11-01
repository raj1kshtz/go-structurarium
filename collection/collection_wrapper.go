package collection

type GenericCollectionWrapper[T comparable] struct {
	collection *GenericCollection[T]
}

func NewGenericCollectionWrapper[T comparable](initialCapacity ...int) *GenericCollectionWrapper[T] {
	return &GenericCollectionWrapper[T]{collection: NewGenericCollection[T](initialCapacity...)}
}

func (w *GenericCollectionWrapper[T]) Add(value T) bool {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "add", values: []T{value}, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericCollectionWrapper[T]) AddAll(values []T) bool {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "addAll", values: values, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericCollectionWrapper[T]) Remove(value T) bool {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "remove", values: []T{value}, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericCollectionWrapper[T]) RemoveAll(values []T) bool {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "removeAll", values: values, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericCollectionWrapper[T]) RetainAll(values []T) bool {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "retainAll", values: values, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericCollectionWrapper[T]) Contains(value T) bool {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "contains", values: []T{value}, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericCollectionWrapper[T]) ContainsAll(values []T) bool {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "containsAll", values: values, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericCollectionWrapper[T]) Size() int {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "size", replyChan: replyChan}
	return (<-replyChan).(int)
}

func (w *GenericCollectionWrapper[T]) IsEmpty() bool {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "isEmpty", replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericCollectionWrapper[T]) Clear() {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "clear", replyChan: replyChan}
	<-replyChan
}

func (w *GenericCollectionWrapper[T]) ToArray() []T {
	replyChan := make(chan interface{})
	w.collection.collectionChan <- collectionRequest[T]{action: "toArray", replyChan: replyChan}
	return (<-replyChan).([]T)
}
