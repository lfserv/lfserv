package boltdb_store

import "github.com/lfserv/bolt"

func (s *BoltStore) DeleteUser(user string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(usersBucket)
		if bucket == nil {
			return errNoBucket
		}

		err := bucket.Delete([]byte(user))
		return err
	})
}
