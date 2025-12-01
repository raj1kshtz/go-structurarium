package tree

type BSTWrapper[T Ordered] struct {
	bst *GenericBST[T]
}

func NewBSTWrapper[T Ordered]() *BSTWrapper[T] {
	return &BSTWrapper[T]{
		bst: NewGenericBST[T](),
	}
}

func (bw *BSTWrapper[T]) Insert(value T) {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "insert", value: value, replyChan: replyChan}
	<-replyChan
}

func (bw *BSTWrapper[T]) Delete(value T) bool {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "delete", value: value, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (bw *BSTWrapper[T]) Search(value T) bool {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "search", value: value, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (bw *BSTWrapper[T]) Min() (T, error) {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "min", replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

func (bw *BSTWrapper[T]) Max() (T, error) {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "max", replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

func (bw *BSTWrapper[T]) InOrder() []T {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "inorder", replyChan: replyChan}
	return (<-replyChan).([]T)
}

func (bw *BSTWrapper[T]) PreOrder() []T {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "preorder", replyChan: replyChan}
	return (<-replyChan).([]T)
}

func (bw *BSTWrapper[T]) PostOrder() []T {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "postorder", replyChan: replyChan}
	return (<-replyChan).([]T)
}

func (bw *BSTWrapper[T]) LevelOrder() []T {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "levelorder", replyChan: replyChan}
	return (<-replyChan).([]T)
}

func (bw *BSTWrapper[T]) Height() int {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "height", replyChan: replyChan}
	return (<-replyChan).(int)
}

func (bw *BSTWrapper[T]) Size() int {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "size", replyChan: replyChan}
	return (<-replyChan).(int)
}

func (bw *BSTWrapper[T]) IsEmpty() bool {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "isEmpty", replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (bw *BSTWrapper[T]) Clear() {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "clear", replyChan: replyChan}
	<-replyChan
}

func (bw *BSTWrapper[T]) Validate() bool {
	replyChan := make(chan interface{})
	bw.bst.bstChan <- bstRequest[T]{action: "validate", replyChan: replyChan}
	return (<-replyChan).(bool)
}
