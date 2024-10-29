package queue

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type QueueWrapperTestSuite struct {
	suite.Suite
	intQueue    *WrapperQueue[int]
	stringQueue *WrapperQueue[string]
}

func TestQueueWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(QueueWrapperTestSuite))
}

func (s *QueueWrapperTestSuite) SetupTest() {
	s.intQueue = NewWrapperQueue[int]()
	s.stringQueue = NewWrapperQueue[string]()
}

func (s *QueueWrapperTestSuite) TestWrapperQueue() {

	s.Run("TestIsEmpty", func() {
		assert.True(s.T(), s.intQueue.IsEmpty(), "Queue should be empty initially")
	})

	s.Run("TestAddFirstAndPollFirst", func() {
		err := s.intQueue.AddFirst(1)
		assert.NoError(s.T(), err)
		assert.False(s.T(), s.intQueue.IsEmpty(), "Queue should not be empty after adding an element")
		value, err := s.intQueue.PollFirst()
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), 1, value)
	})

	s.Run("TestAddLastAndPoll", func() {
		err := s.intQueue.AddLast(2)
		assert.NoError(s.T(), err)
		value, err := s.intQueue.Poll()
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), 2, value)
	})

	s.Run("TestOffer, AddFirst, Peek, Display, Clear", func() {
		err := s.intQueue.Offer(3)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), 1, s.intQueue.Size())

		err = s.intQueue.AddFirst(4)
		assert.NoError(s.T(), err)
		peekedValue, err := s.intQueue.Peek()
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), 4, peekedValue)

		s.intQueue.Display()
		assert.Equal(s.T(), 2, s.intQueue.Size())

		err = s.intQueue.Clear()
		assert.NoError(s.T(), err)
		assert.True(s.T(), s.intQueue.IsEmpty(), "Queue should be empty after clearing")
	})

	s.Run("TestStringQueue", func() {
		err := s.stringQueue.AddFirst("Hello")
		assert.NoError(s.T(), err)

		err = s.stringQueue.AddLast("World")
		assert.NoError(s.T(), err)

		peekedStr, err := s.stringQueue.Peek()
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "Hello", peekedStr)

		firstStr, err := s.stringQueue.PollFirst()
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "Hello", firstStr)

		lastStr, err := s.stringQueue.PollLast()
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "World", lastStr)

		assert.True(s.T(), s.stringQueue.IsEmpty(), "String queue should be empty after polling all elements")
	})

}
