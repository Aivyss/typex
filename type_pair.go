package typex

import (
	"github.com/aivyss/typex/slice"
	"github.com/aivyss/typex/types"
)

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func CreateMap[K comparable, V any](pairs []Pair[K, V]) map[K]V {
	m := make(map[K]V, len(pairs))
	for _, pair := range pairs {
		m[pair.Key] = pair.Value
	}

	return m
}

func CreateNotNilMap[K comparable, V any](pairs []Pair[K, V]) map[K]V {
	notNullPairs := slice.Filter(pairs, func(pair Pair[K, V]) bool {
		return !types.IsNil(pair.Value)
	})

	m := make(map[K]V, len(notNullPairs))
	for _, pair := range notNullPairs {
		m[pair.Key] = pair.Value
	}

	return m
}
