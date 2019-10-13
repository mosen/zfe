package hosts

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeIndexEndpoint(svc HostService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		// req := request.(indexRequest)
		hosts, err := svc.Find()
		if err != nil {
			return indexResponse{Errors: []string{err.Error()}}, nil
		}

		return indexResponse{Data: hosts}, nil
	}
}
