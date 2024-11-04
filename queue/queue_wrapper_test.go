package queue

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type GenericQueueWrapperTestSuite struct {
	suite.Suite
	intQueue *GenericQueueWrapper[int]
}

func TestGenericWrapperQueueSuite(t *testing.T) {
	suite.Run(t, new(GenericQueueWrapperTestSuite))
}

func (t *GenericQueueWrapperTestSuite) SetupTest() {
	t.intQueue = NewGenericQueueWrapper[int]()
}

func (t *GenericQueueWrapperTestSuite) TestQueueWrapper() {
	t.Run("TestIsEmpty", func() {
		t.True(t.intQueue.IsEmpty())
	})

	t.Run("TestEnqueue", func() {
		t.True(t.intQueue.Enqueue(1))
		t.True(t.intQueue.Enqueue(2))
	})

	t.Run("TestQueueToArray", func() {
		vals := t.intQueue.ToArray()
		t.Equal([]int{1, 2}, vals)
	})

	t.Run("TestDequeue", func() {
		val, ok := t.intQueue.Dequeue()
		t.True(ok)
		t.Equal(1, val)
	})

	t.Run("TestQueueSize", func() {
		t.Equal(1, t.intQueue.Size())
	})

	t.Run("TestQueuePeek", func() {
		val := t.intQueue.Peek()
		t.Equal(2, val)
	})

	t.Run("TestQueueClear", func() {
		t.intQueue.Clear()
		t.Equal(0, t.intQueue.Size())
	})
}
