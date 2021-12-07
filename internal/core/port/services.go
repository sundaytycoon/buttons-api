package port

import (
	"context"

	"github.com/sundaytycoon/profile.me-server/internal/core/domain"
)

type UserService interface {
	Get(ctx context.Context, id string) (*domain.User, error)
	Create(ctx context.Context, name string, state string) (*domain.User, error)
}
