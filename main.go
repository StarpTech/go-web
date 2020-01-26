package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/controller"
	"github.com/starptech/go-web/core"
	"github.com/starptech/go-web/models"
)

func main() {
	config := config.NewConfig()
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
	err := mr.Register(user)

	if err != nil {
		server.Echo.Logger.Fatal(err)
	}

	mr.AutoMigrateAll()
	mr.Create(&user)

	// listen
	go func() {
		server.Echo.Logger.Fatal(server.Start(config.Address))
	}()

	server.GracefulShutdown()
}
