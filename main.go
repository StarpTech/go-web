package main

import (
	config "github.com/starptech/go-web/config"
	server "github.com/starptech/go-web/server"
)

func main() {
	config := config.GetConfig()

	// migration
	m := server.Migration{}
	m.Up()

	// start server
	echo := server.NewEngine()
	echo.Logger.Fatal(echo.Start(config.Port))
}
