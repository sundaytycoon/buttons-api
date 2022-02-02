package user

import (
	"context"

	"github.com/sundaytycoon/buttons-api/internal/constants/model"
)

// user repository  직접 구현 한 곳
type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r Repository) FindUserDataByOAuthTokens(ctx context.Context, refreshToken string) *model.UserData {
	return nil
}
