package vector

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WrapperVectorTestSuite struct {
	suite.Suite
	vectorWrapper *WrapperVector[int]
}

func TestWrapperVectorTestSuite(t *testing.T) {
	suite.Run(t, new(WrapperVectorTestSuite))
}

func (s *WrapperVectorTestSuite) SetupTest() {
	s.vectorWrapper = NewWrapperVector[int]()
}

func (s *WrapperVectorTestSuite) TestVectorWrapper() {
	s.Run("TestAdd", func() {
		err := s.vectorWrapper.Add(10)
		s.NoError(err)
		err = s.vectorWrapper.Add(20)
		s.NoError(err)
		s.Equal(2, s.vectorWrapper.Size())
	})

	s.Run("TestAddAt", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.vectorWrapper.Add(10)
		s.vectorWrapper.Add(30)

		err := s.vectorWrapper.AddAt(1, 20)
		s.NoError(err)

		value, err := s.vectorWrapper.Get(1)
		s.NoError(err)
		s.Equal(20, value)
	})

	s.Run("TestGet", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.vectorWrapper.Add(10)
		s.vectorWrapper.Add(20)

		value, err := s.vectorWrapper.Get(0)
		s.NoError(err)
		s.Equal(10, value)

		value, err = s.vectorWrapper.Get(1)
		s.NoError(err)
		s.Equal(20, value)

		_, err = s.vectorWrapper.Get(5)
		s.Error(err)
	})

	s.Run("TestSet", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.vectorWrapper.Add(10)
		s.vectorWrapper.Add(20)

		err := s.vectorWrapper.Set(1, 30)
		s.NoError(err)

		value, err := s.vectorWrapper.Get(1)
		s.NoError(err)
		s.Equal(30, value)
	})

	s.Run("TestRemoveAt", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.vectorWrapper.Add(10)
		s.vectorWrapper.Add(20)
		s.vectorWrapper.Add(30)

		err := s.vectorWrapper.RemoveAt(1)
		s.NoError(err)
		s.Equal(2, s.vectorWrapper.Size())

		value, err := s.vectorWrapper.Get(1)
		s.NoError(err)
		s.Equal(30, value)
	})

	s.Run("TestSize", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.Equal(0, s.vectorWrapper.Size())

		s.vectorWrapper.Add(10)
		s.Equal(1, s.vectorWrapper.Size())
	})

	s.Run("TestIsEmpty", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.True(s.vectorWrapper.IsEmpty())

		s.vectorWrapper.Add(10)
		s.False(s.vectorWrapper.IsEmpty())
	})

	s.Run("TestClear", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.vectorWrapper.Add(10)
		s.vectorWrapper.Add(20)

		err := s.vectorWrapper.Clear()
		s.NoError(err)
		s.True(s.vectorWrapper.IsEmpty())
	})

	s.Run("TestToArray", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.vectorWrapper.Add(10)
		s.vectorWrapper.Add(20)
		s.vectorWrapper.Add(30)

		arr := s.vectorWrapper.ToArray()
		s.Equal([]int{10, 20, 30}, arr)
	})

	s.Run("TestEnsureCapacity", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.vectorWrapper.EnsureCapacity(10)
		// Just ensuring no panic
	})

	s.Run("TestTrimToSize", func() {
		s.vectorWrapper = NewWrapperVector[int]()
		s.vectorWrapper.Add(10)
		s.vectorWrapper.Add(20)

		err := s.vectorWrapper.TrimToSize()
		s.NoError(err)
	})
}
