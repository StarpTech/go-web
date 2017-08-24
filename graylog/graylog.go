package graylog

import (
	"github.com/labstack/gommon/log"
	"github.com/starptech/go-web/config"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

func NewGraylogLogger(config *config.Configuration) *gelf.TCPWriter {
	tw, err := gelf.NewTCPWriter(config.GrayLogAddr)

	if err != nil {
		log.Fatalf("gelf.NewWriter: %s", err)
	}

	return tw
}
