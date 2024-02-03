package utilx

import (
	"errors"
	"reflect"
)

var primitives = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.String,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Uintptr,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
}

var pointers = []reflect.Kind{
	reflect.Array,
	reflect.Chan,
	reflect.Func,
	reflect.Interface,
	reflect.Map,
	reflect.Pointer,
	reflect.Slice,
}

var structs = []reflect.Kind{
	reflect.Struct,
}

var invalids = []reflect.Kind{
	reflect.Invalid,
}

func IsNil[T any](value T) bool {
	valueOf := reflect.ValueOf(value)
	kind := valueOf.Kind()

	if inKindGroup(kind, primitives) {
		return false
	}

	if inKindGroup(kind, pointers) {
		return valueOf.IsNil()
	}

	if inKindGroup(kind, structs) {
		return false
	}

	if inKindGroup(kind, invalids) {
		return true
	}

	return true
}

func inKindGroup(k reflect.Kind, group []reflect.Kind) bool {
	flag := false
	for _, kind := range group {
		if k == kind {
			flag = true
			break
		}
	}

	return flag
}

func ConvertType[T any](original interface{}) (T, error) {
	if IsNil(original) {
		var t T
		return t, errors.New("fail to assert type to generic type(nil value)")
	}

	t, ok := original.(T)

	if !ok {
		var t T
		return t, errors.New("fail to assert type to generic type(not same type)")
	}

	return t, nil
}
