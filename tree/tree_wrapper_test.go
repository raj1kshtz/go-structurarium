package tree

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TreeWrapperTestSuite struct {
	suite.Suite
	treeWrapper *TreeWrapper[int]
}

func TestTreeWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(TreeWrapperTestSuite))
}

func (s *TreeWrapperTestSuite) SetupTest() {
	s.treeWrapper = NewTreeWrapper[int](1)
}

func (s *TreeWrapperTestSuite) TestWrapperCreation() {
	s.NotNil(s.treeWrapper)
	s.Equal(1, s.treeWrapper.GetRoot())
	s.Equal(1, s.treeWrapper.Size())
}

func (s *TreeWrapperTestSuite) TestInsert() {
	err := s.treeWrapper.Insert(1, 2)
	s.NoError(err)
	err = s.treeWrapper.Insert(1, 3)
	s.NoError(err)
	err = s.treeWrapper.Insert(2, 4)
	s.NoError(err)

	s.Equal(4, s.treeWrapper.Size())
}

func (s *TreeWrapperTestSuite) TestInsertInvalidParent() {
	err := s.treeWrapper.Insert(99, 2)
	s.Error(err)
	s.Equal(1, s.treeWrapper.Size())
}

func (s *TreeWrapperTestSuite) TestSearch() {
	s.treeWrapper.Insert(1, 2)
	s.treeWrapper.Insert(1, 3)

	s.True(s.treeWrapper.Search(1))
	s.True(s.treeWrapper.Search(2))
	s.True(s.treeWrapper.Search(3))
	s.False(s.treeWrapper.Search(99))
}

func (s *TreeWrapperTestSuite) TestRemove() {
	s.treeWrapper.Insert(1, 2)
	s.treeWrapper.Insert(1, 3)
	s.treeWrapper.Insert(2, 4)

	removed := s.treeWrapper.Remove(2)
	s.True(removed)
	s.Equal(2, s.treeWrapper.Size())
	s.False(s.treeWrapper.Search(2))
	s.False(s.treeWrapper.Search(4))
}

func (s *TreeWrapperTestSuite) TestTraversals() {
	s.treeWrapper.Insert(1, 2)
	s.treeWrapper.Insert(1, 3)
	s.treeWrapper.Insert(2, 4)
	s.treeWrapper.Insert(2, 5)
	s.treeWrapper.Insert(3, 6)

	preOrder := s.treeWrapper.PreOrder()
	s.Equal([]int{1, 2, 4, 5, 3, 6}, preOrder)

	postOrder := s.treeWrapper.PostOrder()
	s.Equal([]int{4, 5, 2, 6, 3, 1}, postOrder)

	levelOrder := s.treeWrapper.LevelOrder()
	s.Equal([]int{1, 2, 3, 4, 5, 6}, levelOrder)
}

func (s *TreeWrapperTestSuite) TestHeight() {
	s.Equal(1, s.treeWrapper.Height())

	s.treeWrapper.Insert(1, 2)
	s.treeWrapper.Insert(1, 3)
	s.Equal(2, s.treeWrapper.Height())

	s.treeWrapper.Insert(2, 4)
	s.Equal(3, s.treeWrapper.Height())
}

func (s *TreeWrapperTestSuite) TestIsEmpty() {
	s.False(s.treeWrapper.IsEmpty())

	s.treeWrapper.Clear()
	s.True(s.treeWrapper.IsEmpty())
}

func (s *TreeWrapperTestSuite) TestClear() {
	s.treeWrapper.Insert(1, 2)
	s.treeWrapper.Insert(1, 3)
	s.treeWrapper.Insert(2, 4)

	s.treeWrapper.Clear()
	s.Equal(0, s.treeWrapper.Size())
	s.True(s.treeWrapper.IsEmpty())
}

func (s *TreeWrapperTestSuite) TestStringWrapper() {
	strWrapper := NewTreeWrapper[string]("root")

	err := strWrapper.Insert("root", "child1")
	s.NoError(err)
	err = strWrapper.Insert("root", "child2")
	s.NoError(err)

	s.Equal(3, strWrapper.Size())
	s.True(strWrapper.Search("child1"))
	s.True(strWrapper.Search("child2"))

	levelOrder := strWrapper.LevelOrder()
	s.Equal([]string{"root", "child1", "child2"}, levelOrder)
}
