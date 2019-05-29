package store

type MetaStore interface {
	Init() error
	Get(v *RequestVars) (*MetaObject, error)
	UnsafeGet(v *RequestVars) (*MetaObject, error)
	Put(v *RequestVars) (*MetaObject, error)
	Delete(v *RequestVars) error
	AddLocks(repo string, l ...Lock) error
	Locks(repo string) ([]Lock, error)
	AllLocks() ([]Lock, error)
	FilteredLocks(repo, path, cursor, limit string) (locks []Lock, next string, err error)
	DeleteLock(repo, user, id string, force bool) (*Lock, error)
	Close() error
	AddUser(user, pass string) error
	DeleteUser(user string) error
	Users() ([]*User, error)
	Objects() ([]*MetaObject, error)
	Authenticate(user, password string) (string, bool)
}
