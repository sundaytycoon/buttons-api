package auth

type storage struct {
	servicedb servicedb
}

type servicedb interface {
}

func New(servicedb servicedb) *storage {
	return &storage{
		servicedb: servicedb,
	}
}
