package meta

import (
	"time"
)

type Lock struct {
	Id       string    `json:"id"`
	Path     string    `json:"path"`
	Owner    User      `json:"owner"`
	LockedAt time.Time `json:"locked_at"`
}

type LockRequest struct {
	Path string `json:"path"`
}

type LockResponse struct {
	Lock    *Lock  `json:"lock"`
	Message string `json:"message,omitempty"`
}

type LockList struct {
	Locks      []Lock `json:"locks"`
	NextCursor string `json:"next_cursor,omitempty"`
	Message    string `json:"message,omitempty"`
}

type LocksByCreatedAt []Lock

func (c LocksByCreatedAt) Len() int           { return len(c) }
func (c LocksByCreatedAt) Less(i, j int) bool { return c[i].LockedAt.Before(c[j].LockedAt) }
func (c LocksByCreatedAt) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
