package redis_store

import (
	"lfserv/store"
)

func (s *RedisStore) FilteredLocks(repo, path, cursor, limit string) (locks []store.Lock, next string, err error) {
	return nil, "", nil
}
