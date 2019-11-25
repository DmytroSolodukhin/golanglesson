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
}
