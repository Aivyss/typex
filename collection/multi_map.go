package collection

type MultiMap[K comparable, V any] interface {
	Put(key K, value V) MultiMap[K, V]
	Get(key K) []V
	Remove(key K) MultiMap[K, V]
	Clean() MultiMap[K, V]
	Keys() []K
	Values() [][]V
	Entries() []MultiMapEntry[K, V]
}

type MultiMapEntry[K comparable, V any] struct {
	Key    K
	Values []V
}

type multiMapDefault[K comparable, V any] struct {
	inside map[K][]V
}

func (m *multiMapDefault[K, V]) Put(key K, value V) MultiMap[K, V] {
	slice := m.Get(key)
	slice = append(slice, value)
	m.inside[key] = slice

	return m
}

func (m *multiMapDefault[K, V]) Get(key K) []V {
	return m.inside[key]
}

func (m *multiMapDefault[K, V]) Remove(key K) MultiMap[K, V] {
	m.inside[key] = []V{}

	return m
}

func (m *multiMapDefault[K, V]) Clean() MultiMap[K, V] {
	for key := range m.inside {
		m.inside[key] = []V{}
	}

	return m
}

func (m *multiMapDefault[K, V]) Keys() []K {
	var keys []K

	for key := range m.inside {
		keys = append(keys, key)
	}

	return keys
}

func (m *multiMapDefault[K, V]) Values() [][]V {
	var values [][]V

	for _, v := range m.inside {
		values = append(values, v)
	}

	return values
}

func (m *multiMapDefault[K, V]) Entries() []MultiMapEntry[K, V] {
	var result []MultiMapEntry[K, V]

	for k, v := range m.inside {
		result = append(result, MultiMapEntry[K, V]{Key: k, Values: v})
	}

	return result
}

func NewMultiMap[K comparable, V any]() MultiMap[K, V] {
	return &multiMapDefault[K, V]{
		inside: make(map[K][]V),
	}
}
