package grpcgw

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

const (
	HeaderLocation = "Location"
	HeaderStatus   = "Status"
)

func HeaderDispatcher(w http.ResponseWriter) {
	headers := w.Header()
	if v := headers.Get(runtime.MetadataHeaderPrefix + HeaderLocation); v != "" {
		w.Header().Set(HeaderLocation, v)
	}
	if v := headers.Get(runtime.MetadataHeaderPrefix + HeaderStatus); v != "" {
		status, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Error().Err(err).Msgf("parse error http status code at Header Dispatcher")
			status = 503
		}
		w.WriteHeader(int(status))
	}

	for k, _ := range headers {
		if strings.HasPrefix(k, runtime.MetadataHeaderPrefix) {
			originKey := strings.Replace(k, runtime.MetadataHeaderPrefix, "", 1)
			w.Header().Del(originKey)
		}
	}

}
