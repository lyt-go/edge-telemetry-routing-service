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

// Apply merges next into the store using alert-state-machine semantics.
//
// A terminal state ("complete") is sticky: once an alert reaches it, later
// attempts cannot regress it. This is what protects against a late callback
// from an earlier retry attempt overwriting a completion that a later attempt
// already produced.
//
// Apply returns true only when the stored state actually changed as a result
// of this call. Callers gate side effects on that return value, so a no-op
// re-application (e.g. completing an already-complete alert) does not repeat
// the side effect.
func (s *Store) Apply(next model.Alert) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	cur, ok := s.alerts[next.ID]

	// A completion wins outright: record it and report the transition. A
	// repeat completion of an already-complete alert at the same version is
	// not a change, so it returns false and no side effect repeats.
	if next.State == "complete" {
		if ok && cur.State == "complete" {
			return false
		}
		s.alerts[next.ID] = next
		return true
	}

	// Non-terminal updates cannot displace a terminal state. A late progress
	// callback from an earlier attempt must not regress an alert that a later
	// attempt has already completed.
	if ok && cur.State == "complete" {
		return false
	}

	// Otherwise the first write, or a genuine state change, is applied.
	if ok && cur.State == next.State {
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
