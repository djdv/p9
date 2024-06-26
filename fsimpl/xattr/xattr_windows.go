package xattr

import "github.com/hugelgupf/p9/errors"

func List(p string) ([]string, error) {
	return nil, errors.ENOSYS
}

func Get(p string, attr string) ([]byte, error) {
	return nil, errors.ENOSYS
}
