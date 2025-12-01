package maps

type GenericHashMapWrapper[K comparable, V any] struct {
	hashMap *GenericHashMap[K, V]
}

func NewGenericHashMapWrapper[K comparable, V any]() *GenericHashMapWrapper[K, V] {
	return &GenericHashMapWrapper[K, V]{hashMap: NewGenericHashMap[K, V]()}
}

func NewGenericHashMapWrapperWithCapacity[K comparable, V any](initialCapacity int) *GenericHashMapWrapper[K, V] {
	return &GenericHashMapWrapper[K, V]{hashMap: NewGenericHashMapWithCapacity[K, V](initialCapacity)}
}

func NewGenericHashMapWrapperWithCapacityAndLoadFactor[K comparable, V any](initialCapacity int, loadFactor float64) *GenericHashMapWrapper[K, V] {
	return &GenericHashMapWrapper[K, V]{hashMap: NewGenericHashMapWithCapacityAndLoadFactor[K, V](initialCapacity, loadFactor)}
}

func (w *GenericHashMapWrapper[K, V]) Put(key K, value V) bool {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "put", key: key, value: value, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericHashMapWrapper[K, V]) Get(key K) (V, bool) {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "get", key: key, replyChan: replyChan}
	value := (<-replyChan).(V)
	exists := (<-replyChan).(bool)
	return value, exists
}

func (w *GenericHashMapWrapper[K, V]) Remove(key K) bool {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "remove", key: key, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericHashMapWrapper[K, V]) ContainsKey(key K) bool {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "containsKey", key: key, replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericHashMapWrapper[K, V]) Size() int {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "size", replyChan: replyChan}
	return (<-replyChan).(int)
}

func (w *GenericHashMapWrapper[K, V]) IsEmpty() bool {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "isEmpty", replyChan: replyChan}
	return (<-replyChan).(bool)
}

func (w *GenericHashMapWrapper[K, V]) Clear() {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "clear", replyChan: replyChan}
	<-replyChan
}

func (w *GenericHashMapWrapper[K, V]) Keys() []K {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "keys", replyChan: replyChan}
	return (<-replyChan).([]K)
}

func (w *GenericHashMapWrapper[K, V]) Values() []V {
	replyChan := make(chan interface{})
	w.hashMap.hashMapChan <- hashMapRequest[K, V]{action: "values", replyChan: replyChan}
	return (<-replyChan).([]V)
}
