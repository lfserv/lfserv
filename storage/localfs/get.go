package localfs

import "io"

func (s *Store) Get(oid string, fromByte int64) (io.ReadCloser, error) {
	return nil, nil
}
