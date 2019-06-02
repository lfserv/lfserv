package store

import (
	"lfserv/api/types"
)

type MetaStore interface {
	Init() error
	Get(v *types.RequestVars) (*types.MetaObject, error)
	UnsafeGet(v *types.RequestVars) (*types.MetaObject, error)
	Put(v *types.RequestVars) (*types.MetaObject, error)
	Delete(v *types.RequestVars) error
	AddLocks(repo string, l ...types.Lock) error
	Locks(repo string) ([]types.Lock, error)
	AllLocks() ([]types.Lock, error)
	FilteredLocks(repo, path, cursor, limit string) (locks []types.Lock, next string, err error)
	DeleteLock(repo, user, id string, force bool) (*types.Lock, error)
	Close() error
	AddUser(user, pass string) error
	DeleteUser(user string) error
	Users() ([]*types.User, error)
	Objects() ([]*types.MetaObject, error)
	Authenticate(user, password string) (string, bool)
}
