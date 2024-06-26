//go:build linux
// +build linux

package errors

import (
	"errors"
	"syscall"
)

func sysErrno(err error) Errno {
	var systemErr syscall.Errno
	if errors.As(err, &systemErr) {
		return Errno(systemErr)
	}
	return 0
}
