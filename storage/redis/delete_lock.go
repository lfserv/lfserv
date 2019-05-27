package redis

import "lfserv/store/meta"

func (s *Store) DeleteLock(repo, user, id string, force bool) (*meta.Lock, error) {
	return nil, nil
}
