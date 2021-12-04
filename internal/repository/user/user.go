package user

import (
	"context"

	adaptermysql "github.com/sundaytycoon/profile.me-server/internal/adapter/mysql"
	storeuser "github.com/sundaytycoon/profile.me-server/internal/adapter/mysql/store/user"
	"github.com/sundaytycoon/profile.me-server/internal/core/domain"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
	"go.uber.org/dig"
)

// user repository  직접 구현 한 곳
type Repository struct {
	mysqlAdapter *adaptermysql.Adapter
	userStore    userStore
}

func New(params struct {
	dig.In
	ServiceDBAdapter *adaptermysql.Adapter
	UserStore        *storeuser.Store
}) (*Repository, error) {
	return &Repository{
		mysqlAdapter: params.ServiceDBAdapter,
		userStore:    params.UserStore,
	}, nil
}

func (r *Repository) GetUser(ctx context.Context, id string) (*domain.User, error) {
	op := er.GetOperator()

	conn, err := r.mysqlAdapter.Conn(ctx)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	u, err := r.userStore.GetUser(ctx, conn, id)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	return u, nil
}
