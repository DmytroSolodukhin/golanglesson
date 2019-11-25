package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	api "github.com/kazak/golanglesson/api"
	"log"
	"net"
	context "context"
)

type repository interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (api.streamServiceUploadClient, error)
}

type service struct {
	repo repository
}

func main() {
	address := fmt.Sprintf("%v:%v", "localhost", 27017)
	mongoConn, _ := mgo.Dial(address)

	_ = mongoConn.DB("test")

	lis, err := net.Listen("tcp", PortTOConnect)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	api.RegisterStreamServiceServer(server, &service{rep})

	// Register to response gRPC.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
