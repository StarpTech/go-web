package main

import (
	"context"
	"os"
	"os/signal"
	"time"

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

	go func() {
		echo.Logger.Fatal(echo.Start(config.Address))
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := echo.Shutdown(ctx); err != nil {
		echo.Logger.Fatal(err)
	}
}
