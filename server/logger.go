package server

import (
	"github.com/labstack/gommon/log"
)

func NewLogger(colors bool) *log.Logger {
	logger := log.New("server")

	// enable colors to beauty the beast
	if colors {
		logger.EnableColor()
	}

	return logger
}
