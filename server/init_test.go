package server

import (
	"os"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/starptech/go-web/models"
)

var e struct {
	config *Configuration
	logger *log.Logger
	engine *Engine
}

func TestMain(m *testing.M) {
	e.config = &Configuration{
		ConnectionString: "host=localhost user=gorm dbname=gorm sslmode=disable password=mypassword",
		TemplateDir:      "../templates",
		Dialect:          "postgres",
		RedisAddr:        ":6379",
	}

	e.logger = NewLogger(false)
	e.engine = NewEngine(e.config)
	e.engine.SetLogger(e.logger)

	setup()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func setup() {
	m := Migration{Db: e.engine.GetDB()}
	m.Up()
}

func tearDown() {
	u := &models.User{}
	e.engine.GetDB().DropTableIfExists(u)
}
