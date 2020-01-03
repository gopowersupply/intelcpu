package intelcpu

import (
	"errors"
	"fmt"
)

// errCommon - General error text for all errors from this package
const errCommon = "intelcpu common error"

// CPUError - Error object for all error in this package
type CPUError struct {
	inner error
}

// NewCPUError - Wrapper for errors in this package
// You can use this function to make your own CPU error
func NewCPUError(err error) CPUError {
	return CPUError{
		inner: err,
	}
}

func (err CPUError) Error() string {
	return fmt.Sprintf("%s: %s", errCommon, err.inner.Error())
}

// Unwrap - Returns inner error
func (err CPUError) Unwrap() error {
	return err.inner
}

// IsCPUError - Checks that error is error from this package
func IsCPUError(err error) bool {
	return errors.As(err, &CPUError{})
}
