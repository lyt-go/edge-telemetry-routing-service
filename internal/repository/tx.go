package repository

import "sync"

type TxStore struct {
	mu      sync.Mutex
	entries []string
}

func (s *TxStore) Record(value string, work func() error) error {
	if err := work(); err != nil {
		return err
	}
	s.mu.Lock()
	s.entries = append(s.entries, value)
	s.mu.Unlock()
	return nil
}

func (s *TxStore) Entries() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.entries...)
}
