package main

import (
	"cafe/api"
	"cafe/config"
)

func main() {
	a := api.Api{
		Auth: &config.Config.Auth,
	}
	a.Run()
}
