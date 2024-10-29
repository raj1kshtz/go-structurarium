package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GenericStackTestSuite struct {
	suite.Suite
	intStack    *GenericStack[int]
	stringStack *GenericStack[string]
}

func TestGenericStackTestSuite(t *testing.T) {
	suite.Run(t, new(GenericStackTestSuite))
}

func (s *GenericStackTestSuite) SetupTest() {
	s.intStack = NewGenericStack[int]()       // Dynamic stack without maxSize
	s.stringStack = NewGenericStack[string]() // Dynamic stack without maxSize
}

func (s *GenericStackTestSuite) TestIntegerStack() {
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

func (s *GenericStackTestSuite) TestStringStack() {
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

func (s *GenericStackTestSuite) TestDisplay() {
	assert.NoError(s.T(), s.intStack.Push(1))
	assert.NoError(s.T(), s.intStack.Push(2))
	assert.NoError(s.T(), s.intStack.Push(3))

	expectedOutput := "[1 2 3]"
	assert.Equal(s.T(), expectedOutput, s.intStack.display())
}

func (s *GenericStackTestSuite) TestClear() {
	assert.NoError(s.T(), s.intStack.Push(5))
	assert.NoError(s.T(), s.intStack.Push(10))
	assert.NoError(s.T(), s.intStack.Clear())

	// Check that stack is empty after Clear
	_, err := s.intStack.Top()
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

func (s *GenericStackTestSuite) TestIsEmpty() {
	assert.Equal(s.T(), true, s.intStack.IsEmpty())
}
