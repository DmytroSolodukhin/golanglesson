package services

import (
	"google.golang.org/grpc"
	context "context"
	api "github.com/kazak/golanglesson/api"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type repository interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (api.streamServiceUploadClient, error)
}

type service struct {
	repo repository
}

func StartGRPCServer(port string, protocol string) {
	lis, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	api.RegisterPortServiceServer(server, &service{rep})

	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *service) Upload(ctx context.Context, req *api.Request) (*api.streamServiceUploadClient, error) {

}

