package typex

import (
	"errors"
	"github.com/aivyss/typex/util"
)

type Returns[T any] struct {
	value       *T
	initialized bool
	err         error
}

func (r *Returns[T]) Get() (T, error) {
	if util.IsNil(r.value) {
		var zeroValue T
		return zeroValue, r.err
	}

	return *r.value, nil
}

func (r *Returns[T]) Error() error {
	return r.err
}

func (r *Returns[T]) IsInit() bool {
	return r.initialized
}

func NewReturnsWithErr[T any](v *T, err error) Returns[T] {
	if util.IsNil(v) {
		return Returns[T]{
			value:       nil,
			initialized: false,
			err:         err,
		}
	}

	return Returns[T]{
		value:       v,
		initialized: true,
		err:         nil,
	}
}

func NewReturn[T any](v *T) Returns[T] {
	return NewReturnsWithErr(v, errors.New("nil pointer"))
}
