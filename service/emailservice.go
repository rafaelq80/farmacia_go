package service

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/rafaelq80/farmacia_go/config"
	"github.com/rafaelq80/farmacia_go/util"
	"gopkg.in/mail.v2"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

type EmailData struct {
	Name    string
	Message string
}

func (es *EmailService) SendEmail(to, name, subject string) error {

	data := EmailData{
		Name:    name,
		Message: "O seu cadastro foi efetuado com sucesso!",
	}

	// Parse the template
	tmpl, err := template.New("emailTemplate").Parse(util.EmailTemplate)
	if err != nil {
		return fmt.Errorf("erro ao analisar o template: %w", err)
	}

	// Execute the template with the data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("erro ao executar o template: %w", err)
	}

	m := mail.NewMessage()
	m.SetHeader("From", config.AppConfig.SmtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := mail.NewDialer(
		config.AppConfig.SmtpHost,
		config.AppConfig.SmtpPort,
		config.AppConfig.SmtpUser,
		config.AppConfig.SmtpPassword,
	)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("erro ao enviar o e-mail: %w", err)
	}

	return nil
}
