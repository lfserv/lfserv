package redis_store

import (
	"lfserv/api/types"
)

func (s *RedisStore) DeleteLock(repo, user, id string, force bool) (*types.Lock, error) {
	return nil, nil
}
