package redis_store

import (
	"lfserv/store"
)

func (s *RedisStore) DeleteLock(repo, user, id string, force bool) (*store.Lock, error) {
	return nil, nil
}
