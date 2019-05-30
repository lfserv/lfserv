package redis_store

import (
	"lfserv/store"
)

func (s *RedisStore) Put(v *store.RequestVars) (*store.MetaObject, error) {
	return &store.MetaObject{}, nil
}
