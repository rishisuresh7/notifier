package service

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/sirupsen/logrus"

	"notifier/config"
	"notifier/constant"
	"notifier/models"
	"notifier/request"
)

type RestService interface {
	SendSMS(sms *models.SMS)
	SendEmail(email *models.Email)
}

type restService struct {
	logger *logrus.Logger
	client request.Rest
	config *config.Config
	smtpAuth smtp.Auth
}

func NewRestService(l *logrus.Logger, c request.Rest, conf *config.Config, auth smtp.Auth) RestService {
	return &restService{
		logger: l,
		client: c,
		config: conf,
		smtpAuth: auth,
	}
}

func (r *restService) SendSMS(sms *models.SMS) {
	var (
		phones 		 []string
		maskedPhones []string
	)
	for _, phone := range sms.To {
		trimmedPhone := strings.TrimPrefix(phone, "+91")
		phones = append(phones, trimmedPhone)
		maskedPhones = append(maskedPhones, strings.Replace(trimmedPhone, trimmedPhone[0:6], "xxxxxx", -1))
	}

	phoneString := strings.Join(phones, ",")
	url := fmt.Sprintf(urlTemplate, r.config.SMSApiKey, sms.Message, phoneString)
	// TODO: Parse the response to track request IDs
	res, err := r.client.Get(url, nil, nil)
	if err != nil {
		// TODO: implement retries for sending requests
		r.logger.Errorf("sendSMS: unable to send a SMS request: %s", err)
		return
	}

	if res != 200 {
		r.logger.Errorf("sendSMS: unable to send a SMS request: unexpected error happened in sending the request")
		return
	}

	r.logger.Debugf("sent message to %s successfully", strings.Join(maskedPhones, ","))
}

func (r *restService) SendEmail(email *models.Email) {
	err := smtp.SendMail(fmt.Sprintf("%s:%s", constant.SMTPHost, constant.SMTPPort), r.smtpAuth, email.From, email.To, email.Message)
	if err != nil {
		r.logger.Errorf("sendEmail: unable to send email: %s", err)
		return
	}

	r.logger.Debugf("sent message to %s successfully", strings.Join(email.To, ","))
}
