package boltdb_store

import (
	"github.com/lfserv/bolt"
	"lfserv/api/types"
)

func (s *BoltStore) Delete(v *types.RequestVars) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(objectsBucket)
		if bucket == nil {
			return errNoBucket
		}

		err := bucket.Delete([]byte(v.Oid))
		if err != nil {
			return err
		}

		return nil
	})
}
