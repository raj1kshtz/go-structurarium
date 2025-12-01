package tree

import "fmt"

type BSTNode[T comparable] struct {
	Value T
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

type bstRequest[T comparable] struct {
	action    string
	value     T
	replyChan chan interface{}
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

type GenericBST[T Ordered] struct {
	bstChan chan bstRequest[T]
	root    *BSTNode[T]
	size    int
}

func NewGenericBST[T Ordered]() *GenericBST[T] {
	bst := &GenericBST[T]{
		bstChan: make(chan bstRequest[T]),
		root:    nil,
		size:    0,
	}
	go bst.manageBST()
	return bst
}

func (bst *GenericBST[T]) manageBST() {
	for req := range bst.bstChan {
		switch req.action {
		case "insert":
			bst.insert(req.value)
			req.replyChan <- true
		case "delete":
			req.replyChan <- bst.delete(req.value)
		case "search":
			req.replyChan <- bst.search(req.value)
		case "min":
			value, err := bst.min()
			if err != nil {
				req.replyChan <- err
			} else {
				req.replyChan <- value
			}
		case "max":
			value, err := bst.max()
			if err != nil {
				req.replyChan <- err
			} else {
				req.replyChan <- value
			}
		case "inorder":
			req.replyChan <- bst.inOrder()
		case "preorder":
			req.replyChan <- bst.preOrder()
		case "postorder":
			req.replyChan <- bst.postOrder()
		case "levelorder":
			req.replyChan <- bst.levelOrder()
		case "height":
			req.replyChan <- bst.height()
		case "size":
			req.replyChan <- bst.size
		case "isEmpty":
			req.replyChan <- (bst.size == 0)
		case "clear":
			bst.clear()
			req.replyChan <- true
		case "validate":
			req.replyChan <- bst.validate()
		}
	}
}

func (bst *GenericBST[T]) insert(value T) {
	if bst.root == nil {
		bst.root = &BSTNode[T]{Value: value}
		bst.size++
		return
	}
	bst.insertHelper(bst.root, value)
}

func (bst *GenericBST[T]) insertHelper(node *BSTNode[T], value T) *BSTNode[T] {
	if node == nil {
		bst.size++
		return &BSTNode[T]{Value: value}
	}

	if value < node.Value {
		node.Left = bst.insertHelper(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.insertHelper(node.Right, value)
	}

	return node
}

func (bst *GenericBST[T]) delete(value T) bool {
	if bst.root == nil {
		return false
	}

	var deleted bool
	bst.root, deleted = bst.deleteHelper(bst.root, value)
	if deleted {
		bst.size--
	}
	return deleted
}

func (bst *GenericBST[T]) deleteHelper(node *BSTNode[T], value T) (*BSTNode[T], bool) {
	if node == nil {
		return nil, false
	}

	var deleted bool

	if value < node.Value {
		node.Left, deleted = bst.deleteHelper(node.Left, value)
	} else if value > node.Value {
		node.Right, deleted = bst.deleteHelper(node.Right, value)
	} else {
		deleted = true

		if node.Left == nil && node.Right == nil {
			return nil, true
		}

		if node.Left == nil {
			return node.Right, true
		}
		if node.Right == nil {
			return node.Left, true
		}

		successor := bst.findMin(node.Right)
		node.Value = successor.Value
		node.Right, _ = bst.deleteHelper(node.Right, successor.Value)
	}

	return node, deleted
}

func (bst *GenericBST[T]) search(value T) bool {
	return bst.searchHelper(bst.root, value)
}

func (bst *GenericBST[T]) searchHelper(node *BSTNode[T], value T) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	} else if value < node.Value {
		return bst.searchHelper(node.Left, value)
	} else {
		return bst.searchHelper(node.Right, value)
	}
}

func (bst *GenericBST[T]) min() (T, error) {
	if bst.root == nil {
		var zero T
		return zero, fmt.Errorf("tree is empty")
	}
	return bst.findMin(bst.root).Value, nil
}

func (bst *GenericBST[T]) findMin(node *BSTNode[T]) *BSTNode[T] {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (bst *GenericBST[T]) max() (T, error) {
	if bst.root == nil {
		var zero T
		return zero, fmt.Errorf("tree is empty")
	}
	return bst.findMax(bst.root).Value, nil
}

func (bst *GenericBST[T]) findMax(node *BSTNode[T]) *BSTNode[T] {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (bst *GenericBST[T]) inOrder() []T {
	result := make([]T, 0)
	bst.inOrderHelper(bst.root, &result)
	return result
}

func (bst *GenericBST[T]) inOrderHelper(node *BSTNode[T], result *[]T) {
	if node == nil {
		return
	}
	bst.inOrderHelper(node.Left, result)
	*result = append(*result, node.Value)
	bst.inOrderHelper(node.Right, result)
}

func (bst *GenericBST[T]) preOrder() []T {
	result := make([]T, 0)
	bst.preOrderHelper(bst.root, &result)
	return result
}

func (bst *GenericBST[T]) preOrderHelper(node *BSTNode[T], result *[]T) {
	if node == nil {
		return
	}
	*result = append(*result, node.Value)
	bst.preOrderHelper(node.Left, result)
	bst.preOrderHelper(node.Right, result)
}

func (bst *GenericBST[T]) postOrder() []T {
	result := make([]T, 0)
	bst.postOrderHelper(bst.root, &result)
	return result
}

func (bst *GenericBST[T]) postOrderHelper(node *BSTNode[T], result *[]T) {
	if node == nil {
		return
	}
	bst.postOrderHelper(node.Left, result)
	bst.postOrderHelper(node.Right, result)
	*result = append(*result, node.Value)
}

func (bst *GenericBST[T]) levelOrder() []T {
	if bst.root == nil {
		return []T{}
	}

	result := make([]T, 0)
	queue := []*BSTNode[T]{bst.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func (bst *GenericBST[T]) height() int {
	return bst.heightHelper(bst.root)
}

func (bst *GenericBST[T]) heightHelper(node *BSTNode[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := bst.heightHelper(node.Left)
	rightHeight := bst.heightHelper(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (bst *GenericBST[T]) clear() {
	bst.root = nil
	bst.size = 0
}

func (bst *GenericBST[T]) validate() bool {
	if bst.root == nil {
		return true
	}
	return bst.validateHelper(bst.root, nil, nil)
}

func (bst *GenericBST[T]) validateHelper(node *BSTNode[T], min, max *T) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Value <= *min {
		return false
	}

	if max != nil && node.Value >= *max {
		return false
	}

	leftValid := bst.validateHelper(node.Left, min, &node.Value)
	rightValid := bst.validateHelper(node.Right, &node.Value, max)

	return leftValid && rightValid
}
