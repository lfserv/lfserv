package filesystem_store

import (
	"errors"
	"os"
)

const (
	storeRoot = "LFS_CONTENT_FS_ROOT"
)

var (
	errHashMismatch = errors.New("content hash does not match OID")
	errSizeMismatch = errors.New("content size does not match")
)

// FileSystemStore provides a simple file system based storage.
type FileSystemStore struct {
	basePath string
}

// Init create FileSystemStore at the base directory.
func (s *FileSystemStore) Init() error {
	s.basePath = os.Getenv(storeRoot)
	if s.basePath == "" {
		s.basePath = "content-store"
	}
	return os.MkdirAll(s.basePath, 0750)
}
