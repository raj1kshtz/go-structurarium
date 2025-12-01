package queue

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GenericQueueWrapperTestSuite struct {
	suite.Suite
	queueWrapper *GenericQueueWrapper[int]
}

func TestGenericQueueWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(GenericQueueWrapperTestSuite))
}

func (s *GenericQueueWrapperTestSuite) SetupTest() {
	s.queueWrapper = NewGenericQueueWrapper[int]()
}

func (s *GenericQueueWrapperTestSuite) TestQueueWrapper() {
	s.Run("TestEnqueue", func() {
		success := s.queueWrapper.Enqueue(10)
		s.True(success)
		success = s.queueWrapper.Enqueue(20)
		s.True(success)
		s.Equal(2, s.queueWrapper.Size())
	})

	s.Run("TestDequeue", func() {
		s.queueWrapper = NewGenericQueueWrapper[int]()
		s.queueWrapper.Enqueue(10)
		s.queueWrapper.Enqueue(20)

		value, success := s.queueWrapper.Dequeue()
		s.True(success)
		s.Equal(10, value)

		value, success = s.queueWrapper.Dequeue()
		s.True(success)
		s.Equal(20, value)

		_, success = s.queueWrapper.Dequeue()
		s.False(success)
	})

	s.Run("TestPeek", func() {
		s.queueWrapper = NewGenericQueueWrapper[int]()
		s.queueWrapper.Enqueue(10)
		s.queueWrapper.Enqueue(20)

		value := s.queueWrapper.Peek()
		s.Equal(10, value)

		// Verify peek doesn't remove element
		s.Equal(2, s.queueWrapper.Size())
	})

	s.Run("TestSize", func() {
		s.queueWrapper = NewGenericQueueWrapper[int]()
		s.Equal(0, s.queueWrapper.Size())

		s.queueWrapper.Enqueue(10)
		s.Equal(1, s.queueWrapper.Size())

		s.queueWrapper.Enqueue(20)
		s.Equal(2, s.queueWrapper.Size())
	})

	s.Run("TestIsEmpty", func() {
		s.queueWrapper = NewGenericQueueWrapper[int]()
		s.True(s.queueWrapper.IsEmpty())

		s.queueWrapper.Enqueue(10)
		s.False(s.queueWrapper.IsEmpty())
	})

	s.Run("TestClear", func() {
		s.queueWrapper = NewGenericQueueWrapper[int]()
		s.queueWrapper.Enqueue(10)
		s.queueWrapper.Enqueue(20)

		s.queueWrapper.Clear()
		s.True(s.queueWrapper.IsEmpty())
	})

	s.Run("TestToArray", func() {
		s.queueWrapper = NewGenericQueueWrapper[int]()
		s.queueWrapper.Enqueue(10)
		s.queueWrapper.Enqueue(20)
		s.queueWrapper.Enqueue(30)

		arr := s.queueWrapper.ToArray()
		s.Equal([]int{10, 20, 30}, arr)
	})
}
