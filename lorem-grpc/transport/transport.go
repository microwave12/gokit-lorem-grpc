package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/endpoints"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/pb"
)

type grpcServer struct {
	wordHandler      grpctransport.Handler
	sentenceHandler  grpctransport.Handler
	paragraphHandler grpctransport.Handler
}

func (gs *grpcServer) Word(ctx context.Context, r *pb.WordRequest) (*pb.WordResponse, error) {
	_, res, err := gs.wordHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*pb.WordResponse), nil
}

func (gs *grpcServer) Sentence(ctx context.Context, r *pb.SentenceRequest) (*pb.SentenceResponse, error) {
	_, res, err := gs.sentenceHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*pb.SentenceResponse), nil
}

func (gs *grpcServer) Paragraph(ctx context.Context, r *pb.ParagraphRequest) (*pb.ParagraphResponse, error) {
	_, res, err := gs.paragraphHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*pb.ParagraphResponse), nil
}

// NewGRPCService ...
func NewGRPCService(_ context.Context, endpoints endpoints.Endpoints) pb.LoremServer {
	return &grpcServer{
		wordHandler: grpctransport.NewServer(
			endpoints.WordEndpoint,
			DecodeWordRequest,
			EncodeWordResponse,
		),
		sentenceHandler: grpctransport.NewServer(
			endpoints.SentenceEndpoint,
			DecodeSentenceRequest,
			EncodeSentenceResponse,
		),
		paragraphHandler: grpctransport.NewServer(
			endpoints.ParagraphEndpoint,
			DecodeParagraphRequest,
			EncodeParagraphResponse,
		),
	}
}
