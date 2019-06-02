package boltdb_store

import (
	"encoding/json"
	"github.com/lfserv/bolt"
	"lfserv/api/types"
	"sort"
)

func (s *BoltStore) AddLocks(repo string, l ...types.Lock) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(locksBucket)
		if bucket == nil {
			return errNoBucket
		}

		var locks []types.Lock
		data := bucket.Get([]byte(repo))
		if data != nil {
			if err := json.Unmarshal(data, &locks); err != nil {
				return err
			}
		}
		locks = append(locks, l...)
		sort.Sort(types.LocksByCreatedAt(locks))
		data, err := json.Marshal(&locks)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(repo), data)
	})
}
