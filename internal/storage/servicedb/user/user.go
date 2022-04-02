package user

import (
	"context"

	"github.com/jmoiron/sqlx"

	constantsmodel "github.com/sundaytycoon/buttons-api/internal/model"
	"github.com/sundaytycoon/buttons-api/internal/storage/servicedb/ent"
	"github.com/sundaytycoon/buttons-api/internal/utils/er"
)

type Storage struct {
	servicedb servicedb
}

type servicedb interface {
	EntClient() *ent.Client
	DB() *sqlx.DB
}

func New(servicedb servicedb) *Storage {
	return &Storage{
		servicedb: servicedb,
	}
}

func (a *Storage) GetUserById(ctx context.Context, id string) (*constantsmodel.User, error) {
	op := er.GetOperator()

	rows, err := a.servicedb.DB().QueryContext(ctx, "SELECT id, name, state FROM profileme_users WHERE id = ? ORDER BY id DESC", id)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	if rows.Err() != nil {
		return nil, er.WrapOp(rows.Err(), op)
	}

	var u = new(constantsmodel.User)

	for rows.Next() {
		rows.Scan(&u.ID, &u.Name, &u.State)
	}
	defer rows.Close()

	return u, nil
}
