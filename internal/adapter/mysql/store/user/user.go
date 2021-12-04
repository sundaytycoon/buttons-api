package user

import (
	"context"

	"github.com/sundaytycoon/profile.me-server/internal/core/domain"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
	execdbconn "github.com/sundaytycoon/profile.me-server/pkg/execdbconn"
)

type Store struct{}

func New() (*Store, error) {
	return &Store{}, nil
}

func (a *Store) GetUser(ctx context.Context, tx execdbconn.ContextExecutor, id string) (*domain.User, error) {
	op := er.GetOperator()

	rows, err := tx.QueryContext(ctx, "SELECT id, name, state FROM profileme_users WHERE id = ? ORDER BY id DESC", id)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	if rows.Err() != nil {
		return nil, er.WrapOp(rows.Err(), op)
	}

	var u *domain.User

	for rows.Next() {
		rows.Scan(&u.ID, &u.Name, &u.State)
	}
	defer rows.Close()

	return u, nil
}
