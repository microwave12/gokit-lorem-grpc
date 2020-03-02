package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/endpoints"
	"github.com/microwave12/gokit-lorem-grpc/lorem-grpc/pb"
)

type grpcServer struct {
	lorem grpctransport.Handler
}

func (gs *grpcServer) Lorem(ctx context.Context, r *pb.LoremRequest) (*pb.LoremResponse, error) {
	_, res, err := gs.lorem.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*pb.LoremResponse), nil
}

// NewGRPCService ...
func NewGRPCService(_ context.Context, endpoints endpoints.Endpoints) pb.LoremServer {
	return &grpcServer{
		lorem: grpctransport.NewServer(
			endpoints.GenerateLorem,
			DecodeLoremRequest,
			EncodeLoremResponse,
		),
	}
}
