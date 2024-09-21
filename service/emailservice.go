package service

import (
	"fmt"

	"github.com/rafaelq80/farmacia_go/config"
	"gopkg.in/mail.v2"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (es *EmailService) SendEmail(to, subject, body string) error {

	config.LoadAppConfig("config")

	m := mail.NewMessage()
	m.SetHeader("From", config.AppConfig.SmtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := mail.NewDialer(
		config.AppConfig.SmtpHost, 
		config.AppConfig.SmtpPort, 
		config.AppConfig.SmtpUser, 
		config.AppConfig.SmtpPassword,
	)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("Erro ao Enviar o E-mail: %w", err)
	}

	return nil
}
