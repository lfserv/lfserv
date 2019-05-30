package boltdb_store

import (
	"lfserv/store"
)

func (s *BoltStore) Get(v *store.RequestVars) (*store.MetaObject, error) {
	return s.UnsafeGet(v)
}
