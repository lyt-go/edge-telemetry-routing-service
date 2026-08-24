package service

import "edgetelemetry/internal/collector"

func Collect(inputs []string) ([]string, error) {
	results, errors := collector.Produce(inputs)
	var collected []string
	for results != nil || errors != nil {
		select {
		case value, ok := <-results:
			if !ok {
				results = nil
				continue
			}
			collected = append(collected, value)
		case err, ok := <-errors:
			if !ok {
				errors = nil
				continue
			}
			if err != nil {
				return collected, err
			}
		}
	}
	return collected, nil
}
