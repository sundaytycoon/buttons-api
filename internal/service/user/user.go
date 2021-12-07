package user

import (
	"context"

	"github.com/sundaytycoon/profile.me-server/internal/constants/model"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
)

// user service를 직접 구현 한 곳
type Service struct {
	userRepository userRepository
}

func New(ur userRepository) *Service {
	return &Service{
		userRepository: ur,
	}
}

func (s *Service) Get(ctx context.Context, id string) (*model.User, error) {
	op := er.GetOperator()

	u, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return u, nil
}

func (s *Service) Create(ctx context.Context, name, state string) (*model.User, error) {
	op := er.GetOperator()

	u, err := s.userRepository.Save(ctx, &model.User{Name: name, State: state})
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return u, nil
}
