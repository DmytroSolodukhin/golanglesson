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

type gRPCSettings struct {
	Host string
	Port int
}

type grpcServiceInterface interface {
	ConnectGRPC(settings gRPCSettings)
	UploadFile(file string)
}

type GrpcService struct {
	client api.StreamServiceClient
}

func (service GrpcService) ConnectGRPC(settings gRPCSettings) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", settings.Host, settings.Port),
		grpc.WithInsecure(),
	)
	if  err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	service.client = api.NewStreamServiceClient(conn)
}

func (service GrpcService) UploadFile(file string) {

	stream, err := service.client.Upload(context.Background())

	data, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer data.Close()

	buffer := make([]byte, 1024)
	writing := false
	count := 0

	for writing {
		count, err = data.Read(buffer)
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