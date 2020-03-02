package transport

import (
	"context"

	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/endpoints"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/pb"
)

// EncodeLoremRequest ...
func EncodeLoremRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.GenerateLoremRequest)
	return &pb.LoremRequest{
		RequestType: req.RequestType,
		Min:         req.Min,
		Max:         req.Max,
	}, nil
}

// DecodeLoremRequest ...
func DecodeLoremRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoremRequest)
	return endpoints.GenerateLoremRequest{
		RequestType: req.RequestType,
		Min:         req.Min,
		Max:         req.Max,
	}, nil
}

// EncodeLoremResponse ...
func EncodeLoremResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GenerateLoremResponse)
	return &pb.LoremResponse{
		Message: res.Message,
		Err:     res.Err,
	}, nil
}

// DecodeLoremResponse ...
func DecodeLoremResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.LoremResponse)
	return endpoints.GenerateLoremResponse{
		Message: res.Message,
		Err:     res.Err,
	}, nil
}
