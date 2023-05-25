package listener

import (
	"context"

	"github.com/sirupsen/logrus"

	"notifier/driver"
	"notifier/helper"
	"notifier/models"
	"notifier/service"
)

const channelSize = 10000
const listenerChannel = "notification-listener"

type Listener interface {
	Run()
}

type listener struct {
	logger  *logrus.Logger
	helper  helper.Helper
	driver  driver.Driver
	service service.RestService
}

func NewListener(l *logrus.Logger, d driver.Driver, h helper.Helper, s service.RestService) Listener {
	return &listener{
		logger:  l,
		driver:  d,
		helper:  h,
		service: s,
	}
}

func (l *listener) Run() {
	l.logger.Info("Running notification listener")
	channel := l.driver.Subscribe(context.Background(), listenerChannel)
	listener := channel.Channel(channelSize)
	for {
		select {
		case channelMessage := <-listener:
			var cm models.ChannelMessage
			l.helper.UnMarshal([]byte(channelMessage.Payload), &cm)
			switch cm.Medium {
			case sms:
				var channelSMS models.SMS
				l.helper.UnMarshal(cm.Notification, &channelSMS)
				go l.service.SendSMS(&channelSMS)
			case email:
				var channelEmail models.Email
				l.helper.UnMarshal(cm.Notification, &channelEmail)
				go l.service.SendEmail(&channelEmail)
			default:
				l.logger.Warnf("invalid notification type, please validate")
			}
		}
	}
}
