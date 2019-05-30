package boltdb_store

import (
	"encoding/json"
	"fmt"
	"github.com/lfserv/bolt"
	"lfserv/store"
)

func (s *BoltStore) AllLocks() ([]store.Lock, error) {
	var locks []store.Lock
	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(locksBucket)
		if bucket == nil {
			return errNoBucket
		}

		return bucket.ForEach(func(k, v []byte) error {
			var l []store.Lock
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
