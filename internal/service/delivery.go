package service

import (
	"edgetelemetry/internal/remote"
	"edgetelemetry/internal/repository"
)

type DeliveryService struct{ store *repository.TxStore }

func NewDeliveryService(store *repository.TxStore) *DeliveryService {
	return &DeliveryService{store: store}
}

func (s *DeliveryService) Deliver(value string, send func() error) error {
	return s.store.Record(value, func() error {
		var last error
		for attempt := 0; attempt < 2; attempt++ {
			last = remote.Normalize(send())
			if last == nil {
				return nil
			}
		}
		return last
	})
}
