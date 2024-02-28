package recover

import (
	"errors"
	"fmt"
)

func CatchPanicRecover(runnable func() error, recover func(err error)) {
	err := CatchPanic(runnable)
	if err != nil {
		recover(err)
	}
}

func CatchPanic(runnable func() error) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			pErr, ok := rec.(error)
			if ok {
				err = pErr
				return
			}

			err = errors.New(fmt.Sprintf("unknown panic: %v", rec))
		}
	}()

	err = runnable()

	return err
}

func CatchPanicIgnore(runnable func() error) {
	_ = CatchPanic(runnable)
}
