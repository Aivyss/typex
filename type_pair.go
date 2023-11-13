package typex

import "github.com/aivyss/typex/util"

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
	notNullPairs := util.Filter(pairs, func(pair Pair[K, V]) bool {
		return !util.IsNil(pair.Value)
	})

	m := make(map[K]V, len(notNullPairs))
	for _, pair := range notNullPairs {
		m[pair.Key] = pair.Value
	}

	return m
}
