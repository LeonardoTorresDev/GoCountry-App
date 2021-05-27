package errors

import (
	"github.com/pkg/errors"
)

type fileWritingFailed struct {
	error
}

func WrapFileWritingFailed(err error, format string, args ...interface{}) error {
	return &fileWritingFailed{errors.Wrapf(err, format, args...)}
}

func NewFileWritingFailed(format string, args ...interface{}) error {
	return &fileWritingFailed{errors.Errorf(format, args...)}
}

func IsFileWritingFailed(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*fileWritingFailed)
	return ok
}
