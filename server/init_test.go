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
	}

	e.logger = NewLogger(e.config.GrayLogAddr, e.config.IsProduction)
	e.engine = NewEngine(e.config)
	e.engine.SetLogger(e.logger)

	setup()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func setup() {
	m := Migration{Db: e.engine.Db}
	m.Up()
}

func tearDown() {
	u := &models.User{}
	e.engine.Db.DropTableIfExists(u)
}
