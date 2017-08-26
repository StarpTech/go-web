package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Address          string `env:"ADDRESS" envDefault:":8080"`
	Dialect          string `env:"DIALECT,required" envDefault:"postgres"`
	AssetsPublicDir  string `env:"ASSETS_PUBLIC_DIR,required"`
	AssetsBuildDir   string `env:"ASSETS_BUILD_DIR"`
	TemplateDir      string `env:"TPL_DIR"`
	LayoutDir        string `env:"LAYOUT_DIR"`
	RedisAddr        string `env:"REDIS_ADDR" envDefault:":6379"`
	RedisPwd         string `env:"REDIS_PWD"`
	ConnectionString string `env:"CONNECTION_STRING,required"`
	IsProduction     bool   `env:"PRODUCTION"`
	GrayLogAddr      string `env:"GRAYLOG_ADDR"`
	RequestLogger    bool   `env:"REQUEST_LOGGER"`
	LocaleDir        string `env:"LOCALE_DIR" envDefault:"locales"`
	Lang             string `env:"LANG" envDefault:"en_US"`
	LangDomain       string `env:"LANG_DOMAIN" envDefault:"default"`
	JwtSecret        string `env:"JWT_SECRET,required"`
}

func NewConfig(files ...string) *Configuration {
	err := godotenv.Load(files...)

	if err != nil {
		log.Printf("No .env file could be found %q\n", files)
	}

	cfg := Configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
