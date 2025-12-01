package collection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GenericCollectionWrapperTestSuite struct {
	suite.Suite
	collectionWrapper *GenericCollectionWrapper[int]
}

func TestGenericCollectionWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(GenericCollectionWrapperTestSuite))
}

func (s *GenericCollectionWrapperTestSuite) SetupTest() {
	s.collectionWrapper = NewGenericCollectionWrapper[int]()
}

func (s *GenericCollectionWrapperTestSuite) TestCollectionWrapper() {
	s.Run("TestAdd", func() {
		success := s.collectionWrapper.Add(10)
		s.True(success)
		success = s.collectionWrapper.Add(20)
		s.True(success)
		s.Equal(2, s.collectionWrapper.Size())
	})

	s.Run("TestAddAll", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		success := s.collectionWrapper.AddAll([]int{10, 20, 30})
		s.True(success)
		s.Equal(3, s.collectionWrapper.Size())
	})

	s.Run("TestRemove", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		s.collectionWrapper.Add(10)
		s.collectionWrapper.Add(20)

		removed := s.collectionWrapper.Remove(10)
		s.True(removed)
		s.Equal(1, s.collectionWrapper.Size())

		removed = s.collectionWrapper.Remove(30)
		s.False(removed)
	})

	s.Run("TestRemoveAll", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		s.collectionWrapper.AddAll([]int{10, 20, 30, 40})

		removed := s.collectionWrapper.RemoveAll([]int{20, 30})
		s.True(removed)
		s.Equal(2, s.collectionWrapper.Size())
		s.True(s.collectionWrapper.Contains(10))
		s.False(s.collectionWrapper.Contains(20))
	})

	s.Run("TestRetainAll", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		s.collectionWrapper.AddAll([]int{10, 20, 30, 40})

		changed := s.collectionWrapper.RetainAll([]int{20, 30})
		s.True(changed)
		s.Equal(2, s.collectionWrapper.Size())
		s.True(s.collectionWrapper.Contains(20))
		s.True(s.collectionWrapper.Contains(30))
		s.False(s.collectionWrapper.Contains(10))
	})

	s.Run("TestContains", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		s.collectionWrapper.Add(10)

		s.True(s.collectionWrapper.Contains(10))
		s.False(s.collectionWrapper.Contains(20))
	})

	s.Run("TestContainsAll", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		s.collectionWrapper.AddAll([]int{10, 20, 30})

		s.True(s.collectionWrapper.ContainsAll([]int{10, 20}))
		s.False(s.collectionWrapper.ContainsAll([]int{10, 40}))
	})

	s.Run("TestSize", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		s.Equal(0, s.collectionWrapper.Size())

		s.collectionWrapper.Add(10)
		s.Equal(1, s.collectionWrapper.Size())
	})

	s.Run("TestIsEmpty", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		s.True(s.collectionWrapper.IsEmpty())

		s.collectionWrapper.Add(10)
		s.False(s.collectionWrapper.IsEmpty())
	})

	s.Run("TestClear", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		s.collectionWrapper.AddAll([]int{10, 20, 30})

		s.collectionWrapper.Clear()
		s.True(s.collectionWrapper.IsEmpty())
	})

	s.Run("TestToArray", func() {
		s.collectionWrapper = NewGenericCollectionWrapper[int]()
		s.collectionWrapper.AddAll([]int{10, 20, 30})

		arr := s.collectionWrapper.ToArray()
		s.Equal([]int{10, 20, 30}, arr)
	})
}
