package test

import (
	"github.com/aivyss/typex/util"
	"testing"
)

func TestValue(t *testing.T) {
	t.Run("[func Value] no error", func(t *testing.T) {
		var testStr = util.MustPointer("")
		_, err := util.Value(testStr)
		if err != nil {
			t.Fatal("unexpected error 1")
		}
	})

	t.Run("[func Value] error", func(t *testing.T) {
		var testStr *string = nil
		_, err := util.Value(testStr)
		if err == nil {
			t.Fatal("unexpected error 1")
		}
		t.Log(err)
	})
}
