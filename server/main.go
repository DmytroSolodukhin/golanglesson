package main

import (
	"github.com/kazak/golanglesson/service/services"
	"google.golang.org/grpc"
	env "github.com/kazak/golanglesson/server/env"
)

func main() {
	_, _ = services.MongoConnect(env.Settings.MngoDB.Host, env.Settings.MngoDB.Port, env.Settings.MngoDB.DB)
	services.StartGRPCServer(env.Settings.Api.Port, repository)
}
