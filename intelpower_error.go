package intelpower

import (
	"fmt"
)

const ErrCommon = "intelpower common error"

type CommonError struct {
	inner error
}

func NewCommonError(err error) CommonError {
	return CommonError{
		inner: err,
	}
}

func (err CommonError) Error() string {
	return fmt.Sprintf("%s: %s", ErrCommon, err.inner.Error())
}

func (err CommonError) Unwrap() error {
	return err.inner
}
