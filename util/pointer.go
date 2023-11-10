package util

import "errors"

func Pointer[T any](value T) (*T, error) {
	if IsNil(value) {
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
