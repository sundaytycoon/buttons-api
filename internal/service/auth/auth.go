package auth

import (
	"context"
	"fmt"

	"github.com/sundaytycoon/buttons-api/internal/constants/model"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

type authRepository interface {
	GetOAuthRedirectURL(provider, fromHost string) (string, error)
	GetUserInfoFromProvider(ctx context.Context, provider, code string) (*model.UserToken, error)
}

// Service user service를 직접 구현 한 곳
type Service struct {
	authRepository authRepository
}

func New(ar authRepository) *Service {
	return &Service{
		authRepository: ar,
	}
}

func (s *Service) GetWebOAuthRedirectURL(fromHost, provider string) (string, error) {
	op := er.GetOperator()
	v, err := s.authRepository.GetOAuthRedirectURL(fromHost, provider)
	if err != nil {
		return "", er.WrapOp(err, op)
	}
	return v, nil
}

type GetWebCallbackOut struct {
}

func (s *Service) GetWebCallback(ctx context.Context, provider, code, state string) (string, string, error) {
	op := er.GetOperator()

	// 1. get accessToken // refreshToken // expiry // auth_type // email
	ut, err := s.authRepository.GetUserInfoFromProvider(ctx, provider, code)
	if err != nil {
		return "", "", er.WrapOp(err, op)
	}
	fmt.Println(ut)
	// 2. get userdata

	// 2-1 get userdata using refresh_token

	// 2-2 get userdata using email

	// 3. 회원 판단

	// 4. 회원가입

	// 4-1. [] insert user[users / users_oauth_provider / users_device]

	// 4-2. []

	// 4-2. []

	// 5. 회원가입 도중

	// 5-1. 진행중이던 과정부터 시작할 수 있도록 도움

	// 6. 로그인

	// 7. session data 업데ㅇ이트

	return state, "", nil
}
