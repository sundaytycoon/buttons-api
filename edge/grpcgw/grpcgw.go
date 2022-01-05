package grpcgw

import (
	"crypto/tls"
	"fmt"
	"io/fs"
	"mime"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/sundaytycoon/buttons-api/insecure"
	"github.com/sundaytycoon/buttons-api/pkg/er"
	"github.com/sundaytycoon/buttons-api/third_party"
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
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("requested", r.URL.String())
				fmt.Println(r.URL.Path, strings.HasSuffix(r.URL.Path, "/api"))
				if strings.HasPrefix(r.URL.Path, "/api") {
					fmt.Println("/api")
					mux.ServeHTTP(w, r)
					return
				}
				getOpenAPIHandler().ServeHTTP(w, r)
			}),
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{insecure.Cert},
			},
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

	log.Info().Str("endpoint", endpoint).Msgf("run http application")
	if err = gw.httpServer.Serve(httpListener); err != nil {
		if err != http.ErrServerClosed {
			return er.WrapOp(err, op)
		}
	}
	return nil
}

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
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
