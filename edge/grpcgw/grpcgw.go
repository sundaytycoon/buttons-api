package grpcgw

import (
	"context"
	"io/fs"
	"mime"
	"net"
	"net/http"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/sundaytycoon/buttons-api/doc"
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
		runtime.WithForwardResponseOption(
			func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
				HeaderDispatcher(w)
				return nil
			},
		),
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
				Marshaler: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						UseProtoNames:   true,
						EmitUnpopulated: true,
					},
					UnmarshalOptions: protojson.UnmarshalOptions{
						DiscardUnknown: true,
					},
				},
			},
		),
	)

	return &Gateway{
		mux: mux,
		httpServer: &http.Server{
			Handler: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {

					var header [][]string
					for k, v := range r.Header {
						header = append(header, []string{k, strings.Join(v, "/")})
					}
					var reqBody map[string]interface{}
					log.Trace().
						Str("method", r.Method).
						Str("proto", r.Proto).
						Interface("url", r.RequestURI).
						Interface("header", header).
						Interface("body", reqBody).
						Send()
					if strings.HasPrefix(r.URL.Path, "/api") {
						mux.ServeHTTP(w, r)
						return
					}
					if strings.HasPrefix(r.URL.Path, "/test") {
						r.URL.Path = strings.ReplaceAll(r.URL.EscapedPath(), "/test", "")
						getTestHTMLHandler().ServeHTTP(w, r)
						return
					}
					getOpenAPIHandler().ServeHTTP(w, r)
				},
			),
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
	subFS, err := fs.Sub(doc.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

// getTestHTMLHandler serves an test html page
func getTestHTMLHandler() http.Handler {
	// Use subdirectory in embedded files
	mime.AddExtensionType(".svg", "image/svg+xml")
	subFS, err := fs.Sub(doc.Public, "public")
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
