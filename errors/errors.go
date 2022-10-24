package errors

import (
	"errors"
	"log"
	"os"
)

// Const may be used to convert constant strings into `error` values.
type Const string

func (errStr Const) Error() string { return string(errStr) }

// ExtractErrno extracts an [Errno] from an error, best effort.
//
// If the system-specific or Go-specific error cannot be mapped to anything, it
// will be logged and EIO will be returned.
func ExtractErrno(err error) Errno {
	for _, pair := range []struct {
		error
		Errno
	}{
		{os.ErrNotExist, ENOENT},
		{os.ErrExist, EEXIST},
		{os.ErrPermission, EACCES},
		{os.ErrInvalid, EINVAL},
	} {
		if errors.Is(err, pair.error) {
			return pair.Errno
		}
	}

	var errno Errno
	if errors.As(err, &errno) {
		return errno
	}

	// TODO: migrate the implementations of this to use 1.13 semantics.
	if e := sysErrno(err); e != 0 {
		return e
	}

	// Default case.
	//
	// TODO: give the ability to turn this off.
	log.Printf("unknown error: %v", err)
	return EIO
}
