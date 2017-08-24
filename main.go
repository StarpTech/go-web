package main

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/controllers"
	"github.com/starptech/go-web/models"
	"github.com/starptech/go-web/server"
)

func main() {
	config := config.NewConfig()

	// create server
	server := server.NewServer(config)
	// serve files for dev
	server.ServeStaticFiles()

	userCtrl := &controllers.User{server}
	feedCtrl := &controllers.Feed{server}
	healthCtrl := &controllers.Healthcheck{server}
	importCtrl := &controllers.Importer{server}

	// api rest endpoints
	g := server.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON)

	// pages
	u := server.Echo.Group("/users")
	u.GET("/:id", userCtrl.GetUser)
	u.GET("/:id/details", userCtrl.GetUserDetails)

	// special endpoints
	server.Echo.POST("/import", importCtrl.ImportUser)
	server.Echo.GET("/feed", feedCtrl.GetFeed)

	// metric / health endpoint according to RFC 5785
	server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// migration for dev
	user := models.User{Name: "peter"}
	err := server.GetDB().Register(user)
	if err != nil {
		server.Echo.Logger.Fatal(err)
	}
	server.GetDB().AutoMigrateAll()
	server.GetDB().Create(&user)

	// listen
	go func() {
		server.Echo.Logger.Fatal(server.Start(config.Address))
	}()

	server.GracefulShutdown()
}
