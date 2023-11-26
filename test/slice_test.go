package test

import (
	"github.com/aivyss/typex/slice"
	"github.com/aivyss/typex/testTool"
	"sync"
	"testing"
)

func TestForEach(t *testing.T) {
	var sum int
	var mutex sync.Mutex
	testSlice := slice.Range(1, 10)
	slice.ForEach(testSlice, func(elem int) {
		mutex.Lock()
		sum += elem
		mutex.Unlock()
	})

	testTool.Equal(t, 55, sum)
}

func TestMap(t *testing.T) {
	var sum int
	var mutex sync.Mutex

	type original struct {
		Value int
	}

	type convert struct {
		Value int
	}
	testSlice := make([]original, 0, 10)
	expected := make([]convert, 0, 10)
	for i := 0; i < 10; i++ {
		testSlice = append(testSlice, original{Value: i + 1})
		expected = append(expected, convert{Value: i + 1})
	}

	actual := slice.Map(testSlice, func(o original) convert {
		mutex.Lock()
		sum += o.Value
		mutex.Unlock()
		return convert{Value: o.Value}
	})

	testTool.EqualSlice(t, expected, actual)
	testTool.Equal(t, 55, sum)
}

func TestRange(t *testing.T) {
	t.Run("ascend", func(t *testing.T) {
		actual := slice.Range(1, 10)
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		testTool.EqualSlice(t, expected, actual)
	})

	t.Run("descend", func(t *testing.T) {
		actual := slice.Range(5, -3)
		expected := []int{5, 4, 3, 2, 1, 0, -1, -2, -3}

		testTool.EqualSlice(t, expected, actual)
	})
}

func TestContains(t *testing.T) {
	nums := slice.Range(1, 10)
	testTool.True(t, slice.Contains(nums, 1))
	testTool.False(t, slice.Contains(nums, 11))
}
