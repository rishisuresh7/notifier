package factory

import (
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	"notifier/config"
	"notifier/driver"
	"notifier/helper"
	"notifier/repository"
	"notifier/run"
)

var redisSync sync.Once

type Factory interface {
	Runner() run.Runner
	Driver() driver.Driver
}

type factory struct {
	logger    *logrus.Logger
	config    *config.Config
	redisConn *redis.Client
}

func NewFactory(l *logrus.Logger, c *config.Config) Factory {
	return &factory{
		logger: l,
		config: c,
	}
}

func (f *factory) Runner() run.Runner {
	return run.NewRunners(f.logger)
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