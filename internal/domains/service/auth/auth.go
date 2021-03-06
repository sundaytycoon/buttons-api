package auth

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sundaytycoon/buttons-api/internal/model"
	"github.com/sundaytycoon/buttons-api/internal/utils/er"
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

func (s *Service) GetWebOAuthRedirectURL(_ context.Context, fromHost, provider string) (string, error) {
	op := er.GetOperator()
	v, err := s.authRepository.GetOAuthRedirectURL(fromHost, provider)
	if err != nil {
		return "", er.WrapOp(err, op)
	}
	return v, nil
}

func (s *Service) GetWebCallback(ctx context.Context, val url.Values, provider string) (*model.User, error) {
	op := er.GetOperator()

	code := val.Get("code")
	_ = val.Get("scope")
	_ = val.Get("hd")
	_ = val.Get("prompt")
	_ = val.Get("authuser")

	// 1. get accessToken // refreshToken // expiry // auth_type // email
	ut, err := s.authRepository.GetUserInfoFromProvider(ctx, provider, code)
	if err != nil {
		return nil, er.WrapOp(err, op)
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

	return &model.User{
		UserToken: *ut,
	}, nil
}
