package tree

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GenericTreeTestSuite struct {
	suite.Suite
	intTree *GenericTree[int]
}

func TestGenericTreeTestSuite(t *testing.T) {
	suite.Run(t, new(GenericTreeTestSuite))
}

func (s *GenericTreeTestSuite) SetupTest() {
	s.intTree = NewGenericTree[int](1)
}

func (s *GenericTreeTestSuite) TestTreeCreation() {
	s.NotNil(s.intTree)
	s.NotNil(s.intTree.root)
	s.Equal(1, s.intTree.root.Value)
	s.Equal(1, s.intTree.size)
}

func (s *GenericTreeTestSuite) TestInsert() {
	// Insert children of root
	err := s.intTree.insert(1, 2)
	s.NoError(err)
	err = s.intTree.insert(1, 3)
	s.NoError(err)
	err = s.intTree.insert(1, 4)
	s.NoError(err)

	s.Equal(4, s.intTree.size)
	s.Equal(3, len(s.intTree.root.Children))

	// Insert grandchildren
	err = s.intTree.insert(2, 5)
	s.NoError(err)
	err = s.intTree.insert(2, 6)
	s.NoError(err)

	s.Equal(6, s.intTree.size)
}

func (s *GenericTreeTestSuite) TestInsertInvalidParent() {
	err := s.intTree.insert(99, 2)
	s.Error(err)
	s.Equal(1, s.intTree.size)
}

func (s *GenericTreeTestSuite) TestSearch() {
	s.intTree.insert(1, 2)
	s.intTree.insert(1, 3)
	s.intTree.insert(2, 4)

	node := s.intTree.search(1)
	s.NotNil(node)
	s.Equal(1, node.Value)

	node = s.intTree.search(4)
	s.NotNil(node)
	s.Equal(4, node.Value)

	node = s.intTree.search(99)
	s.Nil(node)
}

func (s *GenericTreeTestSuite) TestRemove() {
	s.intTree.insert(1, 2)
	s.intTree.insert(1, 3)
	s.intTree.insert(2, 4)
	s.intTree.insert(2, 5)

	// Remove node with children (should remove subtree)
	removed := s.intTree.remove(2)
	s.True(removed)
	s.Equal(2, s.intTree.size) // Only root and node 3 remain

	// Try to remove non-existent node
	removed = s.intTree.remove(99)
	s.False(removed)

	// Try to remove root (should fail)
	removed = s.intTree.remove(1)
	s.False(removed)
}

func (s *GenericTreeTestSuite) TestPreOrder() {
	s.intTree.insert(1, 2)
	s.intTree.insert(1, 3)
	s.intTree.insert(2, 4)
	s.intTree.insert(2, 5)
	s.intTree.insert(3, 6)

	result := s.intTree.preOrder()
	s.Equal([]int{1, 2, 4, 5, 3, 6}, result)
}

func (s *GenericTreeTestSuite) TestPostOrder() {
	s.intTree.insert(1, 2)
	s.intTree.insert(1, 3)
	s.intTree.insert(2, 4)
	s.intTree.insert(2, 5)
	s.intTree.insert(3, 6)

	result := s.intTree.postOrder()
	s.Equal([]int{4, 5, 2, 6, 3, 1}, result)
}

func (s *GenericTreeTestSuite) TestLevelOrder() {
	s.intTree.insert(1, 2)
	s.intTree.insert(1, 3)
	s.intTree.insert(2, 4)
	s.intTree.insert(2, 5)
	s.intTree.insert(3, 6)

	result := s.intTree.levelOrder()
	s.Equal([]int{1, 2, 3, 4, 5, 6}, result)
}

func (s *GenericTreeTestSuite) TestHeight() {
	// Single node tree
	height := s.intTree.height()
	s.Equal(1, height)

	// Add one level
	s.intTree.insert(1, 2)
	s.intTree.insert(1, 3)
	height = s.intTree.height()
	s.Equal(2, height)

	// Add another level
	s.intTree.insert(2, 4)
	height = s.intTree.height()
	s.Equal(3, height)
}

func (s *GenericTreeTestSuite) TestSize() {
	s.Equal(1, s.intTree.size)

	s.intTree.insert(1, 2)
	s.Equal(2, s.intTree.size)

	s.intTree.insert(1, 3)
	s.intTree.insert(2, 4)
	s.Equal(4, s.intTree.size)
}

func (s *GenericTreeTestSuite) TestClear() {
	s.intTree.insert(1, 2)
	s.intTree.insert(1, 3)
	s.intTree.insert(2, 4)

	s.intTree.clear()
	s.Equal(0, s.intTree.size)
}

func (s *GenericTreeTestSuite) TestStringTree() {
	strTree := NewGenericTree[string]("root")

	err := strTree.insert("root", "child1")
	s.NoError(err)
	err = strTree.insert("root", "child2")
	s.NoError(err)
	err = strTree.insert("child1", "grandchild")
	s.NoError(err)

	s.Equal(4, strTree.size)
	result := strTree.levelOrder()
	s.Equal([]string{"root", "child1", "child2", "grandchild"}, result)
}
