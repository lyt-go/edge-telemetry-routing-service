package service

import (
	"fmt"

	"edgetelemetry/internal/audit"
	"edgetelemetry/internal/resource"
)

type CalibrationService struct {
	pool *resource.Pool
	log  *audit.Log
}

func NewCalibrationService(pool *resource.Pool, log *audit.Log) *CalibrationService {
	return &CalibrationService{pool: pool, log: log}
}

func (s *CalibrationService) Run(deviceIDs []string, failAt int) error {
	for index, id := range deviceIDs {
		handle, ok := s.pool.Acquire()
		if !ok {
			return fmt.Errorf("resource limit at %s", id)
		}
		if index == failAt {
			handle.Close()
			s.log.Add("failed:" + id)
			return fmt.Errorf("calibration failed: %s", id)
		}
		s.log.Add("calibrated:" + id)
		handle.Close()
	}
	return nil
}
