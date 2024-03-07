package collection

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestForEach(t *testing.T) {
	var sum int
	var mutex sync.Mutex
	testSlice := Range(1, 10)
	ForEach(testSlice, func(elem int) {
		mutex.Lock()
		sum += elem
		mutex.Unlock()
	})

	assert.Equal(t, 55, sum)
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

	actual := Map(testSlice, func(o original) convert {
		mutex.Lock()
		sum += o.Value
		mutex.Unlock()
		return convert{Value: o.Value}
	})

	assert.Equal(t, expected, actual)
	assert.Equal(t, 55, sum)
}

func TestRange(t *testing.T) {
	t.Run("ascend", func(t *testing.T) {
		actual := Range(1, 10)
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		assert.Equal(t, expected, actual)
	})

	t.Run("descend", func(t *testing.T) {
		actual := Range(5, -3)
		expected := []int{5, 4, 3, 2, 1, 0, -1, -2, -3}

		assert.Equal(t, expected, actual)
	})
}

func TestContains(t *testing.T) {
	nums := Range(1, 10)
	assert.True(t, Contains(nums, 1))
	assert.False(t, Contains(nums, 11))
}

func TestFindFirst(t *testing.T) {
	strs := []string{"apple", "banana", "tomato"}

	first := FindFirst(strs, func(str *string) bool {
		return *str == "banana"
	})
	assert.NotNil(t, first)
	assert.Equal(t, "banana", *first)

	first = FindFirst(strs, func(str *string) bool {
		return *str == "graph"
	})
	assert.Nil(t, first)
}
