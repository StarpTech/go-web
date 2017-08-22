package logger

import (
	"io"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/starptech/go-web/config"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

var logger *log.Logger

func init() {
	c := config.GetConfig()

	l := log.New("server")

	if c.GrayLogAddr != "" {
		gelfWriter, err := gelf.NewUDPWriter(c.GrayLogAddr)
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		// Log to greylog and stderr
		l.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
	}

	// enable colors to beauty the beast
	if !c.IsProduction {
		l.EnableColor()
		l.SetLevel(log.DEBUG)
	} else {
		l.SetLevel(log.ERROR)
	}

	logger = l

}

func GetLogger() *log.Logger {
	return logger
}
