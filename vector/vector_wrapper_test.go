package vector

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type WrapperVectorTestSuite struct {
	suite.Suite
	intWrapperVector *WrapperVector[int]
}

func TestWrapperVectorTestSuite(t *testing.T) {
	suite.Run(t, new(WrapperVectorTestSuite))
}

func (s *WrapperVectorTestSuite) SetupTest() {
	s.intWrapperVector = NewWrapperVector[int](2)
}

func (s *WrapperVectorTestSuite) TestWrapperVector() {
	s.Run("TestAdd", func() {
		err := s.intWrapperVector.Add(10)
		s.NoError(err)
		err = s.intWrapperVector.Add(20)
		s.NoError(err)
		s.Equal(2, s.intWrapperVector.Size())
	})

	s.Run("TestAddAt", func() {
		s.intWrapperVector = NewWrapperVector[int](3)
		_ = s.intWrapperVector.Add(10)
		_ = s.intWrapperVector.Add(20)

		err := s.intWrapperVector.AddAt(1, 15)
		s.NoError(err)

		value, err := s.intWrapperVector.Get(1)
		s.NoError(err)
		s.Equal(15, value)

		err = s.intWrapperVector.AddAt(10, 25)
		s.Error(err)
	})

	s.Run("TestRemoveAt", func() {
		s.intWrapperVector = NewWrapperVector[int](2)
		_ = s.intWrapperVector.Add(10)
		_ = s.intWrapperVector.Add(20)
		_ = s.intWrapperVector.Add(30)

		err := s.intWrapperVector.RemoveAt(1)
		s.NoError(err)

		s.Equal(2, s.intWrapperVector.Size())

		value, err := s.intWrapperVector.Get(1)
		s.NoError(err)
		s.Equal(30, value)

		err = s.intWrapperVector.RemoveAt(5)
		s.Error(err)
	})

	s.Run("TestGet", func() {
		s.intWrapperVector = NewWrapperVector[int](2)
		_ = s.intWrapperVector.Add(10)
		_ = s.intWrapperVector.Add(20)

		value, err := s.intWrapperVector.Get(1)
		s.NoError(err)
		s.Equal(20, value)

		_, err = s.intWrapperVector.Get(5)
		s.Error(err)
	})

	s.Run("TestSet", func() {
		s.intWrapperVector = NewWrapperVector[int](2)
		_ = s.intWrapperVector.Add(10)
		_ = s.intWrapperVector.Add(20)

		err := s.intWrapperVector.Set(1, 25)
		s.NoError(err)

		value, err := s.intWrapperVector.Get(1)
		s.NoError(err)
		s.Equal(25, value)

		err = s.intWrapperVector.Set(5, 30)
		s.Error(err)
	})

	s.Run("TestSize", func() {
		s.intWrapperVector = NewWrapperVector[int](2)
		_ = s.intWrapperVector.Add(10)
		_ = s.intWrapperVector.Add(20)
		s.Equal(2, s.intWrapperVector.Size())

		_ = s.intWrapperVector.Add(30)
		s.Equal(3, s.intWrapperVector.Size())
	})

	s.Run("TestIsEmpty", func() {
		s.intWrapperVector = NewWrapperVector[int](2)
		s.True(s.intWrapperVector.IsEmpty())

		_ = s.intWrapperVector.Add(10)
		s.False(s.intWrapperVector.IsEmpty())
	})

	s.Run("TestClear", func() {
		s.intWrapperVector = NewWrapperVector[int]()
		_ = s.intWrapperVector.Add(10)
		_ = s.intWrapperVector.Add(20)

		err := s.intWrapperVector.Clear()
		s.NoError(err)
		s.Equal(0, s.intWrapperVector.Size())
		s.True(s.intWrapperVector.IsEmpty())
	})

	s.Run("TestEnsureCapacity", func() {
		s.intWrapperVector.EnsureCapacity(10) // No error expected; this just ensures that it doesn't panic
	})

	s.Run("TestTrimToSize", func() {
		s.intWrapperVector = NewWrapperVector[int](3)
		_ = s.intWrapperVector.Add(10)
		_ = s.intWrapperVector.Add(20)
		_ = s.intWrapperVector.Add(30)

		s.Equal(3, s.intWrapperVector.Size())
		err := s.intWrapperVector.TrimToSize()
		s.NoError(err)
		s.Equal(3, cap(s.intWrapperVector.vector.data))
	})

	s.Run("TestToArray", func() {
		s.intWrapperVector = NewWrapperVector[int](3)
		_ = s.intWrapperVector.Add(10)
		_ = s.intWrapperVector.Add(20)
		_ = s.intWrapperVector.Add(30)

		arr := s.intWrapperVector.ToArray()
		s.Equal([]int{10, 20, 30}, arr)
	})
}
