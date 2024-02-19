package utilx

import "reflect"

func GetGenericType[T any]() reflect.Type {
	return reflect.TypeOf(func(t T) {}).In(0)
}
