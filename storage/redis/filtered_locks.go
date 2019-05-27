package redis

import "lfserv/store/meta"

func (s *Store) FilteredLocks(repo, path, cursor, limit string) (locks []meta.Lock, next string, err error) {
	return nil, "", nil
}
