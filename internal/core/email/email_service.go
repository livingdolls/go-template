package email

import (
	"context"

	"github.com/livingdolls/go-template/internal/core/port"
	"github.com/livingdolls/go-template/pkg/events"
)

type EmailService struct {
	publisher port.MessagePublisher
}

func NewEmailService(publisher port.MessagePublisher) *EmailService {
	return &EmailService{publisher: publisher}
}

func (s *EmailService) SendEmailVerification(ctx context.Context, email, token string) error {
	payload := map[string]interface{}{
		"email": email,
		"token": token,
	}

	return s.publisher.Publish(
		ctx,
		"notifications",
		events.EmailVerificationEvent,
		payload,
	)
}
