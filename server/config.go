package server

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Configuration struct {
	Address          string `env:"ADDRESS" envDefault:":8080"`
	Dialect          string `env:"DIALECT,required" envDefault:"postgres"`
	AssetsPublicDir  string `env:"ASSETS_PUBLIC_DIR,required"`
	AssetsBuildDir   string `env:"ASSETS_BUILD_DIR"`
	TemplateDir      string `env:"TPL_DIR" envDefault:"templates"`
	ConnectionString string `env:"CONNECTION_STRING,required"`
	IsProduction     bool   `env:"PRODUCTION"`
	GrayLogAddr      string `env:"GRAYLOG_ADDR"`
}

func NewConfig(files ...string) *Configuration {
	err := godotenv.Load(files...)

	if err != nil {
		log.Infof("No .env file could be found %q", files)
	}

	cfg := Configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
