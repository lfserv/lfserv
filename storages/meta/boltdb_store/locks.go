package boltdb_store

import (
	"encoding/json"
	"github.com/lfserv/bolt"
	"lfserv/api/types"
)

func (s *BoltStore) Locks(repo string) ([]types.Lock, error) {
	var locks []types.Lock
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
