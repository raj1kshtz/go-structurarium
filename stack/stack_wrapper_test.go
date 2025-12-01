package stack

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WrapperStackTestSuite struct {
	suite.Suite
	stackWrapper *WrapperStack[int]
}

func TestWrapperStackTestSuite(t *testing.T) {
	suite.Run(t, new(WrapperStackTestSuite))
}

func (s *WrapperStackTestSuite) SetupTest() {
	s.stackWrapper = NewWrapperStack[int]()
}

func (s *WrapperStackTestSuite) TestWrapperStack() {
	s.Run("TestPush", func() {
		err := s.stackWrapper.Push(10)
		s.NoError(err)
		err = s.stackWrapper.Push(20)
		s.NoError(err)
		s.Equal(2, s.stackWrapper.Size())
	})

	s.Run("TestPop", func() {
		s.stackWrapper = NewWrapperStack[int]()
		s.stackWrapper.Push(10)
		s.stackWrapper.Push(20)

		value, err := s.stackWrapper.Pop()
		s.NoError(err)
		s.Equal(20, value)

		value, err = s.stackWrapper.Pop()
		s.NoError(err)
		s.Equal(10, value)

		_, err = s.stackWrapper.Pop()
		s.Error(err)
	})

	s.Run("TestPeek", func() {
		s.stackWrapper = NewWrapperStack[int]()
		s.stackWrapper.Push(10)
		s.stackWrapper.Push(20)

		value, err := s.stackWrapper.Peek()
		s.NoError(err)
		s.Equal(20, value)

		// Verify peek doesn't remove element
		s.Equal(2, s.stackWrapper.Size())
	})

	s.Run("TestClear", func() {
		s.stackWrapper = NewWrapperStack[int]()
		s.stackWrapper.Push(10)
		s.stackWrapper.Push(20)

		err := s.stackWrapper.Clear()
		s.NoError(err)
		s.True(s.stackWrapper.IsEmpty())
	})

	s.Run("TestIsEmpty", func() {
		s.stackWrapper = NewWrapperStack[int]()
		s.True(s.stackWrapper.IsEmpty())

		s.stackWrapper.Push(10)
		s.False(s.stackWrapper.IsEmpty())
	})

	s.Run("TestSize", func() {
		s.stackWrapper = NewWrapperStack[int]()
		s.Equal(0, s.stackWrapper.Size())

		s.stackWrapper.Push(10)
		s.Equal(1, s.stackWrapper.Size())

		s.stackWrapper.Push(20)
		s.Equal(2, s.stackWrapper.Size())

		s.stackWrapper.Pop()
		s.Equal(1, s.stackWrapper.Size())
	})
}
