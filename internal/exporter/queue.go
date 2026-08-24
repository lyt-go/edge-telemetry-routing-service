package exporter

import (
	"sync"

	"edgetelemetry/internal/model"
)

type Queue struct {
	mu      sync.Mutex
	pending map[string][]model.Reading
}

func NewQueue() *Queue {
	return &Queue{pending: make(map[string][]model.Reading)}
}

func (q *Queue) Enqueue(key string, readings []model.Reading) {
	q.mu.Lock()
	q.pending[key] = readings
	q.mu.Unlock()
}

func (q *Queue) Take(key string) []model.Reading {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.pending[key]
}
