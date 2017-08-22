package logger

import (
	"io"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/starptech/go-web/config"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

var Log *log.Logger

func init() {
	c := config.GetConfig()

	logger := log.New("server")

	if c.GrayLogAddr != "" {
		gelfWriter, err := gelf.NewUDPWriter(c.GrayLogAddr)
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		// Log to greylog and stderr
		logger.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
	}

	// enable colors to beauty the beast
	if !c.IsProduction {
		logger.EnableColor()
		logger.SetLevel(log.DEBUG)
	} else {
		logger.SetLevel(log.ERROR)
	}

	Log = logger

}
