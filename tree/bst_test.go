package tree

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GenericBSTTestSuite struct {
	suite.Suite
	intBST *GenericBST[int]
}

func TestGenericBSTTestSuite(t *testing.T) {
	suite.Run(t, new(GenericBSTTestSuite))
}

func (s *GenericBSTTestSuite) SetupTest() {
	s.intBST = NewGenericBST[int]()
}

func (s *GenericBSTTestSuite) TestInsertAndSearch() {
	s.intBST.insert(5)
	s.intBST.insert(3)
	s.intBST.insert(7)
	s.intBST.insert(2)
	s.intBST.insert(4)

	s.Equal(5, s.intBST.size)
	s.True(s.intBST.search(5))
	s.True(s.intBST.search(3))
	s.True(s.intBST.search(7))
	s.False(s.intBST.search(10))
}

func (s *GenericBSTTestSuite) TestDelete() {
	s.intBST.insert(5)
	s.intBST.insert(3)
	s.intBST.insert(7)
	s.intBST.insert(2)
	s.intBST.insert(4)
	s.intBST.insert(6)
	s.intBST.insert(8)

	// Delete leaf node
	deleted := s.intBST.delete(2)
	s.True(deleted)
	s.False(s.intBST.search(2))
	s.Equal(6, s.intBST.size)

	// Delete node with one child
	deleted = s.intBST.delete(3)
	s.True(deleted)
	s.False(s.intBST.search(3))
	s.True(s.intBST.search(4))

	// Delete node with two children
	deleted = s.intBST.delete(7)
	s.True(deleted)
	s.False(s.intBST.search(7))
	s.True(s.intBST.search(6))
	s.True(s.intBST.search(8))
}

func (s *GenericBSTTestSuite) TestMinMax() {
	s.intBST.insert(5)
	s.intBST.insert(3)
	s.intBST.insert(7)
	s.intBST.insert(2)
	s.intBST.insert(8)

	min, err := s.intBST.min()
	s.NoError(err)
	s.Equal(2, min)

	max, err := s.intBST.max()
	s.NoError(err)
	s.Equal(8, max)
}

func (s *GenericBSTTestSuite) TestMinMaxEmptyTree() {
	_, err := s.intBST.min()
	s.Error(err)

	_, err = s.intBST.max()
	s.Error(err)
}

func (s *GenericBSTTestSuite) TestInOrder() {
	s.intBST.insert(5)
	s.intBST.insert(3)
	s.intBST.insert(7)
	s.intBST.insert(2)
	s.intBST.insert(4)
	s.intBST.insert(6)
	s.intBST.insert(8)

	result := s.intBST.inOrder()
	s.Equal([]int{2, 3, 4, 5, 6, 7, 8}, result)
}

func (s *GenericBSTTestSuite) TestPreOrder() {
	s.intBST.insert(5)
	s.intBST.insert(3)
	s.intBST.insert(7)

	result := s.intBST.preOrder()
	s.Equal([]int{5, 3, 7}, result)
}

func (s *GenericBSTTestSuite) TestPostOrder() {
	s.intBST.insert(5)
	s.intBST.insert(3)
	s.intBST.insert(7)

	result := s.intBST.postOrder()
	s.Equal([]int{3, 7, 5}, result)
}

func (s *GenericBSTTestSuite) TestLevelOrder() {
	s.intBST.insert(5)
	s.intBST.insert(3)
	s.intBST.insert(7)
	s.intBST.insert(2)
	s.intBST.insert(4)

	result := s.intBST.levelOrder()
	s.Equal([]int{5, 3, 7, 2, 4}, result)
}

func (s *GenericBSTTestSuite) TestHeight() {
	s.Equal(0, s.intBST.height())

	s.intBST.insert(5)
	s.Equal(1, s.intBST.height())

	s.intBST.insert(3)
	s.intBST.insert(7)
	s.Equal(2, s.intBST.height())

	s.intBST.insert(2)
	s.Equal(3, s.intBST.height())
}

func (s *GenericBSTTestSuite) TestClear() {
	s.intBST.insert(5)
	s.intBST.insert(3)
	s.intBST.insert(7)

	s.intBST.clear()
	s.Equal(0, s.intBST.size)
	s.Nil(s.intBST.root)
}

func (s *GenericBSTTestSuite) TestStringBST() {
	strBST := NewGenericBST[string]()

	strBST.insert("dog")
	strBST.insert("cat")
	strBST.insert("elephant")
	strBST.insert("ant")

	s.Equal(4, strBST.size)
	result := strBST.inOrder()
	s.Equal([]string{"ant", "cat", "dog", "elephant"}, result)
}
