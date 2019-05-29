package filesystem_store

import (
	"io"
	"os"
	"path/filepath"
)

// Get takes a Meta object and retreives the content from the store, returning
// it as an io.ReaderCloser. If fromByte > 0, the reader starts from that byte
func (s *FileSystemStore) Get(oid string, fromByte int64) (io.ReadCloser, error) {
	path := filepath.Join(s.basePath, transformKey(oid))

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if fromByte > 0 {
		_, err = f.Seek(fromByte, io.SeekCurrent)
	}
	return f, err
}
