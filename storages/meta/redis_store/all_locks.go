package redis_store

import (
	"lfserv/store"
)

func (s *RedisStore) AllLocks() ([]store.Lock, error) {
	return nil, nil
}
