package tree

type TreeWrapper[T comparable] struct {
	tree *GenericTree[T]
}

func NewTreeWrapper[T comparable](rootValue T) *TreeWrapper[T] {
	return &TreeWrapper[T]{
		tree: NewGenericTree[T](rootValue),
	}
}

func (tw *TreeWrapper[T]) Insert(parentValue T, value T) error {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "insert", parentVal: parentValue, value: value, replyChan: replyChan}
	result := <-replyChan
	if err, ok := result.(error); ok {
		return err
	}
	return nil
}

func (tw *TreeWrapper[T]) Remove(value T) bool {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "remove", value: value, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (tw *TreeWrapper[T]) Search(value T) bool {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "search", value: value, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (tw *TreeWrapper[T]) PreOrder() []T {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "preorder", replyChan: replyChan}
	return (<-replyChan).([]T)
}

func (tw *TreeWrapper[T]) PostOrder() []T {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "postorder", replyChan: replyChan}
	return (<-replyChan).([]T)
}

func (tw *TreeWrapper[T]) LevelOrder() []T {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "levelorder", replyChan: replyChan}
	return (<-replyChan).([]T)
}

func (tw *TreeWrapper[T]) Height() int {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "height", replyChan: replyChan}
	return (<-replyChan).(int)
}

func (tw *TreeWrapper[T]) Size() int {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "size", replyChan: replyChan}
	return (<-replyChan).(int)
}

func (tw *TreeWrapper[T]) IsEmpty() bool {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "isEmpty", replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (tw *TreeWrapper[T]) Clear() {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "clear", replyChan: replyChan}
	<-replyChan
}

func (tw *TreeWrapper[T]) GetRoot() T {
	replyChan := make(chan interface{})
	tw.tree.treeChan <- treeRequest[T]{action: "getRoot", replyChan: replyChan}
	return (<-replyChan).(T)
}
