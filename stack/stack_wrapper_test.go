package stack

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StackWrapperTestSuite struct {
	suite.Suite
	intStack    *Wrapper[int]
	stringStack *Wrapper[string]
}

func TestStackWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(StackWrapperTestSuite))
}

func (s *StackWrapperTestSuite) SetupTest() {
	s.intStack = NewStackWrapper[int]()
	s.stringStack = NewStackWrapper[string]()
}

func (s *StackWrapperTestSuite) TestIntegerStack() {
	assert.NoError(s.T(), s.intStack.Push(10))
	assert.NoError(s.T(), s.intStack.Push(20))

	value, err := s.intStack.Pop()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 20, value)

	topValue, err := s.intStack.Top()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 10, topValue)

	_, err = s.intStack.Pop()
	assert.NoError(s.T(), err)
	_, err = s.intStack.Pop()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "stack underflow", err.Error())

	_, err = s.intStack.Top()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "stack is empty", err.Error())
}

func (s *StackWrapperTestSuite) TestStringStack() {
	assert.NoError(s.T(), s.stringStack.Push("Hello"))
	assert.NoError(s.T(), s.stringStack.Push("World"))

	value, err := s.stringStack.Pop()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "World", value)

	topValue, err := s.stringStack.Top()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "Hello", topValue)

	_, err = s.stringStack.Pop()
	assert.NoError(s.T(), err)
	_, err = s.stringStack.Pop()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "stack underflow", err.Error())

	_, err = s.stringStack.Top()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "stack is empty", err.Error())
}

func (s *StackWrapperTestSuite) TestDisplay() {
	assert.NoError(s.T(), s.intStack.Push(1))
	assert.NoError(s.T(), s.intStack.Push(2))
	assert.NoError(s.T(), s.intStack.Push(3))

	s.intStack.Display()
}

func (s *StackWrapperTestSuite) TestClear() {
	assert.NoError(s.T(), s.intStack.Push(1))
	assert.NoError(s.T(), s.intStack.Push(2))
	assert.NoError(s.T(), s.intStack.Push(3))
	assert.NoError(s.T(), s.intStack.Clear())
}
