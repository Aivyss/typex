package collection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestThreadUnsafeQueue(t *testing.T) {
	queue := NewThreadUnsafeQueue[int]()

	size := 100
	inputs := make([]int, 0, size)
	for i := 0; i < size; i++ {
		randNum := rand.Intn(100)
		inputs = append(inputs, randNum)
		queue.Append(randNum)
	}
	assert.Equal(t, size, queue.Len())
	assert.False(t, queue.IsEmpty())

	cnt := 0
	for !queue.IsEmpty() {
		getVal, ok := queue.Get()
		assert.True(t, ok)
		popVal, ok := queue.Pop()
		assert.True(t, ok)
		assert.Equal(t, getVal, popVal)
		assert.Equal(t, getVal, inputs[cnt])
		t.Log(fmt.Sprintf("inputVal = %d, getVal = %d, popVal = %d", inputs[cnt], getVal, popVal))
		cnt += 1
	}
	assert.Equal(t, size, cnt)
	assert.True(t, queue.IsEmpty())
}

func TestThreadSafeQueue(t *testing.T) {
	queue := NewThreadSafeQueue[int]()

	size := 100
	inputs := make([]int, 0, size)
	for i := 0; i < size; i++ {
		randNum := rand.Intn(100)
		inputs = append(inputs, randNum)
		queue.Append(randNum)
	}
	assert.Equal(t, size, queue.Len())
	assert.False(t, queue.IsEmpty())

	cnt := 0
	for !queue.IsEmpty() {
		getVal, ok := queue.Get()
		assert.True(t, ok)
		popVal, ok := queue.Pop()
		assert.True(t, ok)
		assert.Equal(t, getVal, popVal)
		assert.Equal(t, getVal, inputs[cnt])
		t.Log(fmt.Sprintf("inputVal = %d, getVal = %d, popVal = %d", inputs[cnt], getVal, popVal))
		cnt += 1
	}
	assert.Equal(t, size, cnt)
	assert.True(t, queue.IsEmpty())
}
