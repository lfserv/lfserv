package content

import (
	"io"
)

type Store interface {
	Get(oid string, fromByte int64) (io.ReadCloser, error)
	Put(oid string, size int64, r io.Reader) error
	Exists(oid string) bool
}
