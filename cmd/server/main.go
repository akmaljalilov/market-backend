package main

import (
	"market/internal/bootstrap"
	config2 "market/internal/config"
)

func main() {
	config, err := config2.InitConfig()
	if err != nil {
		panic(err)
	}
	bootstrap.Start(config)
}
