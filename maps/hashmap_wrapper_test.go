package maps

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GenericHashMapWrapperTestSuite struct {
	suite.Suite
	mapWrapper *GenericHashMapWrapper[string, int]
}

func TestGenericHashMapWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(GenericHashMapWrapperTestSuite))
}

func (s *GenericHashMapWrapperTestSuite) SetupTest() {
	s.mapWrapper = NewGenericHashMapWrapper[string, int]()
}

func (s *GenericHashMapWrapperTestSuite) TestWrapper() {
	s.Run("TestPut", func() {
		isNew := s.mapWrapper.Put("key1", 100)
		s.True(isNew)

		isNew = s.mapWrapper.Put("key1", 200)
		s.False(isNew)

		value, exists := s.mapWrapper.Get("key1")
		s.True(exists)
		s.Equal(200, value)
	})

	s.Run("TestGet", func() {
		s.mapWrapper.Put("key1", 100)
		s.mapWrapper.Put("key2", 200)

		value, exists := s.mapWrapper.Get("key1")
		s.True(exists)
		s.Equal(100, value)

		value, exists = s.mapWrapper.Get("key3")
		s.False(exists)
		s.Equal(0, value)
	})

	s.Run("TestRemove", func() {
		s.mapWrapper = NewGenericHashMapWrapper[string, int]()
		s.mapWrapper.Put("key1", 100)
		s.mapWrapper.Put("key2", 200)

		removed := s.mapWrapper.Remove("key1")
		s.True(removed)

		_, exists := s.mapWrapper.Get("key1")
		s.False(exists)
	})

	s.Run("TestContainsKey", func() {
		s.mapWrapper = NewGenericHashMapWrapper[string, int]()
		s.mapWrapper.Put("key1", 100)

		s.True(s.mapWrapper.ContainsKey("key1"))
		s.False(s.mapWrapper.ContainsKey("key2"))
	})

	s.Run("TestSize", func() {
		s.mapWrapper = NewGenericHashMapWrapper[string, int]()
		s.Equal(0, s.mapWrapper.Size())

		s.mapWrapper.Put("key1", 100)
		s.Equal(1, s.mapWrapper.Size())

		s.mapWrapper.Put("key2", 200)
		s.Equal(2, s.mapWrapper.Size())
	})

	s.Run("TestIsEmpty", func() {
		s.mapWrapper = NewGenericHashMapWrapper[string, int]()
		s.True(s.mapWrapper.IsEmpty())

		s.mapWrapper.Put("key1", 100)
		s.False(s.mapWrapper.IsEmpty())
	})

	s.Run("TestClear", func() {
		s.mapWrapper = NewGenericHashMapWrapper[string, int]()
		s.mapWrapper.Put("key1", 100)
		s.mapWrapper.Put("key2", 200)

		s.mapWrapper.Clear()
		s.Equal(0, s.mapWrapper.Size())
		s.True(s.mapWrapper.IsEmpty())
	})

	s.Run("TestKeys", func() {
		s.mapWrapper = NewGenericHashMapWrapper[string, int]()
		s.mapWrapper.Put("key1", 100)
		s.mapWrapper.Put("key2", 200)
		s.mapWrapper.Put("key3", 300)

		keys := s.mapWrapper.Keys()
		s.Equal(3, len(keys))
		s.Contains(keys, "key1")
		s.Contains(keys, "key2")
		s.Contains(keys, "key3")
	})

	s.Run("TestValues", func() {
		s.mapWrapper = NewGenericHashMapWrapper[string, int]()
		s.mapWrapper.Put("key1", 100)
		s.mapWrapper.Put("key2", 200)
		s.mapWrapper.Put("key3", 300)

		values := s.mapWrapper.Values()
		s.Equal(3, len(values))
		s.Contains(values, 100)
		s.Contains(values, 200)
		s.Contains(values, 300)
	})
}
