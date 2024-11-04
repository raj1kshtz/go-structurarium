package queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GenericQueueTestSuite struct {
	suite.Suite
	intQueue *GenericQueue[int]
}

func TestGenericQueueTestSuite(t *testing.T) {
	suite.Run(t, new(GenericQueueTestSuite))
}

func (t *GenericQueueTestSuite) SetupTest() {
	t.intQueue = NewGenericQueue[int]()
}

func (t *GenericQueueTestSuite) TestQueue() {
	t.Run("TestIsEmpty", func() {
		t.True(t.intQueue.isEmpty())
	})

	t.Run("TestEnqueue", func() {
		t.True(t.intQueue.enQueue(1))
		t.True(t.intQueue.enQueue(2))
	})

	t.Run("TestQueueToArray", func() {
		vals := t.intQueue.toArray()
		t.Equal([]int{1, 2}, vals)
	})

	t.Run("TestDequeue", func() {
		val, ok := t.intQueue.deQueue()
		t.True(ok)
		t.Equal(1, val)
	})

	t.Run("TestQueueSize", func() {
		t.Equal(1, t.intQueue.size())
	})

	t.Run("TestQueuePeek", func() {
		t.Equal(2, t.intQueue.peek())
	})

	t.Run("TestQueueClear", func() {
		t.intQueue.clear()
		t.Equal(0, t.intQueue.size())
	})
	
	t.Run("TestEmptyPeek", func() {
		t.Empty(t.intQueue.peek())
	})
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
