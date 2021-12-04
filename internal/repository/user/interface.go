package user

//go:generate mockgen -source=interface.go -destination interface_mock.go -package=user

import (
	"context"

	"github.com/sundaytycoon/profile.me-server/internal/core/domain"
	"github.com/sundaytycoon/profile.me-server/pkg/execdbconn"
)

type userStore interface {
	GetUser(ctx context.Context, tx execdbconn.ContextExecutor, id string) (*domain.User, error)
}
