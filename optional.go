package typex

import (
	"errors"
	"github.com/aivyss/typex/util"
)

type Optional[T any] struct {
	value        T
	defaultValue T
	defaultInit  bool
}

type ElseDo struct {
	flag bool
}

func (o *Optional[T]) SetMustDefault(value T) *Optional[T] {
	if util.IsNil(value) {
		panic("default value is nil")
	}

	o.defaultValue = value
	o.defaultInit = true

	return o
}

func (o *Optional[T]) SetDefault(value T) (*Optional[T], error) {
	if util.IsNil(value) {
		return nil, errors.New("default value is nil")
	}

	o.defaultValue = value
	o.defaultInit = true

	return o, nil
}

func (o *Optional[T]) Get() (T, error) {
	if util.IsNil(o.value) {
		return o.value, errors.New("value is nil")
	}

	return o.value, nil
}

func (o *Optional[T]) MustGet() T {
	value, err := o.Get()
	if err != nil {
		panic(err)
	}

	return value
}

func (o *Optional[T]) GetOrDefault() (T, error) {
	value, err := o.Get()
	if err != nil {
		if !o.defaultInit {
			return o.defaultValue, errors.New("default value is nil")
		}

		return o.defaultValue, nil
	}

	return value, nil
}

func (o *Optional[T]) MustGetOrDefault() T {
	value, err := o.GetOrDefault()
	if err != nil {
		panic(err)
	}

	return value
}

func (o *Optional[T]) IfPresent(consumer func(value T)) *ElseDo {
	flag := true
	if !util.IsNil(o.value) {
		consumer(o.value)
		flag = false
	}

	return &ElseDo{flag: flag}
}

func (o *Optional[T]) IfPresentWithDefault(consumer func(value T)) *ElseDo {
	if !util.IsNil(o.value) {
		consumer(o.value)
		return &ElseDo{flag: false}
	}

	flag := true
	if o.defaultInit {
		consumer(o.defaultValue)
		flag = false
	}

	return &ElseDo{flag: flag}
}

func (e *ElseDo) ElseDo(runnable func()) {
	if e.flag {
		runnable()
	}
}

func Opt[T any](value T) *Optional[T] {
	return &Optional[T]{
		value:       value,
		defaultInit: false,
	}
}

func OptWithDefault[T any](value T, defaultValue T) (*Optional[T], error) {
	if util.IsNil(defaultValue) {
		return nil, errors.New("default value is nil")
	}

	opt := &Optional[T]{
		value:        value,
		defaultValue: defaultValue,
		defaultInit:  true,
	}

	return opt, nil
}

func MustOptWithDefault[T any](value T, defaultValue T) *Optional[T] {
	opt, err := OptWithDefault(value, defaultValue)
	if err != nil {
		panic(err)
	}

	return opt
}
