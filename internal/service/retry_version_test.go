package service_test

import (
	"testing"

	"edgetelemetry/internal/alert"
	"edgetelemetry/internal/service"
	"edgetelemetry/internal/sink"
)

func TestLateAttemptCannotRegressCompletedAlert(t *testing.T) {
	store := alert.NewStore()
	recorder := sink.NewRecorder()
	svc := service.NewRetryService(store, recorder)
	svc.Complete("alert-7", 2)
	svc.LateProgress("alert-7", 1)
	if got := store.Get("alert-7"); got.State != "complete" || got.Version != 2 {
		t.Errorf("late alert attempt replaced the completed state: %#v", got)
	}
	svc.Complete("alert-7", 2)
	if ids := recorder.IDs(); len(ids) != 1 {
		t.Fatalf("completed alert side effect ran more than once: %#v", ids)
	}
}
