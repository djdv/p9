package xattr

import "github.com/djdv/p9/errors"

func List(p string) ([]string, error) {
	return nil, errors.ENOSYS
}

func Get(p string, attr string) ([]byte, error) {
	return nil, errors.ENOSYS
}
