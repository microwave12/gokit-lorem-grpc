package transport

import (
	"context"

	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/endpoints"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/pb"
)

// EncodeWordRequest ...
func EncodeWordRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.LoremRequest)
	return &pb.WordRequest{
		Min: req.Min,
		Max: req.Max,
	}, nil
}

// DecodeWordRequest ...
func DecodeWordRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.WordRequest)
	return endpoints.LoremRequest{
		Min: req.Min,
		Max: req.Max,
	}, nil
}

// EncodeWordResponse ...
func EncodeWordResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.LoremResponse)
	return &pb.WordResponse{
		Message: res.Message,
	}, nil
}

// DecodeWordResponse ...
func DecodeWordResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.WordResponse)
	return endpoints.LoremResponse{
		Message: res.Message,
	}, nil
}

// EncodeSentenceRequest ...
func EncodeSentenceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.LoremRequest)
	return &pb.SentenceRequest{
		Min: req.Min,
		Max: req.Max,
	}, nil
}

// DecodeSentenceRequest ...
func DecodeSentenceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SentenceRequest)
	return endpoints.LoremRequest{
		Min: req.Min,
		Max: req.Max,
	}, nil
}

// EncodeSentenceResponse ...
func EncodeSentenceResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.LoremResponse)
	return &pb.SentenceResponse{
		Message: res.Message,
	}, nil
}

// DecodeSentenceResponse ...
func DecodeSentenceResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.SentenceResponse)
	return endpoints.LoremResponse{
		Message: res.Message,
	}, nil
}

// EncodeParagraphRequest ...
func EncodeParagraphRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.LoremRequest)
	return &pb.ParagraphRequest{
		Min: req.Min,
		Max: req.Max,
	}, nil
}

// DecodeParagraphRequest ...
func DecodeParagraphRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ParagraphRequest)
	return endpoints.LoremRequest{
		Min: req.Min,
		Max: req.Max,
	}, nil
}

// EncodeParagraphResponse ...
func EncodeParagraphResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.LoremResponse)
	return &pb.ParagraphResponse{
		Message: res.Message,
	}, nil
}

// DecodeParagraphResponse ...
func DecodeParagraphResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.ParagraphResponse)
	return endpoints.LoremResponse{
		Message: res.Message,
	}, nil
}
