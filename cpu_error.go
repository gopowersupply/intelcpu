package intelpower

import (
	"fmt"
)

const ErrCPU = "cpu error"

type CPUError struct {
	inner error
}

func NewCPUError(err error) CPUError {
	return CPUError{
		inner: err,
	}
}

func (err CPUError) Error() string {
	return fmt.Sprintf("%s: %s", ErrCPU, err.inner.Error())
}

func (err CPUError) Unwrap() error {
	return err.inner
}
