package collection

import (
	"github.com/aivyss/typex/utilx"
)

type Pair[A, B any] struct {
	First  A
	Second B
}

type MapPair[K comparable, V any] struct {
	Key   K
	Value V
}

func CreateMap[K comparable, V any](pairs []MapPair[K, V]) map[K]V {
	m := make(map[K]V, len(pairs))
	for _, pair := range pairs {
		m[pair.Key] = pair.Value
	}

	return m
}

func CreateNotNilMap[K comparable, V any](pairs []MapPair[K, V]) map[K]V {
	notNullPairs := Filter(pairs, func(pair MapPair[K, V]) bool {
		return !utilx.IsNil(pair.Value)
	})

	m := make(map[K]V, len(notNullPairs))
	for _, pair := range notNullPairs {
		m[pair.Key] = pair.Value
	}

	return m
}
