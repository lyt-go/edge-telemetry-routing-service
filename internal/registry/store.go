package registry

import "sync"

type Store struct {
	mu     sync.RWMutex
	values map[string]int
}

func NewStore() *Store { return &Store{values: make(map[string]int)} }

func (s *Store) Set(id string, value int) {
	s.mu.Lock()
	s.values[id] = value
	s.mu.Unlock()
}

func (s *Store) Snapshot() map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.values
}
