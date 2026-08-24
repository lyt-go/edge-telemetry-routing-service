package snapshot

import (
	"sync"

	"edgetelemetry/internal/model"
)

type Store struct {
	mu      sync.RWMutex
	batches map[string][]model.Reading
}

func NewStore() *Store {
	return &Store{batches: make(map[string][]model.Reading)}
}

func (s *Store) Put(key string, readings []model.Reading) {
	copyOfReadings := append([]model.Reading(nil), readings...)
	s.mu.Lock()
	s.batches[key] = copyOfReadings
	s.mu.Unlock()
}

func (s *Store) Get(key string) []model.Reading {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]model.Reading(nil), s.batches[key]...)
}
