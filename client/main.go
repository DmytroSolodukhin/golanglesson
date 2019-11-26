package main

import (
	env "github.com/kazak/golanglesson/client/env"
	service "github.com/kazak/golanglesson/client/services"
)

func main() {
	grpcService := service.Connect(env.Settings.Api.Host, env.Settings.Api.Port)
	grpcService.UploadFile(env.Settings.FileName)
}