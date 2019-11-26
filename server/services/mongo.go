package services

import (
	"fmt"
	"github.com/globalsign/mgo"
)

type Mongo struct {
	Connection *mgo.Session
	Database   *mgo.Database
}

func MongoConnect(host string, port int, db string) (*Mongo, error) {

	address := fmt.Sprintf("%v:%v", host, port)
	mongoConn, _ := mgo.Dial(address)

	return &Mongo{
		Connection: mongoConn,
		Database:   mongoConn.DB(db),
	}, nil
}
