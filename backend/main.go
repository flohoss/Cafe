package main

import (
	"cafe/api"
	"cafe/config"
	"cafe/service"
)

func main() {
	a := api.Api{
		Auth: &config.C.Auth,
	}
	config.C.Database.Initialize(config.StorageDir)
	service.Initialize()
	a.Hub.Initialize()
	a.Run()
}
