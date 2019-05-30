package boltdb_store

import (
	"encoding/json"
	"github.com/lfserv/bolt"
	"lfserv/store"
)

func (s *BoltStore) DeleteLock(repo, user, id string, force bool) (*store.Lock, error) {
	var deleted *store.Lock
	err := s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(locksBucket)
		if bucket == nil {
			return errNoBucket
		}

		var locks []store.Lock
		data := bucket.Get([]byte(repo))
		if data != nil {
			if err := json.Unmarshal(data, &locks); err != nil {
				return err
			}
		}
		newLocks := make([]store.Lock, 0, len(locks))

		var lock store.Lock
		for _, l := range locks {
			if l.Id == id {
				if l.Owner.Name != user && !force {
					return errNotOwner
				}
				lock = l
			} else if len(l.Id) > 0 {
				newLocks = append(newLocks, l)
			}
		}
		if lock.Id == "" {
			return nil
		}
		deleted = &lock

		if len(newLocks) == 0 {
			return bucket.Delete([]byte(repo))
		}

		data, err := json.Marshal(&newLocks)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(repo), data)
	})
	return deleted, err
}
