package boltdb_store

import (
	"github.com/lfserv/bolt"
	"lfserv/api/types"
)

func (s *BoltStore) Users() ([]*types.User, error) {
	var users []*types.User

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(usersBucket)
		if bucket == nil {
			return errNoBucket
		}

		return bucket.ForEach(func(k, v []byte) error {
			users = append(users, &types.User{Name: string(k)})
			return nil
		})
	})

	return users, err
}
