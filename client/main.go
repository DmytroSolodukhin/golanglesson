package main

import (
	env "github.com/kazak/client/env"
	service "github.com/kazak/client/services"
)

func main() {
	service.ConnectGRPC(env.Settings.Api)
	service.UploadFile(env.Settings.FileName)
}