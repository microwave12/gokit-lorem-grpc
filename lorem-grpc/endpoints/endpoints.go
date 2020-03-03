package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/service"
)

// Endpoints ...
type Endpoints struct {
	WordEndpoint      endpoint.Endpoint
	SentenceEndpoint  endpoint.Endpoint
	ParagraphEndpoint endpoint.Endpoint
}

// MakeWordEndpoint ...
func MakeWordEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)
		min := int(req.Min)
		max := int(req.Max)
		msg, err := svc.Word(ctx, min, max)
		if err != nil {
			return nil, err
		}

		return LoremResponse{Message: msg}, nil
	}
}

// MakeSentenceEndpoint ...
func MakeSentenceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)
		min := int(req.Min)
		max := int(req.Min)
		msg, err := svc.Sentence(ctx, min, max)
		if err != nil {
			return nil, err
		}

		return LoremResponse{Message: msg}, nil
	}
}

// MakeParagraphEndpoint ...
func MakeParagraphEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)
		min := int(req.Min)
		max := int(req.Max)
		msg, err := svc.Paragraph(ctx, min, max)
		if err != nil {
			return nil, err
		}

		return LoremResponse{Message: msg}, nil
	}
}

// Word ...
func (e Endpoints) Word(ctx context.Context, min, max int) (string, error) {
	req := LoremRequest{
		Min: int32(min),
		Max: int32(max),
	}

	res, err := e.WordEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	wordRes := res.(LoremResponse)
	return wordRes.Message, nil
}

// Sentence ...
func (e Endpoints) Sentence(ctx context.Context, min, max int) (string, error) {
	req := LoremRequest{
		Min: int32(min),
		Max: int32(max),
	}

	res, err := e.SentenceEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	sentenceRes := res.(LoremResponse)
	return sentenceRes.Message, nil
}

// Paragraph ...
func (e Endpoints) Paragraph(ctx context.Context, min, max int) (string, error) {
	req := LoremRequest{
		Min: int32(min),
		Max: int32(max),
	}

	res, err := e.ParagraphEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	paragraphRes := res.(LoremResponse)
	return paragraphRes.Message, nil
}
