package test

import (
	"github.com/aivyss/typex"
	"github.com/aivyss/typex/util"
	"testing"
)

func TestReturns(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var testVal *string = nil
		object := typex.NewReturn(testVal)

		if object.IsInit() {
			t.Fatal("unexpected result1")
		}

		if object.Error() == nil {
			t.Fatal("unexpected result2")
		}

		val, err := object.Get()
		if val != "" || err == nil {
			t.Fatal("unexpected result3")
		}
		t.Log(err)
	})

	t.Run("not nil", func(t *testing.T) {
		var testVal = util.MustPointer("abcd")
		object := typex.NewReturn(testVal)

		if !object.IsInit() {
			t.Fatal("unexpected result1")
		}

		if object.Error() != nil {
			t.Fatal("unexpected result2")
		}

		val, err := object.Get()
		if val == "" || err != nil {
			t.Fatal("unexpected result3")
		}
	})
}
