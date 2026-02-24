package main

import (
	"market/app"
	config2 "market/pkg/config"
)

func main() {
	config, err := config2.InitConfig()
	if err != nil {
		panic(err)
	}
	app.Start(config)
}
