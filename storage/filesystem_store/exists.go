package filesystem_store

import (
	"os"
	"path/filepath"
)

// Exists returns true if the object exists in the content store.
func (s *FileSystemStore) Exists(oid string) bool {
	path := filepath.Join(s.basePath, transformKey(oid))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
