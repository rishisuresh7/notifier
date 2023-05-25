package factory

import (
	"log"
	"net/http"
	"net/smtp"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	"notifier/config"
	"notifier/constant"
	"notifier/driver"
	"notifier/helper"
	"notifier/listener"
	"notifier/repository"
	"notifier/request"
	"notifier/run"
	"notifier/service"
)

var redisSync sync.Once

type Factory interface {
	Rest() request.Rest
	Runner() run.Runner
	Driver() driver.Driver
}

type factory struct {
	logger    *logrus.Logger
	config    *config.Config
	redisConn *redis.Client
	httpClient *http.Client
	smtpAuth   smtp.Auth
}

func NewFactory(l *logrus.Logger, c *config.Config) Factory {
	return &factory{
		logger: l,
		config: c,
		httpClient: http.DefaultClient,
		smtpAuth: smtp.PlainAuth("", c.EmailUsername, c.EmailPassword, constant.SMTPHost),
	}
}

func (f *factory) Runner() run.Runner {
	return run.NewRunners(f.logger,
		listener.NewListener(f.logger, f.Driver(), f.Helper(), service.NewRestService(f.logger, f.Rest(), f.config, f.smtpAuth)),
	)
}

func (f *factory) Rest() request.Rest {
	return request.NewRestClient(f.httpClient)
}

func (f *factory) Helper() helper.Helper {
	return helper.NewHelper(f.Driver())
}

func (f *factory) RedisQueryer() repository.RedisQueryer {
	d, err := f.redisDriver()
	if err != nil {
		log.Fatalf("Unable to establish connection to redis: %s", err)
	}

	return repository.NewRedisQueryer(d)
}

func (f *factory) Driver() driver.Driver {
	return driver.NewDriver(f.RedisQueryer())
}