package user

// user repository  직접 구현 한 곳
type Repository struct {
}

func New() *Repository {
	return &Repository{}
}
