package stack

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StackWrapperTestSuite struct {
	suite.Suite
	intStack    *WrapperStack[int]
	stringStack *WrapperStack[string]
}

func TestStackWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(StackWrapperTestSuite))
}

func (s *StackWrapperTestSuite) SetupTest() {
	s.intStack = NewWrapperStack[int]()
	s.stringStack = NewWrapperStack[string]()
}

func (s *StackWrapperTestSuite) TestIntegerStack() {
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

func (s *StackWrapperTestSuite) TestClear() {
	assert.NoError(s.T(), s.intStack.Push(1))
	assert.NoError(s.T(), s.intStack.Push(2))
	assert.NoError(s.T(), s.intStack.Push(3))
	assert.NoError(s.T(), s.intStack.Clear())
}
