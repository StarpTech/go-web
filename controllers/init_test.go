package controllers

import (
	"os"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/core"
	"github.com/starptech/go-web/models"
)

var e struct {
	config *config.Configuration
	logger *log.Logger
	server *core.Server
}

func TestMain(m *testing.M) {
	e.config = &config.Configuration{
		ConnectionString: "host=localhost user=gorm dbname=gorm sslmode=disable password=mypassword",
		TemplateDir:      "../templates/*.html",
		LayoutDir:        "../templates/layouts/*.html",
		Dialect:          "postgres",
		RedisAddr:        ":6379",
	}

	e.server = core.NewServer(e.config)

	setup()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func setup() {
	db := e.server.GetDB()
	cache := e.server.GetCache()

	userCtrl := &User{
		Cache:  &core.CacheStore{Cache: cache},
		Config: e.config,
		Store:  &core.UserStore{DB: db},
	}

	feedCtrl := &Feed{
		Cache:  &core.CacheStore{Cache: cache},
		Config: e.config,
		Store:  &core.UserStore{DB: db},
	}
	healthCtrl := &Healthcheck{
		Cache:  &core.CacheStore{Cache: cache},
		Config: e.config,
		Store:  &core.UserStore{DB: db},
	}
	importCtrl := &Importer{
		Cache:  &core.CacheStore{Cache: cache},
		Config: e.config,
		Store:  &core.UserStore{DB: db},
	}

	g := e.server.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON)

	u := e.server.Echo.Group("/users")
	u.GET("/:id", userCtrl.GetUser)
	u.GET("/:id/details", userCtrl.GetUserDetails)

	e.server.Echo.POST("/import", importCtrl.ImportUser)
	e.server.Echo.GET("/feed", feedCtrl.GetFeed)
	e.server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	e.server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// test data
	user := models.User{Name: "peter"}
	mr := e.server.GetModelRegistry()
	err := mr.Register(user)

	if err != nil {
		e.server.Echo.Logger.Fatal(err)
	}
	mr.AutoMigrateAll()
	mr.Create(&user)
}

func tearDown() {
	e.server.GetModelRegistry().AutoDropAll()
}
