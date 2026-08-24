package service

import (
	"fmt"

	"edgetelemetry/internal/events"
	"edgetelemetry/internal/firmware"
)

type ActivationService struct {
	store *firmware.Store
	bus   *events.Bus
}

func NewActivationService(store *firmware.Store, bus *events.Bus) *ActivationService {
	return &ActivationService{store: store, bus: bus}
}

func (s *ActivationService) Activate(deviceID, version string) error {
	if err := s.store.Activate(deviceID, version); err != nil {
		return err
	}
	if err := s.bus.Publish("activated:" + deviceID + ":" + version); err != nil {
		return fmt.Errorf("publish activation: %w", err)
	}
	return nil
}
