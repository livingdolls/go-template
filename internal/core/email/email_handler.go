package email

import (
	"context"
	"encoding/json"

	"github.com/livingdolls/go-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type EmailHandler struct{}

func NewEmailHandler() *EmailHandler {
	return &EmailHandler{}
}

func (h *EmailHandler) Handle(ctx context.Context, payload []byte) error {
	var message struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}

	if err := json.Unmarshal(payload, &message); err != nil {
		return err
	}

	logger.Log.Info("received payload", zap.String(": ", string(payload)))

	return nil
}
