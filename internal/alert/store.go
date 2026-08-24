package alert

import (
	"sync"

	"edgetelemetry/internal/model"
)

type Store struct {
	mu     sync.Mutex
	alerts map[string]model.Alert
}

func NewStore() *Store { return &Store{alerts: make(map[string]model.Alert)} }

func (s *Store) Apply(next model.Alert) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	current, exists := s.alerts[next.ID]
	if exists && next.Version < current.Version {
		return false
	}
	s.alerts[next.ID] = next
	return true
}

func (s *Store) Get(id string) model.Alert {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.alerts[id]
}
