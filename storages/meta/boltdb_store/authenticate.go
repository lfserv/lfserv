package boltdb_store

import "github.com/lfserv/bolt"

func (s *BoltStore) Authenticate(user, password string) (string, bool) {
	// check admin
	if len(user) > 0 && len(password) > 0 {
		if checkBasicAuth(user, password, true) {
			return user, true
		}
	}

	value := ""

	_ = s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(usersBucket)
		if bucket == nil {
			return errNoBucket
		}
		value = string(bucket.Get([]byte(user)))
		return nil
	})

	return user, value != "" && value == password
}

func checkBasicAuth(user, pass string, ok bool) bool {
	if !ok {
		return false
	}
	//if user != Config.AdminUser || pass != Config.AdminPass {
	//	return false
	//}
	return true
}
