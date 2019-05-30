package boltdb_store

import (
	"errors"
	"github.com/lfserv/bolt"
	"os"
	"time"
)

const (
	boltDbPath = "LFS_META_BOLT_DB"
)

var (
	usersBucket   = []byte("users")
	objectsBucket = []byte("objects")
	locksBucket   = []byte("locks")

	errNoBucket       = errors.New("bucket not found")
	errObjectNotFound = errors.New("object not found")
	errNotOwner       = errors.New("attempt to delete other user's lock")
)

type BoltStore struct {
	dbFile string
	db     *bolt.DB
}

func (s *BoltStore) Init() error {
	s.dbFile = os.Getenv(boltDbPath)
	if s.dbFile == "" {
		s.dbFile = "meta.bolt"
	}
	db, err := bolt.Open(s.dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	s.db = db

	return db.Update(func(tx *bolt.Tx) error {
		buckets := [][]byte{usersBucket, objectsBucket, locksBucket}
		for _, bucket := range buckets {
			if _, err := tx.CreateBucketIfNotExists(bucket); err != nil {
				return err
			}
		}
		return nil
	})
}
