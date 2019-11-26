package services

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"github.com/kazak/golanglesson/client/env"
	api "github.com/kazak/golanglesson/api"
)

func ConnectGRPC() api.StreamServiceClient {
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", env.Settings.Api.Host, env.Settings.Api.Port),
		grpc.WithInsecure(),
	)
	if  err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	return api.NewStreamServiceClient(conn)
}