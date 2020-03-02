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
	var loremEndpoint = grpctransport.NewClient(
		conn, "pb.Lorem", "Lorem",
		transport.EncodeLoremRequest,
		transport.DecodeLoremResponse,
		pb.LoremResponse{},
	).Endpoint()

	return endpoints.Endpoints{
		GenerateLorem: loremEndpoint,
	}
}
