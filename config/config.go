package config

import (
	"fmt"
	"os"
	"path"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

var config *configuration

type configuration struct {
	Address          string `env:"ADDRESS" envDefault:":8080"`
	Dialect          string `env:"DIALECT" envDefault:"postgres"`
	PublicDir        string `env:"PUBLIC_DIR" envDefault:"public"`
	TemplateDir      string `env:"TPL_DIR" envDefault:"templates"`
	ConnectionString string `env:"CONNECTION_STRING" envDefault:""`
	IsProduction     bool   `env:"PRODUCTION"`
	GrayLogAddr      string `env:"GRAYLOG_ADDR" envDefault:""`
}

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Info("No .env file could be found")
	}

	cfg := configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	cfg.TemplateDir = path.Join(wd, cfg.TemplateDir)
	cfg.PublicDir = path.Join(wd, cfg.PublicDir)
	config = &cfg
}

func GetConfig() *configuration {
	return config
}
