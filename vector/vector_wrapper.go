package vector

type WrapperVector[T any] struct {
	vector *GenericVector[T]
}

func NewWrapperVector[T any](initialCapacity ...int) *WrapperVector[T] {
	var capacity int
	if len(initialCapacity) > 0 {
		capacity = initialCapacity[0]
	} else {
		capacity = 0
	}
	return &WrapperVector[T]{
		vector: NewGenericVector[T](capacity),
	}
}

func (v *WrapperVector[T]) Add(item T) error {
	return v.vector.Add(item)
}

func (v *WrapperVector[T]) AddAt(index int, item T) error {
	return v.vector.AddAt(index, item)
}

func (v *WrapperVector[T]) RemoveAt(index int) error {
	return v.vector.RemoveAt(index)
}

func (v *WrapperVector[T]) Set(index int, item T) error {
	return v.vector.Set(index, item)
}

func (v *WrapperVector[T]) Get(index int) (T, error) {
	return v.vector.Get(index)
}

func (v *WrapperVector[T]) Size() int {
	return v.vector.Size()
}

func (v *WrapperVector[T]) Clear() error {
	return v.vector.Clear()
}

func (v *WrapperVector[T]) IsEmpty() bool {
	return v.vector.IsEmpty()
}

func (v *WrapperVector[T]) EnsureCapacity(capacity int) {
	v.vector.EnsureCapacity(capacity)
}

func (v *WrapperVector[T]) ToArray() []T {
	return v.vector.ToArray()
}

func (v *WrapperVector[T]) TrimToSize() error {
	return v.vector.TrimToSize()
}
