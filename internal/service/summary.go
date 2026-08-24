package service

import "edgetelemetry/internal/registry"

type SummaryService struct{ store *registry.Store }

func NewSummaryService(store *registry.Store) *SummaryService { return &SummaryService{store: store} }

func (s *SummaryService) Total(started chan<- struct{}, proceed <-chan struct{}) int {
	values := s.store.Snapshot()
	started <- struct{}{}
	<-proceed
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}
