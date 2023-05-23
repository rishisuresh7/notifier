package run

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type runner struct{
	logger  *logrus.Logger
	runners []Runner
}

func NewRunners(l *logrus.Logger, rs ...Runner) Runner {
	return &runner{
		logger:  l,
		runners: rs,
	}
}

func (s *runner) Run() {
	if len(s.runners) < 1 {
		s.logger.Errorf("No jobs to run, quitting server")
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	for _, rn := range s.runners {
		go rn.Run()
	}

	wg.Wait()
}