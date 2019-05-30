package boltdb_store

import (
	"github.com/lfserv/bolt"
	"lfserv/store"
)

func (s *BoltStore) Users() ([]*store.User, error) {
	var users []*store.User

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(usersBucket)
		if bucket == nil {
			return errNoBucket
		}

		return bucket.ForEach(func(k, v []byte) error {
			users = append(users, &store.User{Name: string(k)})
			return nil
		})
	})

	return users, err
}
