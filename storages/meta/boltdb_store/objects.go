package boltdb_store

import (
	"bytes"
	"encoding/gob"
	"github.com/lfserv/bolt"
	"lfserv/api/types"
)

func (s *BoltStore) Objects() ([]*types.MetaObject, error) {
	var objects []*types.MetaObject

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(objectsBucket)
		if bucket == nil {
			return errNoBucket
		}

		return bucket.ForEach(func(k, v []byte) error {
			var meta types.MetaObject
			dec := gob.NewDecoder(bytes.NewBuffer(v))
			err := dec.Decode(&meta)
			if err != nil {
				return err
			}
			objects = append(objects, &meta)
			return nil
		})
	})

	return objects, err
}
