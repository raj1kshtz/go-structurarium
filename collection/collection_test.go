package collection

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type GenericCollectionTestSuite struct {
	suite.Suite
	collection *GenericCollection[int]
}

func TestGenericCollectionTestSuite(t *testing.T) {
	suite.Run(t, new(GenericCollectionTestSuite))
}

func (s *GenericCollectionTestSuite) SetupTest() {
}

func (s *GenericCollectionTestSuite) TestAll() {
	s.Run("TestInitialCapacity", func() {
		collectionWithCapacity := NewGenericCollection[int](0)
		s.Equal(0, collectionWithCapacity.size(), "Expected initial size to be 0")
	})

	s.Run("TestAdd", func() {
		s.collection = NewGenericCollection[int](1)
		s.collection.add(10)
		s.True(s.collection.contains(10), "Expected collection to contain 10 after adding")
	})

	s.Run("TestRemove", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.add(10)
		removed := s.collection.remove(10)
		s.True(removed, "Expected removal of 10 to be successful")
		s.False(s.collection.contains(10), "Expected collection to not contain 10 after removal")
	})

	s.Run("TestRemoveNotFound", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.add(20)
		removed := s.collection.remove(10)
		s.False(removed, "Expected removal of 10 to be unsuccessful")
	})

	s.Run("TestSize", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.add(10)
		s.Equal(1, s.collection.size(), "Expected size to be 1")
	})

	s.Run("TestContains", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.add(10)
		s.True(s.collection.contains(10), "Expected collection to contain 10")
	})

	s.Run("TestIsEmpty", func() {
		s.collection = NewGenericCollection[int]()
		s.True(s.collection.isEmpty(), "Expected collection to be empty")
		s.collection.add(10)
		s.False(s.collection.isEmpty(), "Expected collection to not be empty after adding an element")
	})

	s.Run("TestAddAll", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.addAll([]int{1, 2, 3})
		s.Equal(3, s.collection.size(), "Expected size to be 3 after addAll")
		s.True(s.collection.contains(1), "Expected collection to contain 1 after addAll")
		s.True(s.collection.contains(2), "Expected collection to contain 2 after addAll")
		s.True(s.collection.contains(3), "Expected collection to contain 3 after addAll")
	})

	s.Run("TestRemoveAll", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.addAll([]int{1, 2, 3, 4})
		removed := s.collection.removeAll([]int{2, 3})
		s.True(removed, "Expected removeAll to successfully remove elements")
		s.Equal(2, s.collection.size(), "Expected size to be 2 after removeAll")
		s.False(s.collection.contains(2), "Expected collection to not contain 2 after removeAll")
		s.False(s.collection.contains(3), "Expected collection to not contain 3 after removeAll")
	})

	s.Run("TestRetainAll", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.addAll([]int{1, 2, 3, 4})
		s.collection.retainAll([]int{2, 3})
		s.Equal(2, s.collection.size(), "Expected size to be 2 after retainAll")
		s.True(s.collection.contains(2), "Expected collection to contain 2 after retainAll")
		s.True(s.collection.contains(3), "Expected collection to contain 3 after retainAll")
		s.False(s.collection.contains(1), "Expected collection to not contain 1 after retainAll")
		s.False(s.collection.contains(4), "Expected collection to not contain 4 after retainAll")
	})

	s.Run("TestContainsAll", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.addAll([]int{1, 2, 3})
		s.True(s.collection.containsAll([]int{1, 2}), "Expected to contain all elements")
		s.False(s.collection.containsAll([]int{1, 4}), "Expected not to contain all elements")
	})

	s.Run("TestClear", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.addAll([]int{1, 2, 3})
		s.collection.clear()
		s.Equal(0, s.collection.size(), "Expected size to be 0 after clear")
	})

	s.Run("TestToArray", func() {
		s.collection = NewGenericCollection[int]()
		s.collection.add(1)
		s.collection.add(2)
		arr := s.collection.toArray()
		expected := []int{1, 2}
		s.Equal(expected, arr, "Expected toArray to return correct elements")
	})
}
