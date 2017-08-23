package main

import (
	"github.com/starptech/go-web/server"
)

func main() {
	e := server.NewEngine()

	m := server.Migration{Db: e.Db}
	m.Up()

	go func() {
		e.Echo.Logger.Fatal(e.Echo.Start(e.Config.Address))
	}()

	e.GracefulShutdown()
}
