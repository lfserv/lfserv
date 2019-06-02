package redis_store

import (
	"lfserv/api/types"
)

func (s *RedisStore) UnsafeGet(v *types.RequestVars) (*types.MetaObject, error) {
	return &types.MetaObject{}, nil
}
