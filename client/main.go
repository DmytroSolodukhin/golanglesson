package main

import (
	"io"
	"log"
	"google.golang.org/grpc"
	"github.com/pkg/errors"
	api "github.com/kazak/golanglesson/api"
	"os"
	"golang.org/x/net/context"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if  err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcClient := api.NewStreamServiceClient(conn)
	stream, err := grpcClient.Upload(context.Background())

	file, err := os.Open("files/googlechrome.dmg")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	buffer := make([]byte, 1024)
	writing := false
	count := 0

	for writing {
		count, err = file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil
				continue
			}

			err = errors.Wrapf(err, "errored while copying from file to buf")
			return
		}

		err = stream.Send(&api.Chunk{
			Content: buffer[:count],
		})
		if err != nil {
			err = errors.Wrapf(err, "failed to send chunk")
			return
		}
	}

}