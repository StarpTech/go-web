package main

import (
	config "github.com/starptech/go-web/config"
	server "github.com/starptech/go-web/server"
)

func main() {
	config := config.GetConfig()
	echo := server.NewEngine()
	echo.Logger.Fatal(echo.Start(config.Port))
}
