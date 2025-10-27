package logger

import (
	"github.com/carinfinin/risk-assessor/internal/config"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var instance *logrus.Logger
var once = sync.Once{}

func Configure(cfg *config.Config) error {
	var err error
	once.Do(func() {
		instance = logrus.New()
		level, parceErr := logrus.ParseLevel(cfg.LoggerLevel)
		if parceErr != nil {
			err = parceErr
			return
		}

		instance.SetLevel(level)
		instance.SetOutput(os.Stdout)

		switch cfg.Format {
		case "json":
			instance.Formatter = &logrus.JSONFormatter{}
		default:
			instance.Formatter = &logrus.TextFormatter{}
		}
	})
	return err
}

func Get() (*logrus.Logger, error) {
	if instance == nil {
		err := Configure(&config.Config{
			Format:      "text",
			LoggerLevel: "info",
		})
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}
