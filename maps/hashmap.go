package maps

import (
	"fmt"
	"hash/fnv"
)

type hashMapRequest[K comparable, V any] struct {
	action    string
	key       K
	value     V
	keys      []K
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

func NewGenericHashMap[K comparable, V any]() *GenericHashMap[K, V] {
	return NewGenericHashMapWithCapacityAndLoadFactor[K, V](16, 0.75)
}

func NewGenericHashMapWithCapacity[K comparable, V any](initialCapacity int) *GenericHashMap[K, V] {
	return NewGenericHashMapWithCapacityAndLoadFactor[K, V](initialCapacity, 0.75)
}

func NewGenericHashMapWithCapacityAndLoadFactor[K comparable, V any](initialCapacity int, loadFactor float64) *GenericHashMap[K, V] {
	hm := &GenericHashMap[K, V]{
		hashMapChan:     make(chan hashMapRequest[K, V]),
		buckets:         make([]*HashMapEntry[K, V], initialCapacity),
		initialCapacity: initialCapacity,
		loadFactor:      loadFactor,
		size:            0,
	}
	go hm.manageHashMap()
	return hm
}

func (hm *GenericHashMap[K, V]) manageHashMap() {
	for req := range hm.hashMapChan {
		switch req.action {
		case "put":
			req.replyChan <- hm.put(req.key, req.value)
		case "get":
			value, exists := hm.get(req.key)
			req.replyChan <- value
			req.replyChan <- exists
		case "remove":
			req.replyChan <- hm.remove(req.key)
		case "containsKey":
			req.replyChan <- hm.containsKey(req.key)
		case "size":
			req.replyChan <- hm.getSize()
		case "isEmpty":
			req.replyChan <- hm.isEmpty()
		case "clear":
			hm.clear()
			req.replyChan <- true
		case "keys":
			req.replyChan <- hm.keys()
		case "values":
			req.replyChan <- hm.values()
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

func (hm *GenericHashMap[K, V]) get(key K) (V, bool) {
	index := hm.hashWithCapacity(key, len(hm.buckets))
	current := hm.buckets[index]

	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}

	var zero V
	return zero, false
}

func (hm *GenericHashMap[K, V]) remove(key K) bool {
	index := hm.hashWithCapacity(key, len(hm.buckets))
	current := hm.buckets[index]

	if current == nil {
		return false
	}

	// Check if the first entry matches
	if current.Key == key {
		hm.buckets[index] = current.Next
		hm.size--
		return true
	}

	// Check remaining entries
	for current.Next != nil {
		if current.Next.Key == key {
			current.Next = current.Next.Next
			hm.size--
			return true
		}
		current = current.Next
	}

	return false
}

func (hm *GenericHashMap[K, V]) containsKey(key K) bool {
	_, exists := hm.get(key)
	return exists
}

func (hm *GenericHashMap[K, V]) getSize() int {
	return hm.size
}

func (hm *GenericHashMap[K, V]) isEmpty() bool {
	return hm.size == 0
}

func (hm *GenericHashMap[K, V]) clear() {
	hm.buckets = make([]*HashMapEntry[K, V], hm.initialCapacity)
	hm.size = 0
}

func (hm *GenericHashMap[K, V]) keys() []K {
	keys := make([]K, 0, hm.size)
	for _, bucket := range hm.buckets {
		for current := bucket; current != nil; current = current.Next {
			keys = append(keys, current.Key)
		}
	}
	return keys
}

func (hm *GenericHashMap[K, V]) values() []V {
	values := make([]V, 0, hm.size)
	for _, bucket := range hm.buckets {
		for current := bucket; current != nil; current = current.Next {
			values = append(values, current.Value)
		}
	}
	return values
}

func (hm *GenericHashMap[K, V]) hash(key K) int {
	return hm.hashWithCapacity(key, len(hm.buckets))
}

func (hm *GenericHashMap[K, V]) hashWithCapacity(key K, capacity int) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return int(h.Sum32()) % capacity
}

func (hm *GenericHashMap[K, V]) checkResize() {
	currentLoadFactor := float64(hm.size) / float64(len(hm.buckets))
	if currentLoadFactor >= hm.loadFactor {
		newCapacity := len(hm.buckets) * 2
		newBuckets := make([]*HashMapEntry[K, V], newCapacity)
		for _, entry := range hm.buckets {
			for current := entry; current != nil; current = current.Next {
				index := hm.hashWithCapacity(current.Key, newCapacity)
				newEntry := &HashMapEntry[K, V]{Key: current.Key, Value: current.Value, Next: newBuckets[index]}
				newBuckets[index] = newEntry
			}
		}
		hm.buckets = newBuckets
	}
}
