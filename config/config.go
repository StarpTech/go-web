package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

var config *configuration

type configuration struct {
	Port             string   `env:"PORT" envDefault:":8080"`
	Dialect          string   `env:"dialect" envDefault:":postgres"`
	ConnectionString string   `env:"connectionString" envDefault:"host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword"`
	IsProduction     bool     `env:"PRODUCTION"`
	Hosts            []string `env:"HOSTS" envSeparator:":"`
}

func init() {

	cfg := configuration{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	config = &cfg
}

func GetConfig() *configuration {
	return config
}
