package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc/metadata"
)

const (
	HandlerTypeHTTP = "http"
	HandlerTypeGRPC = "grpc"

	MDCountryCode = "country_code"
	MDEnv         = "env"
	MDRequestID   = "request_id"
	MDRequestIP   = "request_ip"
	MDHostIP      = "host_ip"
	MDHandler     = "handler"
	MDUserAgent   = "user_agent"
	MDReferer     = "referer"
)

type meta struct {
	Env         string
	CountryCode string
	RequestIP   string
	HostIP      string
	RequestID   string
	Handler     string
	UserAgent   string
	Referer     string
	Actor       string
}

// WithRequestMetadataFromHTTP http request에 담긴 값을 context 객체에 넣는 역할을 한다.
func WithRequestMetadataFromHTTP(r *http.Request, env string) context.Context {
	ctx := r.Context()

	reqIdV := ctx.Value(chi_middleware.RequestIDKey)
	var reqId string
	if v, ok := reqIdV.(string); ok {
		reqId = v
	}
	hostIp := os.Getenv("HOST_IP")
	if hostIp == "" {
		hostIp = "0.0.0.0"
	}

	md := metadata.New(map[string]string{
		MDRequestIP: r.RemoteAddr,
		MDRequestID: reqId,
		MDHostIP:    hostIp,
		MDHandler:   fmt.Sprintf("%s.%s", HandlerTypeHTTP, chi.RouteContext(ctx).RoutePattern()),
		MDUserAgent: r.UserAgent(),
		MDReferer:   r.Referer(),
		MDEnv:       env,
	})

	return metadata.NewIncomingContext(ctx, md)
}

// RequestMetadataFromHTTP get metadata from context
func RequestMetadataFromHTTP(ctx context.Context) meta {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return meta{
			RequestIP:   strings.Join(md.Get(MDRequestIP), ","),
			RequestID:   strings.Join(md.Get(MDRequestID), ","),
			HostIP:      strings.Join(md.Get(MDHostIP), ","),
			Handler:     strings.Join(md.Get(MDHandler), ","),
			UserAgent:   strings.Join(md.Get(MDUserAgent), ","),
			Referer:     strings.Join(md.Get(MDReferer), ","),
			CountryCode: strings.Join(md.Get(MDCountryCode), ","),
			Env:         strings.Join(md.Get(MDEnv), ","),
		}
	}
	return meta{}
}
