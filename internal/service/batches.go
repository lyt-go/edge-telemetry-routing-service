package service

import (
	"edgetelemetry/internal/exporter"
	"edgetelemetry/internal/ingest"
	"edgetelemetry/internal/model"
	"edgetelemetry/internal/snapshot"
)

type BatchService struct {
	parser *ingest.Parser
	store  *snapshot.Store
	queue  *exporter.Queue
}

func NewBatchService(parser *ingest.Parser, store *snapshot.Store, queue *exporter.Queue) *BatchService {
	return &BatchService{parser: parser, store: store, queue: queue}
}

func (s *BatchService) Submit(key string, input []model.Reading) []model.Reading {
	readings := s.parser.Parse(input)
	s.store.Put(key, readings)
	s.queue.Enqueue(key, readings)
	return append([]model.Reading(nil), readings...)
}

func (s *BatchService) Snapshot(key string) []model.Reading { return s.store.Get(key) }
func (s *BatchService) Export(key string) []model.Reading   { return s.queue.Take(key) }
