package boltdb_store

func (s *BoltStore) Close() error {
	return s.db.Close()
}
