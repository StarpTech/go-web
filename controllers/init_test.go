package controllers

import (
	"os"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/logger"
	"github.com/starptech/go-web/models"
	"github.com/starptech/go-web/server"
)

var e struct {
	config *config.Configuration
	logger *log.Logger
	engine *server.Server
}

func TestMain(m *testing.M) {
	e.config = &config.Configuration{
		ConnectionString: "host=localhost user=gorm dbname=gorm sslmode=disable password=mypassword",
		TemplateDir:      "../templates",
		Dialect:          "postgres",
		RedisAddr:        ":6379",
	}

	e.logger = logger.NewLogger(e.config)
	e.engine = server.NewServer(e.config)
	e.engine.SetLogger(e.logger)

	setup()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func setup() {
	userCtrl := new(User)
	feedCtrl := new(Feed)
	healthCtrl := new(Healthcheck)
	importCtrl := new(Importer)

	g := e.engine.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON(e.engine))

	u := e.engine.Echo.Group("/users")
	u.GET("/:id", userCtrl.GetUser(e.engine))
	u.GET("/:id/details", userCtrl.GetUserDetails(e.engine))

	e.engine.Echo.POST("/import", importCtrl.ImportUser(e.engine))
	e.engine.Echo.GET("/feed", feedCtrl.GetFeed(e.engine))
	e.engine.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck(e.engine))
	e.engine.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	user := models.User{Name: "peter"}

	e.engine.GetDB().Register(user)
	e.engine.GetDB().AutoMigrateAll()
	e.engine.GetDB().Create(&user)
}

func tearDown() {
	e.engine.GetDB().AutoDropAll()
}
