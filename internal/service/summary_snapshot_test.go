package service_test

import (
	"testing"

	"edgetelemetry/internal/registry"
	"edgetelemetry/internal/service"
)

func TestSummaryUsesSnapshotFromRequestStart(t *testing.T) {
	store := registry.NewStore()
	store.Set("sensor-a", 1)
	store.Set("sensor-b", 2)
	svc := service.NewSummaryService(store)
	started := make(chan struct{})
	proceed := make(chan struct{})
	result := make(chan int)
	go func() { result <- svc.Total(started, proceed) }()
	<-started
	store.Set("sensor-c", 100)
	close(proceed)
	if total := <-result; total != 3 {
		t.Fatalf("in-flight telemetry summary included a later device update: total=%d", total)
	}
}
