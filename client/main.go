package main

import (
	"log"
	"fmt"
	"google.golang.org/grpc"
	api ""
)

const (
	selfHost = ":9090"
	sendAddress = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(sendAddress, grpc.WithInsecure())
	if  err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcClient := api.NewPortServiceClient(conn)

	restapi.Start(selfHost, grpcClient)
}