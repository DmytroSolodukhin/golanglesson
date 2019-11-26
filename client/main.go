package main

import (
	"io"
	"log"
	"github.com/pkg/errors"
	"github.com/kazak/client/env"
	servioce "github.com/kazak/client/services"
	"os"
	context "context"
)

func main() {
	grpcClient := Ssrvice.ConnectGRPC()
	stream, err := grpcClient.Upload(context.Background())

	file, err := os.Open(env.Settings.FileName)

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