package utilx

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type TestInterface interface{}
type TestStruct struct{}
type TestInterface2 interface {
	dummy() error
}

func TestGetGenericType(t *testing.T) {
	var testInterface TestInterface = TestStruct{}
	var testStruct TestStruct

	assert.True(t, reflect.TypeOf(testInterface).ConvertibleTo(GetGenericType[TestInterface]()))
	assert.False(t, reflect.TypeOf(testInterface).ConvertibleTo(GetGenericType[TestInterface2]()))
	assert.True(t, reflect.TypeOf(testInterface).AssignableTo(GetGenericType[TestInterface]()))
	assert.False(t, reflect.TypeOf(testInterface).AssignableTo(GetGenericType[TestInterface2]()))
	assert.Equal(t, reflect.TypeOf(testStruct), GetGenericType[TestStruct]())
}
