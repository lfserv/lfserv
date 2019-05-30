package boltdb_store

import "github.com/lfserv/bolt"

func (s *BoltStore) AddUser(user, pass string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(usersBucket)
		if bucket == nil {
			return errNoBucket
		}

		err := bucket.Put([]byte(user), []byte(pass))
		if err != nil {
			return err
		}
		return nil
	})
}
