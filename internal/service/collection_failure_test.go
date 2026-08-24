package service_test

import (
	"testing"
	"time"

	"edgetelemetry/internal/service"
)

func TestCollectionReturnsPartialReadingsAndError(t *testing.T) {
	type outcome struct {
		values []string
		err    error
	}
	done := make(chan outcome, 1)
	go func() {
		values, err := service.Collect([]string{"sensor-a", "bad", "sensor-c"})
		done <- outcome{values: values, err: err}
	}()
	select {
	case got := <-done:
		if got.err == nil {
			t.Fatal("failed device collection returned without its error")
		}
		if len(got.values) != 1 || got.values[0] != "sensor-a" {
			t.Fatalf("partial device readings changed: %#v", got.values)
		}
	case <-time.After(300 * time.Millisecond):
		t.Fatal("failed device collection stayed blocked waiting for results")
	}
}
