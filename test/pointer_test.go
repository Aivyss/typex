package test

import (
	"github.com/aivyss/typex/pointer"
	"github.com/aivyss/typex/testTool"
	"testing"
)

func TestPointer(t *testing.T) {
	var nilStrPointer *string = nil
	_, err := pointer.Pointer(nilStrPointer)
	testTool.NotNil(t, err)

	testStr := "abcd"
	_, err = pointer.Pointer(testStr)
	testTool.Nil(t, err)

	pp, err := pointer.Pointer(&testStr)
	testTool.Nil(t, err)
	testTool.Equal(t, testStr, **pp)
}

func TestValue(t *testing.T) {
	t.Run("[func Value] no error", func(t *testing.T) {
		var testStr = pointer.MustPointer("")
		_, err := pointer.Value(testStr)
		if err != nil {
			t.Fatal("unexpected error 1")
		}
	})

	t.Run("[func Value] error", func(t *testing.T) {
		var testStr *string = nil
		_, err := pointer.Value(testStr)
		if err == nil {
			t.Fatal("unexpected error 1")
		}
		t.Log(err)
	})
}
