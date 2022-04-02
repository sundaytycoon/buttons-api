package auth

//import (
//	"context"
//
//	"github.com/rs/zerolog/log"
//
//	buttonsapi "github.com/sundaytycoon/buttons-api"
//	"github.com/sundaytycoon/buttons-api/internal/utils/er"
//)
//
//type GetWebCallbackIn struct {
//	Provider string
//
//	Code     string
//	State    string
//	Scope    string
//	AuthUser string
//	Prompt   string
//	Hd       string
//}
//
//type GetWebCallbackOut struct {
//	Provider       string
//	ToHost         string
//	TemporaryToken string
//}
//
//func (h *Handler) GetWebCallback(ctx context.Context, in *GetWebCallbackIn) (*GetWebCallbackOut, error) {
//	op := er.GetOperator()
//
//	if in.Code == "" {
//		err := er.New("code is required", buttonsapi.ErrBadRequest)
//		return nil, er.WrapOp(err, op)
//	}
//
//	if in.State == "" {
//		err := er.New("state is required", buttonsapi.ErrBadRequest)
//		return nil, er.WrapOp(err, op)
//	}
//
//	providers := []string{buttonsapi.Google}
//	var isIn bool
//	for _, p := range providers {
//		if p == in.Provider {
//			isIn = true
//		}
//	}
//	if isIn == false {
//		err := er.New("provider is not valid", buttonsapi.ErrBadRequest)
//		return nil, er.WrapOp(err, op)
//	}
//
//	toHost, tempToken, err := h.authService.GetWebCallback(ctx, in.Provider, in.Code, in.State)
//	if err != nil {
//		log.Error().Err(err).Send()
//		return nil, er.WrapOp(err, op)
//	}
//
//	res := &GetWebCallbackOut{
//		Provider:       in.Provider,
//		ToHost:         toHost,
//		TemporaryToken: tempToken,
//	}
//
//	return res, nil
//}
