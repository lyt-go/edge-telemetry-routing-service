package service

import (
	"fmt"

	"edgetelemetry/internal/alert"
	"edgetelemetry/internal/model"
	"edgetelemetry/internal/sink"
)

type RetryService struct {
	store *alert.Store
	sink  *sink.Recorder
}

func NewRetryService(store *alert.Store, sink *sink.Recorder) *RetryService {
	return &RetryService{store: store, sink: sink}
}

func (s *RetryService) Complete(id string, version int) {
	next := model.Alert{ID: id, Version: version, State: "complete"}
	if s.store.Apply(next) {
		s.sink.Emit(fmt.Sprintf("%s:%d", id, version))
	}
}

func (s *RetryService) LateProgress(id string, version int) {
	s.store.Apply(model.Alert{ID: id, Version: version, State: "running"})
}
