package vector

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type GenericVectorTestSuite struct {
	suite.Suite
	intVector *GenericVector[int]
}

func TestGenericVectorTestSuite(t *testing.T) {
	suite.Run(t, new(GenericVectorTestSuite))
}

func (s *GenericVectorTestSuite) SetupTest() {
	s.intVector = NewGenericVector[int](5)
}

func (s *GenericVectorTestSuite) TestVector() {
	s.Run("TestAdd", func() {
		err := s.intVector.add(10)
		s.NoError(err)
		err = s.intVector.add(20)
		s.NoError(err)
		s.Equal(2, s.intVector.size())
	})

	s.Run("TestAddAt", func() {
		s.intVector = NewGenericVector[int](15)
		_ = s.intVector.add(10)
		_ = s.intVector.add(20)

		err := s.intVector.addAt(1, 15)
		s.NoError(err)

		value, err := s.intVector.get(1)
		s.NoError(err)
		s.Equal(15, value)

		err = s.intVector.addAt(10, 25)
		s.Error(err)
	})

	s.Run("TestRemoveAt", func() {
		s.intVector = NewGenericVector[int](5)
		_ = s.intVector.add(10)
		_ = s.intVector.add(20)
		_ = s.intVector.add(30)

		err := s.intVector.removeAt(1)
		s.NoError(err)

		s.Equal(2, s.intVector.size())

		value, err := s.intVector.get(1)
		s.NoError(err)
		s.Equal(30, value)

		err = s.intVector.removeAt(5)
		s.Error(err)
	})

	s.Run("TestGet", func() {
		s.intVector = NewGenericVector[int](5)
		_ = s.intVector.add(10)
		_ = s.intVector.add(20)

		value, err := s.intVector.get(1)
		s.NoError(err)
		s.Equal(20, value)

		_, err = s.intVector.get(5)
		s.Error(err)
	})

	s.Run("TestSet", func() {
		_ = s.intVector.add(10)
		_ = s.intVector.add(20)

		err := s.intVector.set(1, 25)
		s.NoError(err)

		value, err := s.intVector.get(1)
		s.NoError(err)
		s.Equal(25, value)

		err = s.intVector.set(5, 30)
		s.Error(err)
	})

	s.Run("TestSize", func() {
		s.intVector = NewGenericVector[int](5)
		_ = s.intVector.add(10)
		_ = s.intVector.add(20)
		s.Equal(2, s.intVector.size())

		_ = s.intVector.add(30)
		s.Equal(3, s.intVector.size())
	})

	s.Run("TestIsEmpty", func() {
		s.intVector = NewGenericVector[int]()
		s.True(s.intVector.isEmpty())

		_ = s.intVector.add(10)
		s.False(s.intVector.isEmpty())
	})

	s.Run("TestClear", func() {
		_ = s.intVector.add(10)
		_ = s.intVector.add(20)

		err := s.intVector.clear()
		s.NoError(err)
		s.Equal(0, s.intVector.size())
		s.True(s.intVector.isEmpty())
	})

	s.Run("TestEnsureCapacity", func() {
		s.intVector.ensureCapacity(10) // No error expected; this just ensures that it doesn't panic
	})

	s.Run("TestTrimToSize", func() {
		s.intVector = NewGenericVector[int](3)
		_ = s.intVector.Add(10)
		_ = s.intVector.Add(20)
		_ = s.intVector.Add(30)

		s.Equal(3, s.intVector.Size())
		err := s.intVector.trimToSize()
		s.NoError(err)
		s.Equal(3, cap(s.intVector.data))
	})

	s.Run("TestToArray", func() {
		s.intVector = NewGenericVector[int](3)
		_ = s.intVector.Add(10)
		_ = s.intVector.Add(20)
		_ = s.intVector.Add(30)

		arr := s.intVector.toArray()
		s.Equal([]int{10, 20, 30}, arr)
	})

}
