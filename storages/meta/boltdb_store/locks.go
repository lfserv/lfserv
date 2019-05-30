package boltdb_store

import (
	"encoding/json"
	"github.com/lfserv/bolt"
	"lfserv/store"
)

func (s *BoltStore) Locks(repo string) ([]store.Lock, error) {
	var locks []store.Lock
	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(locksBucket)
		if bucket == nil {
			return errNoBucket
		}

		data := bucket.Get([]byte(repo))
		if data != nil {
			if err := json.Unmarshal(data, &locks); err != nil {
				return err
			}
		}
		return nil
	})
	return locks, err
}
