package collection

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type GenericCollectionWrapperTestSuite struct {
	suite.Suite
	wrapper *GenericCollectionWrapper[int]
}

func TestGenericCollectionWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(GenericCollectionWrapperTestSuite))
}

func (s *GenericCollectionWrapperTestSuite) SetupTest() {
}

func (s *GenericCollectionWrapperTestSuite) TestAll() {
	s.Run("TestAdd", func() {
		s.wrapper = NewGenericCollectionWrapper[int]()
		s.wrapper.Add(10)
		s.True(s.wrapper.Contains(10), "Expected collection to contain 10 after adding")
	})

	s.Run("TestSize", func() {
		s.wrapper = NewGenericCollectionWrapper[int]()
		s.wrapper.Add(10)
		s.Equal(1, s.wrapper.Size(), "Expected size to be 1")
	})

	s.Run("TestContains", func() {
		s.wrapper = NewGenericCollectionWrapper[int]()
		s.wrapper.Add(10)
		s.True(s.wrapper.Contains(10), "Expected collection to contain 10")
	})

	s.Run("TestRemove", func() {
		s.wrapper = NewGenericCollectionWrapper[int]()
		s.wrapper.Add(10)
		s.wrapper.Remove(10)
		s.False(s.wrapper.Contains(10), "Expected collection to not contain 10 after removal")
	})

	s.Run("TestIsEmpty", func() {
		s.wrapper = NewGenericCollectionWrapper[int]()
		s.True(s.wrapper.IsEmpty(), "Expected collection to be empty")
		s.wrapper.Add(10)
		s.False(s.wrapper.IsEmpty(), "Expected collection to not be empty after adding an element")
	})

	s.Run("TestAddAll", func() {
		s.wrapper = NewGenericCollectionWrapper[int]()
		s.wrapper.AddAll([]int{1, 2, 3})
		s.Equal(3, s.wrapper.Size(), "Expected size to be 3 after addAll")
		s.True(s.wrapper.Contains(1), "Expected collection to contain 1 after addAll")
		s.True(s.wrapper.Contains(2), "Expected collection to contain 2 after addAll")
		s.True(s.wrapper.Contains(3), "Expected collection to contain 3 after addAll")
	})

	s.Run("TestClear", func() {
		s.wrapper = NewGenericCollectionWrapper[int]()
		s.wrapper.AddAll([]int{1, 2, 3})
		s.wrapper.Clear()
		s.Equal(0, s.wrapper.Size(), "Expected size to be 0 after clear")
	})

	s.Run("TestToArray", func() {
		s.wrapper = NewGenericCollectionWrapper[int]()
		s.wrapper.Add(1)
		s.wrapper.Add(2)
		arr := s.wrapper.ToArray()
		expected := []int{1, 2}
		s.Equal(expected, arr, "Expected toArray to return correct elements")
	})
}
