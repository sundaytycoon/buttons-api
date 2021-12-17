package grpcgw

import (
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/sundaytycoon/buttons-api/pkg/er"
)

type GRPCServer interface {
	Start()
	Close() error
}

type Gateway struct {
	mux        *runtime.ServeMux
	httpServer *http.Server
}

func New() *Gateway {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)

	return &Gateway{
		mux: mux,
		httpServer: &http.Server{
			Handler: mux,
		},
	}
}

func (gw *Gateway) Start(endpoint string) error {
	op := er.GetOperator()

	gw.httpServer.Addr = endpoint
	httpListener, err := net.Listen("tcp", gw.httpServer.Addr)
	if err != nil {
		return er.WrapOp(err, op)
	}

	log.Info().Msgf("run http application")
	if err = gw.httpServer.Serve(httpListener); err != nil {
		if err != http.ErrServerClosed {
			return er.WrapOp(err, op)
		}
	}
	return nil
}

type GRPCHandler interface {
	Name() string
	Connect(grpcEndpoint string, mux *runtime.ServeMux) error
}

func (gw *Gateway) ConnectWithHandlers(grpcEndpoint string, handlers ...GRPCHandler) error {
	op := er.GetOperator()
	for _, h := range handlers {
		log.Info().Msgf("connect %s application", h.Name())
		if e := h.Connect(grpcEndpoint, gw.mux); e != nil {
			return er.WrapOp(e, h.Name()+"|"+op)
		}
	}
	return nil
}

func (gw *Gateway) Close() error {
	op := er.GetOperator()

	if err := gw.httpServer.Close(); err != nil {
		return er.WrapOp(err, op)
	}

	return nil
}
