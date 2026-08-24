package audit

import "sync"

type Log struct {
	mu     sync.Mutex
	events []string
}

func (l *Log) Add(event string) {
	l.mu.Lock()
	l.events = append(l.events, event)
	l.mu.Unlock()
}

func (l *Log) Events() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]string(nil), l.events...)
}
