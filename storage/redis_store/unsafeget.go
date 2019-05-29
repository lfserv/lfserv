package redis_store

import (
	"lfserv/store"
)

func (s *RedisStore) UnsafeGet(v *store.RequestVars) (*store.MetaObject, error) {
	return &store.MetaObject{}, nil
}
