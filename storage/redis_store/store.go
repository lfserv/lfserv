package redis_store

// RedisStore implements a metadata storage. It stores user credentials and Meta information
// for objects. The storage is handled by boltdb.
type RedisStore struct {
	Uri string
}

func (s *RedisStore) Init() error {
	return nil
}
