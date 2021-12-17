package user

//go:generate mockgen -source=interface.go -destination interface_mock.go -package=user

import (
	"context"
	"database/sql"

	"github.com/sundaytycoon/buttons-api/edge/mysql"
	"github.com/sundaytycoon/buttons-api/internal/constants/model"
)

type userStore interface {
	GetUser(ctx context.Context, tx mysql.ContextExecutor, id string) (*model.User, error)
}

type mysqlClient interface {
	Conn(ctx context.Context) (*sql.Conn, error)
	Close() error
}
