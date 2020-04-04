package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/internal/controller"
	"github.com/starptech/go-web/internal/core"
	"github.com/starptech/go-web/internal/models"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// create server
	server := core.NewServer(config)
	// serve files for dev
	server.ServeStaticFiles()

	userCtrl := &controller.User{}
	userListCtrl := &controller.UserList{}
	healthCtrl := &controller.Healthcheck{}

	// api endpoints
	g := server.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON)

	// pages
	u := server.Echo.Group("/users")
	u.GET("", userListCtrl.GetUsers)
	u.GET("/:id", userCtrl.GetUser)

	// metric / health endpoint according to RFC 5785
	server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// migration for dev
	user := models.User{Name: "Peter"}
	mr := server.GetModelRegistry()
	err = mr.Register(user)

	if err != nil {
		server.Echo.Logger.Fatal(err)
	}

	mr.AutoMigrateAll()
	mr.Create(&user)
	// Start server
	go func() {
		if err := server.Start(config.Address); err != nil {
			server.Echo.Logger.Info("shutting down the server")
		}
	}()

	server.GracefulShutdown()
}
