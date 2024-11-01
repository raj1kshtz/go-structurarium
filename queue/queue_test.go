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
	assert.NoError(s.T(), s.intQueue.offer(10))
	assert.NoError(s.T(), s.intQueue.addFirst(20))
	assert.NoError(s.T(), s.intQueue.addLast(30))
	assert.NoError(s.T(), s.intQueue.addLast(40))

	value, err := s.intQueue.peek()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 20, value)

	value, err = s.intQueue.poll()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 20, value)

	value, err = s.intQueue.pollFirst()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 10, value)

	value, err = s.intQueue.pollLast()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 40, value)

	value = s.intQueue.size()
	assert.Equal(s.T(), 1, value)

	_, err = s.intQueue.poll()
	assert.NoError(s.T(), err)
}

func (s *GenericQueueTestSuite) TestPeekEmptyQueue() {
	_, err := s.intQueue.peek()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestPollFirstEmptyQueue() {
	_, err := s.intQueue.pollFirst()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestPollLastEmptyQueue() {
	_, err := s.intQueue.pollLast()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestDisplay() {
	assert.NoError(s.T(), s.intQueue.offer(1))
	assert.NoError(s.T(), s.intQueue.offer(2))
	assert.NoError(s.T(), s.intQueue.offer(3))

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
	assert.NoError(s.T(), s.intQueue.offer(5))
	assert.NoError(s.T(), s.intQueue.offer(10))
	assert.NoError(s.T(), s.intQueue.addFirst(10))
	assert.NoError(s.T(), s.intQueue.addLast(10))

	_, err := s.intQueue.pollFirst()
	assert.NoError(s.T(), err)

	_, err = s.intQueue.pollLast()
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), 2, s.intQueue.size())

	assert.NoError(s.T(), s.intQueue.clear())

	_, err = s.intQueue.peek()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())

	_, err = s.intQueue.pollLast()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestSizeEmptyQueue() {
	assert.Equal(s.T(), 0, s.intQueue.size())
}

func (s *GenericQueueTestSuite) TestAddFirstOnEmptyQueue() {
	assert.NoError(s.T(), s.intQueue.addFirst(100))
	assert.Equal(s.T(), 1, s.intQueue.size())
}

func (s *GenericQueueTestSuite) TestAddLastOnEmptyQueue() {
	assert.NoError(s.T(), s.intQueue.addLast(200))
	assert.Equal(s.T(), 1, s.intQueue.size())
}

func (s *GenericQueueTestSuite) TestPollEmptyQueue() {
	_, err := s.intQueue.poll()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "queue is empty", err.Error())
}

func (s *GenericQueueTestSuite) TestIsEmpty() {
	assert.Equal(s.T(), true, s.intQueue.isEmpty())
}
