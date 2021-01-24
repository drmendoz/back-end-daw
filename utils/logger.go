package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type LocalFormatter struct {
	log.Formatter
}

var Log *log.Logger = log.New()

func init() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(1)
	}
	Log.SetFormatter(LocalFormatter{&log.TextFormatter{}})
	Log.SetOutput(file)
}

func (u LocalFormatter) Format(e *log.Entry) ([]byte, error) {
	e.Time = e.Time.Local()
	return u.Formatter.Format(e)
}
