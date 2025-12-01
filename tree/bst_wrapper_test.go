package tree

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BSTWrapperTestSuite struct {
	suite.Suite
	bstWrapper *BSTWrapper[int]
}

func TestBSTWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(BSTWrapperTestSuite))
}

func (s *BSTWrapperTestSuite) SetupTest() {
	s.bstWrapper = NewBSTWrapper[int]()
}

func (s *BSTWrapperTestSuite) TestInsertAndSearch() {
	s.bstWrapper.Insert(5)
	s.bstWrapper.Insert(3)
	s.bstWrapper.Insert(7)
	s.bstWrapper.Insert(2)
	s.bstWrapper.Insert(4)

	s.Equal(5, s.bstWrapper.Size())
	s.True(s.bstWrapper.Search(5))
	s.True(s.bstWrapper.Search(3))
	s.False(s.bstWrapper.Search(10))
}

func (s *BSTWrapperTestSuite) TestDelete() {
	s.bstWrapper.Insert(5)
	s.bstWrapper.Insert(3)
	s.bstWrapper.Insert(7)
	s.bstWrapper.Insert(2)
	s.bstWrapper.Insert(4)

	deleted := s.bstWrapper.Delete(2)
	s.True(deleted)
	s.False(s.bstWrapper.Search(2))
	s.Equal(4, s.bstWrapper.Size())

	deleted = s.bstWrapper.Delete(10)
	s.False(deleted)
}

func (s *BSTWrapperTestSuite) TestMinMax() {
	s.bstWrapper.Insert(5)
	s.bstWrapper.Insert(3)
	s.bstWrapper.Insert(7)
	s.bstWrapper.Insert(2)
	s.bstWrapper.Insert(8)

	min, err := s.bstWrapper.Min()
	s.NoError(err)
	s.Equal(2, min)

	max, err := s.bstWrapper.Max()
	s.NoError(err)
	s.Equal(8, max)
}

func (s *BSTWrapperTestSuite) TestInOrderTraversal() {
	s.bstWrapper.Insert(5)
	s.bstWrapper.Insert(3)
	s.bstWrapper.Insert(7)
	s.bstWrapper.Insert(2)
	s.bstWrapper.Insert(4)
	s.bstWrapper.Insert(6)
	s.bstWrapper.Insert(8)

	result := s.bstWrapper.InOrder()
	s.Equal([]int{2, 3, 4, 5, 6, 7, 8}, result)
}

func (s *BSTWrapperTestSuite) TestAllTraversals() {
	s.bstWrapper.Insert(5)
	s.bstWrapper.Insert(3)
	s.bstWrapper.Insert(7)

	preOrder := s.bstWrapper.PreOrder()
	s.Equal([]int{5, 3, 7}, preOrder)

	postOrder := s.bstWrapper.PostOrder()
	s.Equal([]int{3, 7, 5}, postOrder)

	levelOrder := s.bstWrapper.LevelOrder()
	s.Equal([]int{5, 3, 7}, levelOrder)
}

func (s *BSTWrapperTestSuite) TestHeight() {
	s.Equal(0, s.bstWrapper.Height())

	s.bstWrapper.Insert(5)
	s.Equal(1, s.bstWrapper.Height())

	s.bstWrapper.Insert(3)
	s.bstWrapper.Insert(7)
	s.Equal(2, s.bstWrapper.Height())
}

func (s *BSTWrapperTestSuite) TestIsEmpty() {
	s.True(s.bstWrapper.IsEmpty())

	s.bstWrapper.Insert(5)
	s.False(s.bstWrapper.IsEmpty())

	s.bstWrapper.Clear()
	s.True(s.bstWrapper.IsEmpty())
}

func (s *BSTWrapperTestSuite) TestClear() {
	s.bstWrapper.Insert(5)
	s.bstWrapper.Insert(3)
	s.bstWrapper.Insert(7)

	s.bstWrapper.Clear()
	s.Equal(0, s.bstWrapper.Size())
	s.True(s.bstWrapper.IsEmpty())
}

func (s *BSTWrapperTestSuite) TestValidate() {
	s.bstWrapper.Insert(5)
	s.bstWrapper.Insert(3)
	s.bstWrapper.Insert(7)
	s.bstWrapper.Insert(2)
	s.bstWrapper.Insert(4)

	s.True(s.bstWrapper.Validate())
}

func (s *BSTWrapperTestSuite) TestStringBSTWrapper() {
	strWrapper := NewBSTWrapper[string]()

	strWrapper.Insert("dog")
	strWrapper.Insert("cat")
	strWrapper.Insert("elephant")
	strWrapper.Insert("ant")

	s.Equal(4, strWrapper.Size())
	s.True(strWrapper.Search("cat"))
	s.False(strWrapper.Search("zebra"))

	result := strWrapper.InOrder()
	s.Equal([]string{"ant", "cat", "dog", "elephant"}, result)
}
