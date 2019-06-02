package redis_store

import (
	"lfserv/api/types"
)

func (s *RedisStore) FilteredLocks(repo, path, cursor, limit string) (locks []types.Lock, next string, err error) {
	return nil, "", nil
}
