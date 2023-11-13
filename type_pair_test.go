package typex

import (
	"github.com/aivyss/typex/util"
	"testing"
)

func TestPairs(t *testing.T) {
	m := CreateMap[string, string]([]Pair[string, string]{
		{Key: "key1", Value: "value1"},
		{Key: "key2", Value: "value2"},
		{Key: "key3", Value: "value3"},
	})

	_, ok1 := m["key1"]
	_, ok2 := m["key2"]
	_, ok3 := m["key3"]

	if !ok1 || !ok2 || !ok3 || len(m) != 3 {
		t.Fatal("unexpected result 1")
	}

	m2 := CreateNotNilMap[string, *string]([]Pair[string, *string]{
		{Key: "key1", Value: util.MustPointer("value1")},
		{Key: "key2", Value: nil},
		{Key: "key3", Value: util.MustPointer("value3")},
	})

	_, ok1 = m2["key1"]
	_, ok2 = m2["key2"]
	_, ok3 = m2["key3"]

	if !ok1 || ok2 || !ok3 || len(m2) != 2 {
		t.Fatal("unexpected result 2")
	}
}
