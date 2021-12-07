package user

import (
	"context"

	"github.com/sundaytycoon/profile.me-server/internal/core/domain"
	"github.com/sundaytycoon/profile.me-server/internal/core/port"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
)

// user service를 직접 구현 한 곳
type service struct {
	userRepository port.UserRepository
}

func New(userRepo port.UserRepository) *service {
	return &service{
		userRepository: userRepo,
	}
}

func (s *service) Get(ctx context.Context, id string) (*domain.User, error) {
	op := er.GetOperator()

	u, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return u, nil
}

func (s *service) Create(ctx context.Context, name, state string) (*domain.User, error) {
	op := er.GetOperator()

	u, err := s.userRepository.Save(ctx, &domain.User{Name: name, State: state})
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return u, nil
}
