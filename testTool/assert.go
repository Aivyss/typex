package testTool

import (
	"github.com/aivyss/typex/types"
	"testing"
)

func True(t *testing.T, v any) {
	checkNilT(t)

	boolean, err := types.ConvertType[bool](v)
	if err != nil {
		booleanPointer, err := types.ConvertType[*bool](v)
		if err != nil {
			t.Logf("not a bool / bool pointer\n")
			t.Logf("value: %v\n", v)
			t.Fail()
		}

		if !*booleanPointer {
			t.Logf("false value\n")
			t.Fail()
		}
	}

	if !boolean {
		t.Logf("false value\n")
		t.Fail()
	}
}

func False(t *testing.T, v any) {
	checkNilT(t)

	boolean, err := types.ConvertType[bool](v)
	if err != nil {
		booleanPointer, err := types.ConvertType[*bool](v)
		if err != nil {
			t.Logf("not a bool / bool pointer\n")
			t.Logf("value: %v\n", v)
			t.Fail()
		}

		if *booleanPointer {
			t.Logf("true value\n")
			t.Fail()
		}
	}

	if boolean {
		t.Logf("true value\n")
		t.Fail()
	}
}

func Equal[T comparable](t *testing.T, expected T, actual T) {
	checkNilT(t)

	if types.IsNil(expected) && types.IsNil(actual) {
		return
	}

	if types.IsNil(expected) || types.IsNil(actual) {
		t.Logf("not a equal value\n")
		t.Fail()
		return
	}

	if expected != actual {
		t.Logf("not a equal value\n")
		t.Fail()
		return
	}
}

func EqualSlice[T comparable](t *testing.T, expected []T, actual []T) {
	actualLen := len(actual)
	if len(expected) != actualLen {
		t.Logf("not a equal slice (not same size)\n")
		t.Logf("expected: %v\n", expected)
		t.Logf("actual: %v\n", actual)
		t.Fail()
		return
	}

	for i := 0; i < actualLen; i++ {
		if actual[i] != expected[i] {
			t.Logf("not a equal slice\n")
			t.Logf("[index %d] expected: %v actual: %v\n", i, expected[i], actual[i])
			t.Fail()
			return
		}
	}
}

func checkNilT(t *testing.T) {
	if t == nil {
		t.Logf("*testing.T is nil")
		t.Fail()
		return
	}
}

func Nil(t *testing.T, v any) {
	if !types.IsNil(v) {
		t.Logf("value isn't nil\n")
		t.Fail()
	}
}

func NotNil(t *testing.T, v any) {
	if types.IsNil(v) {
		t.Logf("value is nil\n")
		t.Fail()
	}
}
