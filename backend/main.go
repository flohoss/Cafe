package main

import (
	"cafe/api"
	"cafe/config"
)

func main() {
	a := api.Api{
		Auth: &config.C.Auth,
	}
	config.C.Database.Initialize(config.StorageDir)
	a.Run()
}
