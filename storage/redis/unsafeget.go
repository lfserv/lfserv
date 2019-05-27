package redis

import "lfserv/store/meta"

func (s *Store) UnsafeGet(v *meta.RequestVars) (*meta.Object, error) {
	return &meta.Object{}, nil
}
