package typex

type Set[T comparable] interface {
	Contains(value T) bool
	Set(value T) bool
	Values() []T
	Remove(value T) bool
}
