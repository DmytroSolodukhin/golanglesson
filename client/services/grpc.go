package services

import (
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"io"
	"log"
	api "github.com/kazak/golanglesson/api"
	"os"
	context "context"
)

type GrpcServiceInterface interface {
	UploadFile(file string)
}

type grpcService struct {
	client api.StreamServiceClient
}

func Connect(host string, port int) GrpcServiceInterface {
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", host, port),
		grpc.WithInsecure(),
	)
	if  err != nil {
		log.Fatal(err)
	}
	return &grpcService{
		client: api.NewStreamServiceClient(conn),
	}
}

func (service *grpcService) UploadFile(file string) {

	stream, err := service.client.Upload(context.Background())
	data, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer data.Close()

	buffer := make([]byte, 1024)
	writing := true
	count := 0
	fmt.Println("start to upload file")

	for writing {
		count, err = data.Read(buffer)

		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil
				continue
				fmt.Println("file loaded")
			}

			errorMessage := "errored while copying from file to buf"
			err = errors.Wrapf(err, errorMessage)
			fmt.Println(errorMessage)

			return
		}
		fmt.Println("chunk sending...")

		err = stream.Send(&api.Chunk{
			Content: buffer[:count],
		})
		if err != nil {
			err = errors.Wrapf(err, "failed to send chunk")
			return
		}
	}
}