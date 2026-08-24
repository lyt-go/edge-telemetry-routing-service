package service

import (
	"edgetelemetry/internal/catalog"
	"edgetelemetry/internal/device"
	"edgetelemetry/internal/model"
)

type ProfileService struct {
	builder device.Builder
	cache   *catalog.Cache
}

func NewProfileService(cache *catalog.Cache) *ProfileService {
	return &ProfileService{cache: cache}
}

func (s *ProfileService) Register(id string, labels map[string]string) error {
	profile, err := s.builder.Build(id, labels)
	s.cache.Put(profile)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProfileService) Lookup(id string) *model.DeviceProfile { return s.cache.Get(id) }
