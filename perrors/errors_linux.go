// +build linux

package perrors

import "syscall"

func sysErrno(err error) Errno {
	se, ok := err.(syscall.Errno)
	if ok {
		return Errno(se)
	}
	return 0
}
