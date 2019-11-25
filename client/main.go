package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"google.golang.org/grpc"
	api "github.com/kazak/golanglesson/api"
	"os"
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

	grpcClient := api.NewStreamServiceClient(conn)

	file, err := os.Open("files/googlechrome.dmg")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		_, err = reader.Read(buffer);
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		chunk := api.Chunk{Content: buffer}
	}

}