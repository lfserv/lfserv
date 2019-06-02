package boltdb_store

import (
	"lfserv/api/types"
)

func (s *BoltStore) Get(v *types.RequestVars) (*types.MetaObject, error) {
	return s.UnsafeGet(v)
}
