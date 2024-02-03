package collection

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestHashSet(t *testing.T) {
	t.Run("【test】constructor", func(t *testing.T) {
		// given
		elements := []string{"a", "b"}

		// then
		s1 := NewSet(elements...)

		// when
		assert.False(t, s1.IsEmpty())
		slice := s1.ToSlice()
		sort.Slice(elements, func(i int, j int) bool {
			return elements[i] < elements[j]
		})
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
		assert.Equal(t, elements, slice)
	})

	t.Run("【test】Add", func(t *testing.T) {
		// given
		s := NewSet[string]()
		slice := s.ToSlice()
		assert.Equal(t, 0, len(slice))

		// then
		s.Add("a")

		// when
		slice = s.ToSlice()
		assert.Equal(t, 1, len(slice))
		assert.Equal(t, "a", slice[0])
	})

	t.Run("【test】AddAll / ToSlice", func(t *testing.T) {
		// given
		elements := []string{"a", "b", "a"}
		expectedSlice := []string{"a", "b"}

		// then
		s := NewSet[string]()
		s.AddAll(elements)

		// when
		slice := s.ToSlice()
		sort.Slice(expectedSlice, func(i, j int) bool {
			return expectedSlice[i] < expectedSlice[j]
		})
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
		assert.Equal(t, expectedSlice, slice)
	})

	t.Run("【test】Contains", func(t *testing.T) {
		// given
		elements := []string{"a", "b", "a"}
		expectedValues := []string{"a", "b"}
		notExpectedValues := []string{"c", "aa", "bc"}
		s := NewSet(elements...)

		// then
		for _, expectedValue := range expectedValues {
			assert.True(t, s.Contains(expectedValue))
		}
		for _, notExpectedValue := range notExpectedValues {
			assert.False(t, s.Contains(notExpectedValue))
		}
	})

	t.Run("【test】ContainsAll", func(t *testing.T) {
		// given
		elements := []string{"a", "b", "a"}
		expectedValues := []string{"a", "b"}
		notExpectedValues := []string{"c", "aa", "bc"}
		s := NewSet(elements...)

		// then
		assert.True(t, s.ContainsAll(expectedValues))
		assert.False(t, s.ContainsAll(notExpectedValues))
		assert.False(t, s.ContainsAll(append(expectedValues, notExpectedValues...)))
	})

	t.Run("【test】IsEmpty", func(t *testing.T) {
		// given
		s := NewSet[string]()

		// then
		assert.True(t, s.IsEmpty())

		s.Add("something")
		assert.False(t, s.IsEmpty())
	})

	t.Run("【test】Remove", func(t *testing.T) {
		// given
		elements := []string{"a", "b", "a"}
		s := NewSet(elements...)

		// then
		assert.True(t, s.Contains(elements[0]))
		s.Remove(elements[0])
		assert.False(t, s.Contains(elements[0]))
	})

	t.Run("【test】RemoveAll", func(t *testing.T) {
		// given
		elements := []string{"a", "b", "a", "c"}
		expectedValues := []string{"a", "b", "c"}
		s := NewSet(elements...)

		// then
		assert.True(t, s.ContainsAll(expectedValues))
		s.RemoveAll(expectedValues)
		assert.False(t, s.Contains(expectedValues[0]))
		assert.False(t, s.Contains(expectedValues[1]))
		assert.False(t, s.Contains(expectedValues[2]))
		assert.True(t, s.IsEmpty())
	})

	t.Run("【test】Clear", func(t *testing.T) {
		// given
		elements := []string{"a", "b", "a", "c"}
		s := NewSet(elements...)

		// then
		assert.False(t, s.IsEmpty())
		s.Clear()
		assert.True(t, s.IsEmpty())
	})
}
