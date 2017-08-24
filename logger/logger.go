package logger

import (
	"github.com/labstack/gommon/log"
	"github.com/starptech/go-web/config"
)

func NewLogger(config *config.Configuration) *log.Logger {
	logger := log.New("server")

	// enable colors to beauty the beast
	if config.CliColors {
		logger.EnableColor()
	}

	return logger
}
