package cronjob

import (
	"github.com/sirupsen/logrus"
)

type CronJob interface {
	Name() string
	Schedule() string
	Job(l *logrus.Logger) func()
}

func NewCronJobs() []CronJob {
	return []CronJob{
	}
}
