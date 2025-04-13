package email

import (
	"github.com/livingdolls/go-template/internal/infrastructure/messagebroker"
	"github.com/livingdolls/go-template/pkg/events"
)

func init() {
	messagebroker.GlobalRegistry().RegisterHandler(events.EmailVerificationEvent, NewEmailHandler())
}
