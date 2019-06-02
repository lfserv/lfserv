package boltdb_store

import (
	"encoding/json"
	"fmt"
	"github.com/lfserv/bolt"
	"lfserv/api/types"
)

func (s *BoltStore) AllLocks() ([]types.Lock, error) {
	var locks []types.Lock
	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(locksBucket)
		if bucket == nil {
			return errNoBucket
		}

		return bucket.ForEach(func(k, v []byte) error {
			var l []types.Lock
			if err := json.Unmarshal(v, &l); err != nil {
				return err
			}
			for _, lv := range l {
				lv.Path = fmt.Sprintf("%s:%s", k, lv.Path)
				locks = append(locks, lv)
			}
			return nil
		})
	})
	return locks, err
}
