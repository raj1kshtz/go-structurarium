package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GenericQueueTestSuite struct {
	suite.Suite
	intQueue    *GenericQueue[int]
	stringQueue *GenericQueue[string]
}

func TestGenericQueueTestSuite(t *testing.T) {
	suite.Run(t, new(GenericQueueTestSuite))
}

func (s *GenericQueueTestSuite) SetupTest() {
	s.intQueue = NewGenericQueue[int]()
	s.stringQueue = NewGenericQueue[string]()
}

func (s *GenericQueueTestSuite) TestIntegerQueue() {
	assert.NoError(s.T(), s.intQueue.Offer(10))
	assert.NoError(s.T(), s.intQueue.AddFirst(20))
	assert.NoError(s.T(), s.intQueue.AddLast(30))
	assert.NoError(s.T(), s.intQueue.AddLast(40))

	value, err := s.intQueue.Peek()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 20, value)

	value, err = s.intQueue.Poll()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 20, value)

	value, err = s.intQueue.PollFirst()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 10, value)

	value, err = s.intQueue.PollLast()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 40, value)

	value = s.intQueue.Size()
	assert.Equal(s.T(), 1, value)

	_, err = s.intQueue.Poll()
	assert.NoError(s.T(), err)
}

func (s *GenericQueueTestSuite) TestPeekEmptyQueue() {
	_, err := s.intQueue.Peek()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestPollFirstEmptyQueue() {
	_, err := s.intQueue.PollFirst()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestPollLastEmptyQueue() {
	_, err := s.intQueue.PollLast()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestDisplay() {
	assert.NoError(s.T(), s.intQueue.Offer(1))
	assert.NoError(s.T(), s.intQueue.Offer(2))
	assert.NoError(s.T(), s.intQueue.Offer(3))

	expectedOutput := "[1 2 3]"
	assert.Equal(s.T(), expectedOutput, s.intQueue.display())
}

func (s *GenericQueueTestSuite) TestDisplayEmptyQueue() {
	expectedOutput := "[]"
	assert.Equal(s.T(), expectedOutput, s.intQueue.display())
}

func (s *GenericQueueTestSuite) TestHelperFunctionReplyChanReceive() {
	replyChan := make(chan interface{})

	go func() { replyChan <- nil }()
	err := replyChanReceive(replyChan)
	assert.NoError(s.T(), err)

	go func() { replyChan <- fmt.Errorf("test error") }()
	err = replyChanReceive(replyChan)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "test error", err.Error())
}

func (s *GenericQueueTestSuite) TestClear() {
	assert.NoError(s.T(), s.intQueue.Offer(5))
	assert.NoError(s.T(), s.intQueue.Offer(10))
	assert.NoError(s.T(), s.intQueue.AddFirst(10))
	assert.NoError(s.T(), s.intQueue.AddLast(10))

	_, err := s.intQueue.PollFirst()
	assert.NoError(s.T(), err)

	_, err = s.intQueue.PollLast()
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), 2, s.intQueue.Size())

	assert.NoError(s.T(), s.intQueue.Clear())

	_, err = s.intQueue.Peek()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())

	_, err = s.intQueue.PollLast()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestSizeEmptyQueue() {
	assert.Equal(s.T(), 0, s.intQueue.Size())
}

func (s *GenericQueueTestSuite) TestAddFirstOnEmptyQueue() {
	assert.NoError(s.T(), s.intQueue.AddFirst(100))
	assert.Equal(s.T(), 1, s.intQueue.Size())
}

func (s *GenericQueueTestSuite) TestAddLastOnEmptyQueue() {
	assert.NoError(s.T(), s.intQueue.AddLast(200))
	assert.Equal(s.T(), 1, s.intQueue.Size())
}

func (s *GenericQueueTestSuite) TestPollEmptyQueue() {
	_, err := s.intQueue.Poll()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestIsEmpty() {
	assert.Equal(s.T(), true, s.intQueue.IsEmpty())
}
