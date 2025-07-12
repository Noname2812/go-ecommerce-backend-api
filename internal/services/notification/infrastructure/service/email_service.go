package notificationserviceimpl

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"

	notificationservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/application/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
)

const (
	TEMPLATE_EMAIL_OTP = "otp-auth.html"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}
type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

type emailService struct {
	logger *zap.Logger
	config *setting.EmailSetting
}

// SendRegisterOTP implements notificationservice.EmailService.
func (e *emailService) SendRegisterOTP(email string, otp string) error {
	htmlBody, err := getMailTemplate(TEMPLATE_EMAIL_OTP, map[string]interface{}{
		"otp": otp,
	})
	if err != nil {
		e.logger.Error("failed to get mail template", zap.Error(err))
		return err
	}

	contentEmail := Mail{
		From:    EmailAddress{Address: e.config.Username, Name: "SHOP DEV"},
		To:      []string{email},
		Subject: "OTP Verification",
		Body:    htmlBody,
	}

	messageMail := buildMessage(contentEmail)
	auth := smtp.PlainAuth("", e.config.Username, e.config.Password, e.config.Host)
	addr := fmt.Sprintf("%s:%d", e.config.Host, e.config.Port)

	err = send(addr, auth, []string{email}, e.config.Username, messageMail)
	if err != nil {
		e.logger.Error("failed to send email", zap.Error(err))
		return err
	}

	e.logger.Info("SendRegisterOTP", zap.String("email", email), zap.String("otp", otp))

	return nil
}

func NewNotificationService(config *setting.EmailSetting, logger *zap.Logger) notificationservice.EmailService {
	return &emailService{
		logger: logger,
		config: config,
	}
}

// ------------ utils ------------

// send email
func send(addr string, auth smtp.Auth, to []string, from string, messageMail string) error {
	err := smtp.SendMail(addr, auth, from, to, []byte(messageMail))
	return err
}

// get mail template
func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates-email/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

func buildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
