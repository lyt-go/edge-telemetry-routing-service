package notify

import (
	"sync"

	"edgetelemetry/internal/session"
)

type Logger struct {
	mu      sync.Mutex
	entries []string
}

func (l *Logger) Record(ctx session.Context) {
	l.mu.Lock()
	l.entries = append(l.entries, ctx.DeviceID)
	l.mu.Unlock()
}

func (l *Logger) Entries() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]string(nil), l.entries...)
}
