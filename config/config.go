package config

import (
	"fmt"
	"os"
	"path"

	"github.com/caarlos0/env"
)

var config *configuration

type configuration struct {
	Port             string   `env:"PORT" envDefault:":8080"`
	Dialect          string   `env:"dialect" envDefault:"postgres"`
	PublicDir        string   `env:"publicDir" envDefault:"public"`
	TemplateDir      string   `env:"templateDir" envDefault:"templates"`
	ConnectionString string   `env:"connectionString" envDefault:"host=localhost user=gorm dbname=gorm sslmode=disable password=mypassword"`
	IsProduction     bool     `env:"PRODUCTION"`
	Hosts            []string `env:"HOSTS" envSeparator:":"`
}

func init() {
	cfg := configuration{}
	err := env.Parse(&cfg)
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
