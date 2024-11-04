package maps

import "unsafe"

type hashMapRequest[K comparable, V any] struct {
	action    string
	key       K
	value     V
	replyChan chan interface{}
}

type HashMapEntry[K comparable, V any] struct {
	Key   K
	Value V
	Next  *HashMapEntry[K, V]
}

type GenericHashMap[K comparable, V any] struct {
	hashMapChan     chan hashMapRequest[K, V]
	buckets         []*HashMapEntry[K, V]
	initialCapacity int
	loadFactor      float64
	size            int
}

// NewGenericHashMap creates a new GenericHashMap with default initial capacity and load factor
func NewGenericHashMap[K comparable, V any]() *GenericHashMap[K, V] {
	return NewGenericHashMapWithCapacityAndLoadFactor[K, V](16, 0.75)
}

// NewGenericHashMapWithCapacity creates a new GenericHashMap with specified initial capacity
func NewGenericHashMapWithCapacity[K comparable, V any](initialCapacity int) *GenericHashMap[K, V] {
	return NewGenericHashMapWithCapacityAndLoadFactor[K, V](initialCapacity, 0.75)
}

// NewGenericHashMapWithCapacityAndLoadFactor creates a new GenericHashMap with specified initial capacity and load factor
func NewGenericHashMapWithCapacityAndLoadFactor[K comparable, V any](initialCapacity int, loadFactor float64) *GenericHashMap[K, V] {
	return &GenericHashMap[K, V]{
		hashMapChan:     make(chan hashMapRequest[K, V]),
		buckets:         make([]*HashMapEntry[K, V], initialCapacity),
		initialCapacity: initialCapacity,
		loadFactor:      loadFactor,
	}
}

func (hm *GenericHashMap[K, V]) manageHashMap() {
	for req := range hm.hashMapChan {
		switch req.action {
		case "put":
			req.replyChan <- hm.put(req.key, req.value)
		}
	}
}

func (hm *GenericHashMap[K, V]) put(key K, value V) bool {
	index := hm.hashWithCapacity(key, len(hm.buckets))
	entry := &HashMapEntry[K, V]{Key: key, Value: value}

	if hm.buckets[index] == nil {
		hm.buckets[index] = entry
		hm.size++
		hm.checkResize()
		return true
	} else {
		current := hm.buckets[index]
		for current != nil {
			if current.Key == key {
				current.Value = value
				return false
			}
			if current.Next == nil {
				break
			}
			current = current.Next
		}
		current.Next = entry
		hm.size++
		hm.checkResize()
		return true
	}
}

func (hm *GenericHashMap[K, V]) hash(key K) int {
	return int(uintptr(unsafe.Pointer(&key)) % uintptr(len(hm.buckets)))
}

func (hm *GenericHashMap[K, V]) checkResize() {
	currentLoadFactor := float64(hm.size) / float64(len(hm.buckets))
	if currentLoadFactor >= hm.loadFactor {
		newCapacity := len(hm.buckets) * 2
		newBuckets := make([]*HashMapEntry[K, V], newCapacity)
		for _, entry := range hm.buckets {
			for current := entry; current != nil; current = current.Next {
				index := hm.hashWithCapacity(current.Key, newCapacity)
				newBuckets[index] = &HashMapEntry[K, V]{Key: current.Key, Value: current.Value, Next: newBuckets[index]}
			}
		}
		hm.buckets = newBuckets
	}
}

func (hm *GenericHashMap[K, V]) hashWithCapacity(key K, capacity int) int {
	return int(uintptr(unsafe.Pointer(&key)) % uintptr(capacity))
}
