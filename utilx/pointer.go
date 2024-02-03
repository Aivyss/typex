package utilx

import (
	"errors"
	"fmt"
	"runtime"
)

func Pointer[T any](value T) *T {
	return &value
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
