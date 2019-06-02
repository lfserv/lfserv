package boltdb_store

import (
	"bytes"
	"encoding/gob"
	"github.com/lfserv/bolt"
	"lfserv/api/types"
)

func (s *BoltStore) Put(v *types.RequestVars) (*types.MetaObject, error) {
	// Check if it exists first
	if meta, err := s.Get(v); err == nil {
		meta.Existing = true
		return meta, nil
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	meta := types.MetaObject{Oid: v.Oid, Size: v.Size}
	err := enc.Encode(meta)
	if err != nil {
		return nil, err
	}

	err = s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(objectsBucket)
		if bucket == nil {
			return errNoBucket
		}

		err = bucket.Put([]byte(v.Oid), buf.Bytes())
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &meta, nil
}
