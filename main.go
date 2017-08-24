package main

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/controllers"
	"github.com/starptech/go-web/logger"
	"github.com/starptech/go-web/server"
)

func main() {
	config := config.NewConfig()
	logger := logger.NewLogger(config)

	// create server
	server := server.NewServer(config)
	server.SetLogger(logger)
	server.ServeStaticFiles()

	userCtrl := new(controllers.User)
	feedCtrl := new(controllers.Feed)
	healthCtrl := new(controllers.Healthcheck)
	importCtrl := new(controllers.Importer)

	// api rest endpoints
	g := server.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON(server))

	// pages
	u := server.Echo.Group("/users")
	u.GET("/:id", userCtrl.GetUser(server))
	u.GET("/:id/details", userCtrl.GetUserDetails(server))

	// special endpoints
	server.Echo.POST("/import", importCtrl.ImportUser(server))
	server.Echo.GET("/feed", feedCtrl.GetFeed(server))

	// metric / health endpoint according to RFC 5785
	server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck(server))
	server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// listen
	go func() {
		logger.Fatal(server.Start(config.Address))
	}()

	server.GracefulShutdown()
}
