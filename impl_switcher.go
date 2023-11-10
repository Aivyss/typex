package typex

import (
	"errors"
	"github.com/aivyss/typex/util"
)

type DefaultSwitcher[T comparable] struct {
	m   map[T]func(v *T)
	def func(v *T)
}

func NewSwitcher[T comparable]() Switcher[T] {
	return &DefaultSwitcher[T]{
		m:   make(map[T]func(v *T)),
		def: func(v *T) {},
	}
}

func (s *DefaultSwitcher[T]) Set(key T, consumer func(v *T)) error {
	if util.IsNil(consumer) {
		return errors.New("consumer is nil")
	}

	s.m[key] = consumer

	return nil
}

func (s *DefaultSwitcher[T]) SetDefault(consumer func(v *T)) error {
	if util.IsNil(consumer) {
		return errors.New("consumer is nil")
	}

	s.def = consumer

	return nil
}

func (s *DefaultSwitcher[T]) Switch(key T) error {
	consumer := s.m[key]

	if util.IsNil(consumer) {
		return errors.New("consumer is nil")
	}

	consumer(&key)

	return nil
}
