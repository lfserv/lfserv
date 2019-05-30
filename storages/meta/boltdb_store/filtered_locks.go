package boltdb_store

import (
	"fmt"
	"lfserv/store"
	"math"
	"strconv"
)

func (s *BoltStore) FilteredLocks(repo, path, cursor, limit string) (locks []store.Lock, next string, err error) {
	locks, err = s.Locks(repo)
	if err != nil {
		return
	}

	if cursor != "" {
		lastSeen := -1
		for i, l := range locks {
			if l.Id == cursor {
				lastSeen = i
				break
			}
		}

		if lastSeen > -1 {
			locks = locks[lastSeen:]
		} else {
			err = fmt.Errorf("cursor (%s) not found", cursor)
			return
		}
	}

	if path != "" {
		var filtered []store.Lock
		for _, l := range locks {
			if l.Path == path {
				filtered = append(filtered, l)
			}
		}

		locks = filtered
	}

	if limit != "" {
		var size int
		size, err = strconv.Atoi(limit)
		if err != nil || size < 0 {
			locks = make([]store.Lock, 0)
			err = fmt.Errorf("invalid limit amount: %s", limit)
			return
		}

		size = int(math.Min(float64(size), float64(len(locks))))
		if size+1 < len(locks) {
			next = locks[size].Id
		}
		locks = locks[:size]
	}

	return locks, next, nil
}
