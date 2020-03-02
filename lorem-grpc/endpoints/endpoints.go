package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/service"
)

// Endpoints ...
type Endpoints struct {
	GenerateLorem endpoint.Endpoint
}

// MakeGenerateLoremEndpoint ...
func MakeGenerateLoremEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GenerateLoremRequest)

		var (
			min, max int
		)

		min = int(req.Min)
		max = int(req.Max)

		msg, err := svc.Lorem(ctx, req.RequestType, min, max)
		if err != nil {
			return nil, err
		}

		return GenerateLoremResponse{Message: msg}, nil
	}
}

// Lorem ...
func (e Endpoints) Lorem(ctx context.Context, requestType string, min, max int) (string, error) {
	req := GenerateLoremRequest{
		RequestType: requestType,
		Min:         int32(min),
		Max:         int32(max),
	}

	res, err := e.GenerateLorem(ctx, req)
	if err != nil {
		return "", err
	}

	loremRes := res.(GenerateLoremResponse)
	if loremRes.Err != "" {
		return "", errors.New(loremRes.Err)
	}

	return loremRes.Message, nil
}
