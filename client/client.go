package client

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/endpoints"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/pb"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/service"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/transport"
	"google.golang.org/grpc"
)

// New ...
func New(conn *grpc.ClientConn) service.Service {
	var wordEndpoint = grpctransport.NewClient(
		conn, "pb.Lorem", "Word",
		transport.EncodeWordRequest,
		transport.DecodeWordResponse,
		pb.WordResponse{},
	).Endpoint()

	var sentenceEndpoint = grpctransport.NewClient(
		conn, "pb.Lorem", "Sentence",
		transport.EncodeSentenceRequest,
		transport.DecodeSentenceResponse,
		pb.SentenceResponse{},
	).Endpoint()

	var paragraphEndpoint = grpctransport.NewClient(
		conn, "pb.Lorem", "Paragraph",
		transport.EncodeParagraphRequest,
		transport.DecodeParagraphResponse,
		pb.ParagraphResponse{},
	).Endpoint()

	return endpoints.Endpoints{
		WordEndpoint:      wordEndpoint,
		SentenceEndpoint:  sentenceEndpoint,
		ParagraphEndpoint: paragraphEndpoint,
	}
}
