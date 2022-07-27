package templatefs

import (
	"github.com/hugelgupf/p9/errors"
	"github.com/hugelgupf/p9/p9"
)

// NotSymlinkFile denies Readlink with EINVAL.
//
// EINVAL is returned by readlink(2) when the file is not a symlink.
type NotSymlinkFile struct{}

// Readlink implements p9.File.Readlink.
func (NotSymlinkFile) Readlink() (string, error) {
	return "", errors.EINVAL
}

// NotDirectoryFile denies any directory operations with ENOTDIR.
//
// Those operations are Create, Mkdir, Symlink, Link, Mknod, RenameAt,
// UnlinkAt, and Readdir.
type NotDirectoryFile struct{}

// Create implements p9.File.Create.
func (NotDirectoryFile) Create(name string, mode p9.OpenFlags, permissions p9.FileMode, _ p9.UID, _ p9.GID) (p9.File, p9.QID, uint32, error) {
	return nil, p9.QID{}, 0, errors.ENOTDIR
}

// Mkdir implements p9.File.Mkdir.
func (NotDirectoryFile) Mkdir(name string, permissions p9.FileMode, _ p9.UID, _ p9.GID) (p9.QID, error) {
	return p9.QID{}, errors.ENOTDIR
}

// Symlink implements p9.File.Symlink.
func (NotDirectoryFile) Symlink(oldname string, newname string, _ p9.UID, _ p9.GID) (p9.QID, error) {
	return p9.QID{}, errors.ENOTDIR
}

// Link implements p9.File.Link.
func (NotDirectoryFile) Link(target p9.File, newname string) error {
	return errors.ENOTDIR
}

// Mknod implements p9.File.Mknod.
func (NotDirectoryFile) Mknod(name string, mode p9.FileMode, major uint32, minor uint32, _ p9.UID, _ p9.GID) (p9.QID, error) {
	return p9.QID{}, errors.ENOTDIR
}

// RenameAt implements p9.File.RenameAt.
func (NotDirectoryFile) RenameAt(oldname string, newdir p9.File, newname string) error {
	return errors.ENOTDIR
}

// UnlinkAt implements p9.File.UnlinkAt.
func (NotDirectoryFile) UnlinkAt(name string, flags uint32) error {
	return errors.ENOTDIR
}

// Readdir implements p9.File.Readdir.
func (NotDirectoryFile) Readdir(offset uint64, count uint32) (p9.Dirents, error) {
	return nil, errors.ENOTDIR
}

// ReadOnlyFile returns EROFS for FSync, SetAttr, Remove, Rename, WriteAt, and nil for Flush.
type ReadOnlyFile struct{}

// FSync implements p9.File.FSync.
func (ReadOnlyFile) FSync() error {
	return errors.EROFS
}

// SetAttr implements p9.File.SetAttr.
func (ReadOnlyFile) SetAttr(valid p9.SetAttrMask, attr p9.SetAttr) error {
	return errors.EROFS
}

// Remove implements p9.File.Remove.
func (ReadOnlyFile) Remove() error {
	return errors.EROFS
}

// Rename implements p9.File.Rename.
func (ReadOnlyFile) Rename(directory p9.File, name string) error {
	return errors.EROFS
}

// WriteAt implements p9.File.WriteAt.
func (ReadOnlyFile) WriteAt(p []byte, offset int64) (int, error) {
	return 0, errors.EROFS
}

// Flush implements p9.File.Flush.
func (ReadOnlyFile) Flush() error {
	return nil
}

// ReadOnlyDir denies any directory and file operations with EROFS
//
// Those operations are Create, Mkdir, Symlink, Link, Mknod, RenameAt,
// UnlinkAt, Readdir, Rename, SetAttr, FSync, and Remove.
type ReadOnlyDir struct{}

// Create implements p9.File.Create.
func (ReadOnlyDir) Create(name string, mode p9.OpenFlags, permissions p9.FileMode, _ p9.UID, _ p9.GID) (p9.File, p9.QID, uint32, error) {
	return nil, p9.QID{}, 0, errors.EROFS
}

// Mkdir implements p9.File.Mkdir.
func (ReadOnlyDir) Mkdir(name string, permissions p9.FileMode, _ p9.UID, _ p9.GID) (p9.QID, error) {
	return p9.QID{}, errors.EROFS
}

// Symlink implements p9.File.Symlink.
func (ReadOnlyDir) Symlink(oldname string, newname string, _ p9.UID, _ p9.GID) (p9.QID, error) {
	return p9.QID{}, errors.EROFS
}

// Link implements p9.File.Link.
func (ReadOnlyDir) Link(target p9.File, newname string) error {
	return errors.EROFS
}

// Mknod implements p9.File.Mknod.
func (ReadOnlyDir) Mknod(name string, mode p9.FileMode, major uint32, minor uint32, _ p9.UID, _ p9.GID) (p9.QID, error) {
	return p9.QID{}, errors.EROFS
}

// RenameAt implements p9.File.RenameAt.
func (ReadOnlyDir) RenameAt(oldname string, newdir p9.File, newname string) error {
	return errors.EROFS
}

// UnlinkAt implements p9.File.UnlinkAt.
func (ReadOnlyDir) UnlinkAt(name string, flags uint32) error {
	return errors.EROFS
}

// Readdir implements p9.File.Readdir.
func (ReadOnlyDir) Readdir(offset uint64, count uint32) (p9.Dirents, error) {
	return nil, errors.EROFS
}

// FSync implements p9.File.FSync.
func (ReadOnlyDir) FSync() error {
	return errors.EROFS
}

// SetAttr implements p9.File.SetAttr.
func (ReadOnlyDir) SetAttr(valid p9.SetAttrMask, attr p9.SetAttr) error {
	return errors.EROFS
}

// Remove implements p9.File.Remove.
func (ReadOnlyDir) Remove() error {
	return errors.EROFS
}

// Rename implements p9.File.Rename.
func (ReadOnlyDir) Rename(directory p9.File, name string) error {
	return errors.EROFS
}

// IsDir returns EISDIR for ReadAt and WriteAt.
type IsDir struct{}

// WriteAt implements p9.File.WriteAt.
func (IsDir) WriteAt(p []byte, offset int64) (int, error) {
	return 0, errors.EISDIR
}

// ReadAt implements p9.File.ReadAt.
func (IsDir) ReadAt(p []byte, offset int64) (int, error) {
	return 0, errors.EISDIR
}
