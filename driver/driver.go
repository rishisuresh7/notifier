package driver

import "notifier/repository"

type Driver interface {
	repository.RedisQueryer
}

type driver struct {
	repository.RedisQueryer
}

func NewDriver(rd repository.RedisQueryer) Driver {
	return &driver{
		RedisQueryer: rd,
	}
}
