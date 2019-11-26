package main

import (
	env "github.com/kazak/golanglesson/client/env"
	service "github.com/kazak/golanglesson/client/services"
)

func main() {
	service.ConnectGRPC(env.Settings.Api)
	service.UploadFile(env.Settings.FileName)
}