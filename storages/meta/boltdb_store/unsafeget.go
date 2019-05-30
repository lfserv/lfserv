package boltdb_store

import (
	"bytes"
	"encoding/gob"
	"github.com/lfserv/bolt"
	"lfserv/store"
)

func (s *BoltStore) UnsafeGet(v *store.RequestVars) (*store.MetaObject, error) {
	var meta store.MetaObject

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(objectsBucket)
		if bucket == nil {
			return errNoBucket
		}

		value := bucket.Get([]byte(v.Oid))
		if len(value) == 0 {
			return errObjectNotFound
		}

		dec := gob.NewDecoder(bytes.NewBuffer(value))
		return dec.Decode(&meta)
	})

	if err != nil {
		return nil, err
	}

	return &meta, nil
}
