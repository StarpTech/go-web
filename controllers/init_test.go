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
	userCtrl := &User{e.server}
	feedCtrl := &Feed{e.server}
	healthCtrl := &Healthcheck{e.server}
	importCtrl := &Importer{e.server}

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

	// bootstrap db
	err := e.server.GetDB().Register(user)
	if err != nil {
		log.Fatal(err)
	}

	e.server.GetDB().AutoMigrateAll()
	e.server.GetDB().Create(&user)
}

func tearDown() {
	e.server.GetDB().AutoDropAll()
}
