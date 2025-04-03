package port

import "context"

type EmailSender interface {
	Send(ctx context.Context, to, subject, body string) error
}

type EmailService interface {
	SendVerificationEmail(ctx context.Context, email, token string) error
}
