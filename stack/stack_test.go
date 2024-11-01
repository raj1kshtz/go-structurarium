package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GenericStackTestSuite struct {
	suite.Suite
	intStack *GenericStack[int]
}

func TestGenericStackTestSuite(t *testing.T) {
	suite.Run(t, new(GenericStackTestSuite))
}

func (s *GenericStackTestSuite) SetupTest() {
	s.intStack = NewGenericStack[int]()
}

func (s *GenericStackTestSuite) TestIntegerStack() {
	assert.Equal(s.T(), true, s.intStack.IsEmpty())
	assert.NoError(s.T(), s.intStack.Push(10))
	assert.NoError(s.T(), s.intStack.Push(20))

	value, err := s.intStack.Pop()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 20, value)

	topValue, err := s.intStack.Peek()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 10, topValue)

	assert.Equal(s.T(), 1, s.intStack.Size())

	_, err = s.intStack.Pop()
	assert.NoError(s.T(), err)
	_, err = s.intStack.Pop()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "stack underflow", err.Error())

	_, err = s.intStack.Peek()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "stack is empty", err.Error())
}

func (s *GenericStackTestSuite) TestClear() {
	assert.NoError(s.T(), s.intStack.Push(5))
	assert.NoError(s.T(), s.intStack.Push(10))
	assert.NoError(s.T(), s.intStack.Clear())

	_, err := s.intStack.Peek()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "stack is empty", err.Error())

	_, err = s.intStack.Pop()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "stack underflow", err.Error())
}

func (s *GenericStackTestSuite) TestHelperFunctionReplyChanReceive() {
	replyChan := make(chan interface{})

	go func() { replyChan <- nil }()
	err := replyChanReceive(replyChan)
	assert.NoError(s.T(), err)

	go func() { replyChan <- fmt.Errorf("test error") }()
	err = replyChanReceive(replyChan)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "test error", err.Error())
}
