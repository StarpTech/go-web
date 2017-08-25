package main

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/controllers"
	"github.com/starptech/go-web/core"
	"github.com/starptech/go-web/models"
)

func main() {
	config := config.NewConfig()

	// create server
	server := core.NewServer(config)
	// serve files for dev
	server.ServeStaticFiles()

	cache := server.GetCache()
	db := server.GetDB()

	userCtrl := &controllers.User{
		Cache:  &core.CacheStore{Cache: cache},
		Config: config,
		Store:  &core.UserStore{DB: db},
	}

	feedCtrl := &controllers.Feed{
		Cache:  &core.CacheStore{Cache: cache},
		Config: config,
		Store:  &core.UserStore{DB: db},
	}
	healthCtrl := &controllers.Healthcheck{
		Cache:  &core.CacheStore{Cache: cache},
		Config: config,
		Store:  &core.UserStore{DB: db},
	}
	importCtrl := &controllers.Importer{
		Cache:  &core.CacheStore{Cache: cache},
		Config: config,
		Store:  &core.UserStore{DB: db},
	}

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
