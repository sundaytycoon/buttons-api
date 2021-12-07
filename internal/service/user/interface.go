package user

//go:generate mockgen -source=interface.go -destination interface_mock.go -package=user

import (
	"context"

	"github.com/sundaytycoon/profile.me-server/internal/constants/model"
)

type userRepository interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	Save(ctx context.Context, u *model.User) (*model.User, error)
}
