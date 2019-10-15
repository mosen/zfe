package hosts

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeHostsIndexEndpoint(svc HostService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		// req := request.(indexRequest)
		hosts, err := svc.Find()
		if err != nil {
			return hostsIndexResponse{Errors: []string{err.Error()}}, nil
		}

		return hostsIndexResponse{Data: hosts}, nil
	}
}

func makeTemplatesIndexEndpoint(svc TemplateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		// req := request.(indexRequest)
		templates, err := svc.Find()
		if err != nil {
			return templatesIndexResponse{Errors: []string{err.Error()}}, nil
		}

		return templatesIndexResponse{Data: templates}, nil
	}
}
