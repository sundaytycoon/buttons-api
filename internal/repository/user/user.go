package user

import (
	"context"

	"github.com/sundaytycoon/buttons-api/internal/constants/model"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

// user repository  직접 구현 한 곳
type Repository struct {
	mysqlClient mysqlClient
	userStore   userStore
}

func New(m mysqlClient, u userStore) *Repository {
	return &Repository{
		mysqlClient: m,
		userStore:   u,
	}
}

func (r *Repository) GetUser(ctx context.Context, id string) (*model.User, error) {
	op := er.GetOperator()

	conn, err := r.mysqlClient.Conn(ctx)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	u, err := r.userStore.GetUser(ctx, conn, id)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	return u, nil
}

func (r *Repository) Save(ctx context.Context, u *model.User) (*model.User, error) {
	return nil, nil
}
