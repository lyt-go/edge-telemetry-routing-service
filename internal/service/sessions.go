package service

import (
	"edgetelemetry/internal/notify"
	"edgetelemetry/internal/session"
)

type SessionService struct {
	pool   *session.Pool
	logger *notify.Logger
}

func NewSessionService(pool *session.Pool, logger *notify.Logger) *SessionService {
	return &SessionService{pool: pool, logger: logger}
}

func (s *SessionService) Handle(deviceID string, tags []string, ready <-chan struct{}, done chan<- struct{}) {
	ctx := s.pool.Acquire(deviceID, tags)
	s.pool.Release(ctx)
	go func() {
		<-ready
		s.logger.Record(*ctx)
		done <- struct{}{}
	}()
}
