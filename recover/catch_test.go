package recover

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCatchPanic(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		errMsg := "test_error"

		err := CatchPanic(func() error {
			return errors.New(errMsg)
		})
		assert.Equal(t, errMsg, err.Error())
	})

	t.Run("basic panic test", func(t *testing.T) {
		panicMsg := "test_error"

		err := CatchPanic(func() error {
			panic(panicMsg)
		})

		assert.Equal(t, fmt.Sprintf("unknown panic: %s", panicMsg), err.Error())
	})

	t.Run("basic err panic test", func(t *testing.T) {
		panicMsg := "test_error"

		err := CatchPanic(func() error {
			panic(errors.New(panicMsg))
		})

		assert.Equal(t, panicMsg, err.Error())
	})
}

func TestCatchPanicIgnore(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		var wg sync.WaitGroup
		errMsg := "test_error"

		wg.Add(1)
		CatchPanicIgnore(func() error {
			wg.Done()
			return errors.New(errMsg)
		})
		wg.Wait()
	})

	t.Run("basic panic test", func(t *testing.T) {
		var wg sync.WaitGroup
		panicMsg := "test_error"

		wg.Add(1)
		CatchPanicIgnore(func() error {
			wg.Done()
			panic(panicMsg)
		})
		wg.Wait()
	})

	t.Run("basic err panic test", func(t *testing.T) {
		var wg sync.WaitGroup
		panicMsg := "test_error"

		wg.Add(1)
		CatchPanicIgnore(func() error {
			wg.Done()
			panic(errors.New(panicMsg))
		})
		wg.Wait()
	})
}

func TestCatchPanicRecover(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		var wg sync.WaitGroup
		errMsg := "test_error"

		wg.Add(1)
		CatchPanicRecover(
			func() error {
				return errors.New(errMsg)
			},
			func(err error) {
				assert.Equal(t, errMsg, err.Error())
				wg.Done()
			},
		)
		wg.Wait()
	})

	t.Run("basic panic test", func(t *testing.T) {
		var wg sync.WaitGroup
		panicMsg := "test_error"

		wg.Add(1)
		CatchPanicRecover(
			func() error {
				panic(panicMsg)
			}, func(err error) {
				assert.Equal(t, fmt.Sprintf("unknown panic: %s", panicMsg), err.Error())
				wg.Done()
			},
		)
		wg.Wait()
	})

	t.Run("basic err panic test", func(t *testing.T) {
		var wg sync.WaitGroup
		panicMsg := "test_error"

		wg.Add(1)
		CatchPanicRecover(
			func() error {
				panic(errors.New(panicMsg))
			}, func(err error) {
				assert.Equal(t, panicMsg, err.Error())
				wg.Done()
			},
		)
		wg.Wait()
	})
}
