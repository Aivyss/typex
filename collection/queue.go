package collection

import "sync"

type Queue[T any] interface {
	Append(elem T)
	Len() int
	IsEmpty() bool
	Pop() (T, bool)
	Get() (T, bool)
}

func NewThreadUnsafeQueue[T any]() Queue[T] {
	return &threadUnsafeQueue[T]{
		list: []T{},
	}
}

type threadUnsafeQueue[T any] struct {
	list []T
}

func (t *threadUnsafeQueue[T]) Append(elem T) {
	t.list = append(t.list, elem)
}

func (t *threadUnsafeQueue[T]) Pop() (T, bool) {
	if t.IsEmpty() {
		return *new(T), false
	}

	elem := t.list[0]
	t.list = t.list[1:]
	return elem, true
}

func (t *threadUnsafeQueue[T]) Get() (T, bool) {
	if t.IsEmpty() {
		return *new(T), false
	}

	return t.list[0], true
}

func (t *threadUnsafeQueue[T]) IsEmpty() bool {
	return t.Len() == 0
}

func (t *threadUnsafeQueue[T]) Len() int {
	return len(t.list)
}

func NewThreadSafeQueue[T any]() Queue[T] {
	return &threadSafeQueue[T]{
		list: []T{},
	}
}

type threadSafeQueue[T any] struct {
	mutex sync.Mutex
	list  []T
}

func (t *threadSafeQueue[T]) Append(elem T) {
	t.mutex.Lock()
	t.list = append(t.list, elem)
	t.mutex.Unlock()
}

func (t *threadSafeQueue[T]) Pop() (T, bool) {
	if t.IsEmpty() {
		return *new(T), false
	}

	t.mutex.Lock()
	elem := t.list[0]
	t.list = t.list[1:]
	t.mutex.Unlock()

	return elem, true
}

func (t *threadSafeQueue[T]) Get() (T, bool) {
	if t.IsEmpty() {
		return *new(T), false
	}

	t.mutex.Lock()
	elem := t.list[0]
	t.mutex.Unlock()

	return elem, true
}

func (t *threadSafeQueue[T]) IsEmpty() bool {
	return t.Len() == 0
}

func (t *threadSafeQueue[T]) Len() int {
	t.mutex.Lock()
	size := len(t.list)
	t.mutex.Unlock()

	return size
}
