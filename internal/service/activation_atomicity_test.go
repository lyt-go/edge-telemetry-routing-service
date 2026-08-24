package service_test

import (
	"testing"

	"edgetelemetry/internal/events"
	"edgetelemetry/internal/firmware"
	"edgetelemetry/internal/service"
)

func TestActivationFailureLeavesNoVersionOrSuccessEvent(t *testing.T) {
	store := firmware.NewStore()
	store.SetFail(true)
	bus := &events.Bus{}
	svc := service.NewActivationService(store, bus)
	if err := svc.Activate("sensor-a", "v2.4.0"); err == nil {
		t.Fatal("rejected firmware activation returned success")
	}
	if version := store.Version("sensor-a"); version != "" {
		t.Fatalf("rejected firmware activation left an active version: %q", version)
	}
	if published := bus.Events(); len(published) != 0 {
		t.Fatalf("rejected firmware activation published a success event: %#v", published)
	}
}
