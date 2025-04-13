package messagebroker

import (
	"sync"
)

var (
	once           sync.Once
	globalRegistry *HandlerRegistry
)

func GlobalRegistry() *HandlerRegistry {
	once.Do(func() {
		globalRegistry = NewHandlerRegistry()
	})
	return globalRegistry
}
