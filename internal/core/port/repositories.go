package port

import (
	"context"

	"github.com/sundaytycoon/profile.me-server/internal/core/domain"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*domain.User, error)
	Save(ctx context.Context, u domain.User) (*domain.User, error)
}
