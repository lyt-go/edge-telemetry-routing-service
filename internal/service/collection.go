package service

import "edgetelemetry/internal/collector"

func Collect(inputs []string) ([]string, error) {
	results, errors := collector.Produce(inputs)
	var collected []string
	for value := range results {
		collected = append(collected, value)
	}
	if err := <-errors; err != nil {
		return collected, err
	}
	return collected, nil
}
