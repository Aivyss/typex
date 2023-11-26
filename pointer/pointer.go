package pointer

import (
	"errors"
	"fmt"
	"github.com/aivyss/typex/types"
	"runtime"
)

func Pointer[T any](value T) (*T, error) {
	if types.IsNil(value) {
		return nil, errors.New("value is nil")
	}

	return &value, nil
}

func MustPointer[T any](value T) *T {
	pointer, err := Pointer(value)
	if err != nil {
		panic(err)
	}

	return pointer
}

func Value[T any](p *T) (t T, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			pc, _, _, _ := runtime.Caller(5)
			previousFunc := runtime.FuncForPC(pc)
			_, l := previousFunc.FileLine(pc)
			err = errors.New(fmt.Sprintf("fail to convert the pointer into a value (%s:%d)", previousFunc.Name(), l))
		}
	}()

	t = *p

	return t, err
}
