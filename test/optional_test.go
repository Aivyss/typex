package test

import (
	"github.com/aivyss/typex"
	"github.com/aivyss/typex/pointer"
	"testing"
)

func TestOpt(t *testing.T) {
	t.Run("ifPresent-elseDo", func(t *testing.T) {
		var num *int = nil
		defaultValue := pointer.MustPointer(123)

		opt := typex.MustOptWithDefault(num, defaultValue)
		opt.IfPresent(func(value *int) {
			t.Fatal("wrong result")
		})

		opt.IfPresentWithDefault(func(value *int) {
			if value != defaultValue {
				t.Fatal("not same value")
			}
		}).ElseDo(func() {
			t.Fatal("wrong result")
		})
	})

	t.Run("strings", func(t *testing.T) {
		str := "value"
		defaultStr := "default"

		opt := typex.MustOptWithDefault(str, defaultStr)
		value := opt.MustGetOrDefault()

		if value != str {
			t.Fatal("not same value")
		}

		opt.IfPresentWithDefault(func(value string) {
			if value != str {
				t.Fatal("not same value")
			}
		})
	})

	t.Run("*strings", func(t *testing.T) {
		var strNil *string = nil
		strPointer, err := pointer.Pointer("strings")
		if err != nil {
			t.Fatal(err)
		}

		opt := typex.MustOptWithDefault(strNil, strPointer)
		defaultValue := opt.MustGetOrDefault()

		if strPointer != defaultValue {
			t.Fatal("not same value")
		}

		opt.IfPresentWithDefault(func(value *string) {
			if value == nil {
				t.Fatal("value is nil")
			}
		})
	})

	t.Run("[]strings", func(t *testing.T) {
		var sliceNil []string = nil
		expectedDefault := []string{"test"}
		opt := typex.MustOptWithDefault(sliceNil, expectedDefault)
		defaultValue := opt.MustGetOrDefault()

		if len(defaultValue) != 1 || defaultValue[0] != "test" {
			t.Fatal("not same value")
		}
	})

	t.Run("*[]strings", func(t *testing.T) {
		var sliceNil *[]string = nil
		expectedDefault := &[]string{"test"}
		opt := typex.MustOptWithDefault(sliceNil, expectedDefault)
		defaultValue := opt.MustGetOrDefault()

		if defaultValue == nil || len(*defaultValue) != 1 || (*defaultValue)[0] != "test" {
			t.Fatal("not same value")
		}
	})

	t.Run("map", func(t *testing.T) {
		var mapNil map[string]any = nil
		expectedDefault := map[string]any{"aa": true}
		opt := typex.MustOptWithDefault(mapNil, expectedDefault)
		defaultValue := opt.MustGetOrDefault()

		if defaultValue == nil || len(defaultValue) == 0 || defaultValue["aa"] == nil {
			t.Fatal("not same value")
		}
	})

	t.Run("*map", func(t *testing.T) {
		var mapNil *map[string]any = nil
		expectedDefault := pointer.MustPointer(map[string]any{"aa": true})
		opt := typex.MustOptWithDefault(mapNil, expectedDefault)
		defaultValue := opt.MustGetOrDefault()

		if defaultValue == nil || len(*defaultValue) == 0 || (*defaultValue)["aa"] == nil {
			t.Fatal("not same value")
		}
	})

	t.Run("struct", func(t *testing.T) {
		type someType struct {
			value string
		}

		value := someType{value: "test"}
		defaultValue := someType{value: "defaultValue"}
		opt := typex.MustOptWithDefault(value, defaultValue)
		actual := opt.MustGetOrDefault()

		if actual != value {
			t.Fatal("not same value")
		}
	})

	t.Run("*struct", func(t *testing.T) {
		type someType struct {
			value string
		}

		var value *someType = nil
		defaultValue := &someType{value: "defaultValue"}
		opt := typex.MustOptWithDefault(value, defaultValue)
		actual := opt.MustGetOrDefault()

		if actual != defaultValue {
			t.Fatal("not same value")
		}
	})
}
