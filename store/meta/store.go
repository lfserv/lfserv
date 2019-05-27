package meta

type Store interface {
	Get(v *RequestVars) (*Object, error)
	UnsafeGet(v *RequestVars) (*Object, error)
	Put(v *RequestVars) (*Object, error)
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
	Objects() ([]*Object, error)
	Authenticate(user, password string) (string, bool)
}
