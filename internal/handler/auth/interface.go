package auth

//go:generate mockgen -source=interface.go -destination interface_mock.go -package=user

import (
	"context"

	"github.com/sundaytycoon/buttons-api/internal/constants/model"
)

type sessionService interface {
	GetRedirectURL(ctx context.Context, id string) (*model.User, error)
}
