package boltdb_store

import (
	"bytes"
	"encoding/gob"
	"github.com/lfserv/bolt"
	"lfserv/store"
)

func (s *BoltStore) Objects() ([]*store.MetaObject, error) {
	var objects []*store.MetaObject

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(objectsBucket)
		if bucket == nil {
			return errNoBucket
		}

		return bucket.ForEach(func(k, v []byte) error {
			var meta store.MetaObject
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
