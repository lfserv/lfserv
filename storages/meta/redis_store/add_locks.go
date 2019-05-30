package redis_store

import (
	"lfserv/store"
)

func (s *RedisStore) AddLocks(repo string, l ...store.Lock) error {
	return nil
}
