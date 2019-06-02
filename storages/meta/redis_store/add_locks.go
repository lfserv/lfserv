package redis_store

import (
	"lfserv/api/types"
)

func (s *RedisStore) AddLocks(repo string, l ...types.Lock) error {
	return nil
}
