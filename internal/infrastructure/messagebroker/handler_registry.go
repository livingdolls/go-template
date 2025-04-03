package messagebroker

import (
	"sync"

	"github.com/livingdolls/go-template/internal/core/port"
)

type HandlerRegistry struct {
	handlers map[string]port.MessageHandler
	mu       sync.RWMutex
}

func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		handlers: make(map[string]port.MessageHandler),
	}
}

func (r *HandlerRegistry) RegisterHandler(messageType string, handler port.MessageHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.handlers[messageType] = handler
}

func (r *HandlerRegistry) GetHandler(messageType string) (port.MessageHandler, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	handler, exists := r.handlers[messageType]
	return handler, exists
}

func (r *HandlerRegistry) UnregisterHandler(messageType string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.handlers, messageType)
}

func (r *HandlerRegistry) GetAllHandlers() map[string]port.MessageHandler {
	r.mu.Lock()
	defer r.mu.Unlock()

	handlersCopy := make(map[string]port.MessageHandler)
	for key, value := range r.handlers {
		handlersCopy[key] = value
	}

	return handlersCopy
}
