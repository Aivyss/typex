package test

import (
	"fmt"
	"github.com/aivyss/typex"
	"github.com/aivyss/typex/pointer"
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
		var testVal = pointer.MustPointer("abcd")
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

	t.Run("temp", func(t *testing.T) {
		type testStruct struct {
			Value string
		}

		v := &testStruct{
			Value: "test1",
		}
		object := typex.NewReturn(v)
		st, err := object.Get()
		if err != nil {
			fmt.Println("err =", err)
		}
		fmt.Println("st.Value =", st.Value)
		v.Value = "test2"
		st, err = object.Get()
		if err != nil {
			fmt.Println("err =", err)
		}
		fmt.Println("st.Value =", st.Value)
	})
}
