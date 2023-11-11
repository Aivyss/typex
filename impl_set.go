package typex

type DefaultSet[T comparable] map[T]bool

func NewSet[T comparable]() Set[T] {
	return DefaultSet[T](make(map[T]bool))
}

func NewSetWithArgs[T comparable](args ...T) Set[T] {
	set := DefaultSet[T](make(map[T]bool))
	for _, arg := range args {
		set.Set(arg)
	}

	return set
}

func (s DefaultSet[T]) Contains(value T) bool {
	raw := s.raw()
	if t, ok := raw[value]; t && ok {
		return true
	}

	return false
}

func (s DefaultSet[T]) Set(value T) bool {
	raw := s.raw()

	raw[value] = true

	return raw[value]
}

func (s DefaultSet[T]) Values() []T {
	raw := s.raw()
	results := make([]T, 0, len(raw))
	for value, ok := range raw {
		if ok {
			results = append(results, value)
		}
	}

	return results
}

func (s DefaultSet[T]) Remove(value T) bool {
	raw := s.raw()

	raw[value] = false

	return raw[value]
}

func (s DefaultSet[T]) raw() map[T]bool {
	return s
}
