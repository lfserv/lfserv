package redis_store

import (
	"lfserv/store"
)

func (s *RedisStore) Locks(repo string) ([]store.Lock, error) {
	return nil, nil
}
