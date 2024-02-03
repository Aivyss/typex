package utilx

import (
	"errors"
)

type Switcher[T comparable] interface {
	Set(key T, consumer func(v *T)) error
	SetDefault(consumer func(v *T)) error
	Switch(key T) error
}

type defaultSwitcher[T comparable] struct {
	m   map[T]func(v *T)
	def func(v *T)
}

func NewSwitcher[T comparable]() Switcher[T] {
	return &defaultSwitcher[T]{
		m:   make(map[T]func(v *T)),
		def: func(v *T) {},
	}
}

func (s *defaultSwitcher[T]) Set(key T, consumer func(v *T)) error {
	if IsNil(consumer) {
		return errors.New("consumer is nil")
	}

	s.m[key] = consumer

	return nil
}

func (s *defaultSwitcher[T]) SetDefault(consumer func(v *T)) error {
	if IsNil(consumer) {
		return errors.New("consumer is nil")
	}

	s.def = consumer

	return nil
}

func (s *defaultSwitcher[T]) Switch(key T) error {
	consumer := s.m[key]

	if IsNil(consumer) {
		return errors.New("consumer is nil")
	}

	consumer(&key)

	return nil
}
