package server

import (
	"io"
	"os"

	"github.com/labstack/gommon/log"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

func NewLogger(grayLogAddr string, isProduction bool) *log.Logger {
	l := log.New("server")

	if grayLogAddr != "" {
		gelfWriter, err := gelf.NewUDPWriter(grayLogAddr)
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		// Log to greylog and stderr
		l.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
	}

	// enable colors to beauty the beast
	if !isProduction {
		l.EnableColor()
		l.SetLevel(log.DEBUG)
	} else {
		l.SetLevel(log.ERROR)
	}

	return l
}
