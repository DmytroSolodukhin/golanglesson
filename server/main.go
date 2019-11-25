package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/globalsign/mgo"
	"io"
	"lesson/server/api"
	"log"
	"os"
)

func main() {
	address := fmt.Sprintf("%v:%v", "localhost", 27017)
	mongoConn, _ := mgo.Dial(address)

	_ = mongoConn.DB("test")


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
