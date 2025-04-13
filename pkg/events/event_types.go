package events

const (
	EmailVerificationEvent = "email.verification"
	NotificationSendEvent  = "notification.send"
)

type Event struct {
	Type     string
	Payload  []byte
	Metadata map[string]interface{}
}
