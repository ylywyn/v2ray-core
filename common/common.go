// Package common contains common utilities that are shared among other packages.
// See each sub-package for detail.
package common

import (
	"errors"
)

var (
	ErrObjectReleased   = errors.New("Object already released.")
	ErrBadConfiguration = errors.New("Bad configuration.")
	ErrObjectNotFound   = errors.New("Object not found.")
	ErrDuplicatedName   = errors.New("Duplicated name.")
)

// Must panics if err is not nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
