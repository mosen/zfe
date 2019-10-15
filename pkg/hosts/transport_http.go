package hosts

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/examples/shipping/cargo"
	"github.com/gorilla/mux"
)
import "net/http"
import kitlog "github.com/go-kit/kit/log"
import kittransport "github.com/go-kit/kit/transport"
import kithttp "github.com/go-kit/kit/transport/http"

func MakeHandler(hostSvc HostService, templateSvc TemplateService, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	hostsIndexHandler := kithttp.NewServer(
		makeHostsIndexEndpoint(hostSvc),
		decodeIndexRequest,
		encodeResponse,
		opts...,
	)

	templatesIndexHandler := kithttp.NewServer(
		makeTemplatesIndexEndpoint(templateSvc),
		decodeIndexRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/hosts/v1/", hostsIndexHandler).Methods("GET")
	r.Handle("/templates/v1/", templatesIndexHandler).Methods("GET")

	return r
}

func decodeIndexRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case cargo.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	//case ErrInvalidArgument:
	//	w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
