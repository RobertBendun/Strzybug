package utils

import (
	"fmt"
	"runtime"
)

func Wrap(err error) error {
	if err == nil {
		return nil
	}

	fpcs := make([]uintptr, 1)
	if runtime.Callers(2, fpcs) == 0 {
		panic("no caller for wrap() function")
	}

	caller := runtime.FuncForPC(fpcs[0]-1)
	if caller == nil {
		panic("wrap() caller is nil")
	}

	return fmt.Errorf("%s: %w", caller.Name(), err)
}
