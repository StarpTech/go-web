package main

import (
	config "github.com/starptech/go-web/config"
	db "github.com/starptech/go-web/db"
	server "github.com/starptech/go-web/server"
)

func main() {
	config := config.GetConfig()

	// migration
	database := db.GetDB()
	server.Up(database)

	// start server
	echo := server.NewEngine()
	echo.Logger.Fatal(echo.Start(config.Port))
}
