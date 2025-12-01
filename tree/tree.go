package tree

import "fmt"

type TreeNode[T comparable] struct {
	Value    T
	Children []*TreeNode[T]
	Parent   *TreeNode[T]
}

type treeRequest[T comparable] struct {
	action    string
	value     T
	parentVal T
	callback  func(*TreeNode[T])
	replyChan chan interface{}
}

type TraversalOrder int

const (
	PreOrder TraversalOrder = iota
	PostOrder
	LevelOrder
)

type GenericTree[T comparable] struct {
	treeChan chan treeRequest[T]
	root     *TreeNode[T]
	size     int
}

func NewGenericTree[T comparable](rootValue T) *GenericTree[T] {
	t := &GenericTree[T]{
		treeChan: make(chan treeRequest[T]),
		root: &TreeNode[T]{
			Value:    rootValue,
			Children: make([]*TreeNode[T], 0),
			Parent:   nil,
		},
		size: 1,
	}
	go t.manageTree()
	return t
}

func (t *GenericTree[T]) manageTree() {
	for req := range t.treeChan {
		switch req.action {
		case "insert":
			req.replyChan <- t.insert(req.parentVal, req.value)
		case "remove":
			req.replyChan <- t.remove(req.value)
		case "search":
			node := t.search(req.value)
			req.replyChan <- (node != nil)
		case "find":
			node := t.search(req.value)
			req.replyChan <- node
		case "traverse":
			order := PreOrder // Default to pre-order
			req.replyChan <- t.traverse(order)
		case "preorder":
			req.replyChan <- t.preOrder()
		case "postorder":
			req.replyChan <- t.postOrder()
		case "levelorder":
			req.replyChan <- t.levelOrder()
		case "height":
			req.replyChan <- t.height()
		case "size":
			req.replyChan <- t.size
		case "isEmpty":
			req.replyChan <- t.size == 0
		case "clear":
			t.clear()
			req.replyChan <- true
		case "getRoot":
			if t.root != nil {
				req.replyChan <- t.root.Value
			} else {
				var zero T
				req.replyChan <- zero
			}
		}
	}
}

func (t *GenericTree[T]) insert(parentValue T, value T) error {
	parentNode := t.search(parentValue)
	if parentNode == nil {
		return fmt.Errorf("parent node with value %v not found", parentValue)
	}

	newNode := &TreeNode[T]{
		Value:    value,
		Children: make([]*TreeNode[T], 0),
		Parent:   parentNode,
	}
	parentNode.Children = append(parentNode.Children, newNode)
	t.size++
	return nil
}

func (t *GenericTree[T]) remove(value T) bool {
	if t.root == nil || t.root.Value == value {
		return false
	}

	node := t.search(value)
	if node == nil {
		return false
	}

	parent := node.Parent
	if parent != nil {
		for i, child := range parent.Children {
			if child.Value == value {
				parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
				break
			}
		}
	}

	t.size -= t.countNodes(node)
	return true
}

func (t *GenericTree[T]) search(value T) *TreeNode[T] {
	if t.root == nil {
		return nil
	}
	return t.searchHelper(t.root, value)
}

func (t *GenericTree[T]) searchHelper(node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	for _, child := range node.Children {
		if result := t.searchHelper(child, value); result != nil {
			return result
		}
	}

	return nil
}

func (t *GenericTree[T]) preOrder() []T {
	result := make([]T, 0)
	t.preOrderHelper(t.root, &result)
	return result
}

func (t *GenericTree[T]) preOrderHelper(node *TreeNode[T], result *[]T) {
	if node == nil {
		return
	}

	*result = append(*result, node.Value)
	for _, child := range node.Children {
		t.preOrderHelper(child, result)
	}
}

func (t *GenericTree[T]) postOrder() []T {
	result := make([]T, 0)
	t.postOrderHelper(t.root, &result)
	return result
}

func (t *GenericTree[T]) postOrderHelper(node *TreeNode[T], result *[]T) {
	if node == nil {
		return
	}

	for _, child := range node.Children {
		t.postOrderHelper(child, result)
	}
	*result = append(*result, node.Value)
}

func (t *GenericTree[T]) levelOrder() []T {
	if t.root == nil {
		return []T{}
	}

	result := make([]T, 0)
	queue := []*TreeNode[T]{t.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Value)

		queue = append(queue, node.Children...)
	}

	return result
}

func (t *GenericTree[T]) traverse(order TraversalOrder) []T {
	switch order {
	case PreOrder:
		return t.preOrder()
	case PostOrder:
		return t.postOrder()
	case LevelOrder:
		return t.levelOrder()
	default:
		return t.preOrder()
	}
}

func (t *GenericTree[T]) height() int {
	if t.root == nil {
		return 0
	}
	return t.heightHelper(t.root)
}

func (t *GenericTree[T]) heightHelper(node *TreeNode[T]) int {
	if node == nil {
		return 0
	}

	if len(node.Children) == 0 {
		return 1
	}

	maxHeight := 0
	for _, child := range node.Children {
		childHeight := t.heightHelper(child)
		if childHeight > maxHeight {
			maxHeight = childHeight
		}
	}

	return maxHeight + 1
}

func (t *GenericTree[T]) countNodes(node *TreeNode[T]) int {
	if node == nil {
		return 0
	}

	count := 1
	for _, child := range node.Children {
		count += t.countNodes(child)
	}

	return count
}

func (t *GenericTree[T]) clear() {
	var zero T
	t.root = &TreeNode[T]{
		Value:    zero,
		Children: make([]*TreeNode[T], 0),
		Parent:   nil,
	}
	t.size = 0
}

func replyChanReceive(replyChan chan interface{}) error {
	result := <-replyChan
	if err, ok := result.(error); ok {
		return err
	}
	return nil
}
