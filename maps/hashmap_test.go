package maps

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GenericHashMapTestSuite struct {
	suite.Suite
	intMap *GenericHashMap[string, int]
}

func TestGenericHashMapTestSuite(t *testing.T) {
	suite.Run(t, new(GenericHashMapTestSuite))
}

func (s *GenericHashMapTestSuite) SetupTest() {
	s.intMap = NewGenericHashMap[string, int]()
}

func (s *GenericHashMapTestSuite) TestHashMap() {
	s.Run("TestPut", func() {
		isNew := s.intMap.put("key1", 100)
		s.True(isNew, "Expected new key to be added")

		isNew = s.intMap.put("key1", 200)
		s.False(isNew, "Expected existing key to be updated")

		value, exists := s.intMap.get("key1")
		s.True(exists)
		s.Equal(200, value)
	})

	s.Run("TestGet", func() {
		s.intMap = NewGenericHashMap[string, int]()
		s.intMap.put("key1", 100)
		s.intMap.put("key2", 200)

		value, exists := s.intMap.get("key1")
		s.True(exists)
		s.Equal(100, value)

		value, exists = s.intMap.get("key3")
		s.False(exists)
		s.Equal(0, value)
	})

	s.Run("TestRemove", func() {
		s.intMap = NewGenericHashMap[string, int]()
		s.intMap.put("key1", 100)
		s.intMap.put("key2", 200)

		removed := s.intMap.remove("key1")
		s.True(removed)

		_, exists := s.intMap.get("key1")
		s.False(exists)

		removed = s.intMap.remove("key3")
		s.False(removed)
	})

	s.Run("TestContainsKey", func() {
		s.intMap = NewGenericHashMap[string, int]()
		s.intMap.put("key1", 100)

		s.True(s.intMap.containsKey("key1"))
		s.False(s.intMap.containsKey("key2"))
	})

	s.Run("TestSize", func() {
		s.intMap = NewGenericHashMap[string, int]()
		s.Equal(0, s.intMap.getSize())

		s.intMap.put("key1", 100)
		s.Equal(1, s.intMap.getSize())

		s.intMap.put("key2", 200)
		s.Equal(2, s.intMap.getSize())

		s.intMap.remove("key1")
		s.Equal(1, s.intMap.getSize())
	})

	s.Run("TestIsEmpty", func() {
		s.intMap = NewGenericHashMap[string, int]()
		s.True(s.intMap.isEmpty())

		s.intMap.put("key1", 100)
		s.False(s.intMap.isEmpty())

		s.intMap.remove("key1")
		s.True(s.intMap.isEmpty())
	})

	s.Run("TestClear", func() {
		s.intMap = NewGenericHashMap[string, int]()
		s.intMap.put("key1", 100)
		s.intMap.put("key2", 200)
		s.intMap.put("key3", 300)

		s.intMap.clear()
		s.Equal(0, s.intMap.getSize())
		s.True(s.intMap.isEmpty())
	})

	s.Run("TestKeys", func() {
		s.intMap = NewGenericHashMap[string, int]()
		s.intMap.put("key1", 100)
		s.intMap.put("key2", 200)
		s.intMap.put("key3", 300)

		keys := s.intMap.keys()
		s.Equal(3, len(keys))
		s.Contains(keys, "key1")
		s.Contains(keys, "key2")
		s.Contains(keys, "key3")
	})

	s.Run("TestValues", func() {
		s.intMap = NewGenericHashMap[string, int]()
		s.intMap.put("key1", 100)
		s.intMap.put("key2", 200)
		s.intMap.put("key3", 300)

		values := s.intMap.values()
		s.Equal(3, len(values))
		s.Contains(values, 100)
		s.Contains(values, 200)
		s.Contains(values, 300)
	})

	s.Run("TestResize", func() {
		// Create a small map to trigger resizing
		smallMap := NewGenericHashMapWithCapacity[string, int](2)

		// Add enough elements to trigger resize
		for i := 0; i < 10; i++ {
			smallMap.put(string(rune('a'+i)), i*10)
		}

		// Verify all elements are still accessible after resize
		for i := 0; i < 10; i++ {
			value, exists := smallMap.get(string(rune('a' + i)))
			s.True(exists)
			s.Equal(i*10, value)
		}
	})
}
