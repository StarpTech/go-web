package server

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

func NewGraylogLogger(grayLogAddr string) *gelf.TCPWriter {
	tw, err := gelf.NewTCPWriter(grayLogAddr)

	if err != nil {
		log.Fatalf("gelf.NewWriter: %s", err)
	}

	return tw
}
