package user

//go:generate mockgen -source=interface.go -destination interface_mock.go -package=user

import (
	"context"

	"github.com/sundaytycoon/profile.me-server/internal/constants/model"
)

type userService interface {
	Get(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, name, state string) (*model.User, error)
}
