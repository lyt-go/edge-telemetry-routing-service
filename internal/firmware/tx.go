package firmware

import (
	"fmt"
	"sync"
)

type Store struct {
	mu      sync.Mutex
	active  map[string]string
	failSet bool
}

func NewStore() *Store             { return &Store{active: make(map[string]string)} }
func (s *Store) SetFail(fail bool) { s.failSet = fail }

func (s *Store) Activate(deviceID, version string) error {
	s.mu.Lock()
	s.active[deviceID] = version
	s.mu.Unlock()
	if s.failSet {
		return fmt.Errorf("activation rejected")
	}
	return nil
}

func (s *Store) Version(deviceID string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.active[deviceID]
}
