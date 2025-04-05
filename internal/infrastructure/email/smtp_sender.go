package email

import (
	"context"
	"fmt"

	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"go.uber.org/zap"
	"gopkg.in/mail.v2"
)

type MailSender struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func NewMailSender(host string, port int, username, password, from string) *MailSender {
	return &MailSender{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (s *MailSender) Send(ctx context.Context, to, subject, body string) error {
	m := mail.NewMessage()

	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)

	d := mail.NewDialer(s.host, s.port, s.username, s.password)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err := d.DialAndSend(m); err != nil {
		logger.Log.Error("failed to send mail", zap.Error(err))
		return fmt.Errorf("failed to send email: %w", err)
	}

	logger.Log.Info("successfully to send mail, detail :", zap.String("to :", to), zap.String("subject", subject))
	return nil
}
