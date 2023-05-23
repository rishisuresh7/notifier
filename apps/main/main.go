package main

import (
	"github.com/sirupsen/logrus"
	ft "github.com/x-cray/logrus-prefixed-formatter"

	"notifier/config"
	"notifier/constant"
	"notifier/factory"
)

var Version = "0.0.0"

func main() {
	l := logrus.New()
	l.Level = logrus.DebugLevel
	l.Formatter = &ft.TextFormatter{
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: constant.TimeStampFormat,
	}

	conf, err := config.NewConfig()
	if err != nil {
		l.Fatalf("Error in generating config: %s", err)
	}

	f := factory.NewFactory(l, conf)
	l.Infof("Running notifier server version: %s", Version)
	runner := f.Runner()
	runner.Run()
}
