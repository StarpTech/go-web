package main

import (
	"github.com/starptech/go-web/server"
)

func main() {
	config := server.NewConfig()
	logger := server.NewLogger(config.CliColors)

	engine := server.NewEngine(config)
	engine.SetLogger(logger)
	engine.ServeStaticFiles()

	m := server.Migration{Db: engine.GetDB()}
	m.Up()

	go func() {
		logger.Fatal(engine.Start(config.Address))
	}()

	engine.GracefulShutdown()
}
