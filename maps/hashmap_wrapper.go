package maps

type GenericHashMapWrapper[K comparable, V any] struct {
	hashMap *GenericHashMap[K, V]
}

// NewGenericHashMapWrapper creates a new instance of GenericHashMapWrapper with default capacity
func NewGenericHashMapWrapper[K comparable, V any]() *GenericHashMapWrapper[K, V] {
	return &GenericHashMapWrapper[K, V]{hashMap: NewGenericHashMap[K, V]()}
}

// NewGenericHashMapWrapperWithCapacity creates a new instance with specified initial capacity
func NewGenericHashMapWrapperWithCapacity[K comparable, V any](initialCapacity int) *GenericHashMapWrapper[K, V] {
	return &GenericHashMapWrapper[K, V]{hashMap: NewGenericHashMapWithCapacity[K, V](initialCapacity)}
}

// NewGenericHashMapWrapperWithCapacityAndLoadFactor creates a new instance with specified capacity and load factor
func NewGenericHashMapWrapperWithCapacityAndLoadFactor[K comparable, V any](initialCapacity int, loadFactor float64) *GenericHashMapWrapper[K, V] {
	return &GenericHashMapWrapper[K, V]{hashMap: NewGenericHashMapWithCapacityAndLoadFactor[K, V](initialCapacity, loadFactor)}
}

// Put adds or updates a key-value pair in the map. Returns true if new key was added, false if existing key was updated
func (w *GenericHashMapWrapper[K, V]) Put(key K, value V) bool {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "put", key: key, value: value, replyChan: replyChan}
	return (<-replyChan).(bool)
}

// Get retrieves the value associated with the given key. Returns the value and true if found, zero value and false otherwise
func (w *GenericHashMapWrapper[K, V]) Get(key K) (V, bool) {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "get", key: key, replyChan: replyChan}
	value := (<-replyChan).(V)
	exists := (<-replyChan).(bool)
	return value, exists
}

// Remove removes the key-value pair for the given key. Returns true if the key was found and removed
func (w *GenericHashMapWrapper[K, V]) Remove(key K) bool {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "remove", key: key, replyChan: replyChan}
	return (<-replyChan).(bool)
}

// ContainsKey checks if the map contains the given key
func (w *GenericHashMapWrapper[K, V]) ContainsKey(key K) bool {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "containsKey", key: key, replyChan: replyChan}
	return (<-replyChan).(bool)
}

// Size returns the number of key-value pairs in the map
func (w *GenericHashMapWrapper[K, V]) Size() int {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "size", replyChan: replyChan}
	return (<-replyChan).(int)
}

// IsEmpty returns true if the map contains no key-value pairs
func (w *GenericHashMapWrapper[K, V]) IsEmpty() bool {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "isEmpty", replyChan: replyChan}
	return (<-replyChan).(bool)
}

// Clear removes all key-value pairs from the map
func (w *GenericHashMapWrapper[K, V]) Clear() {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "clear", replyChan: replyChan}
	<-replyChan
}

// Keys returns a slice of all keys in the map
func (w *GenericHashMapWrapper[K, V]) Keys() []K {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "keys", replyChan: replyChan}
	return (<-replyChan).([]K)
}

// Values returns a slice of all values in the map
func (w *GenericHashMapWrapper[K, V]) Values() []V {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "values", replyChan: replyChan}
	return (<-replyChan).([]V)
}
