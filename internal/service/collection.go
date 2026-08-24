package service

import "edgetelemetry/internal/collector"

func Collect(inputs []string) ([]string, error) {
	results, errors := collector.Produce(inputs)
	var collected []string
	// Consume results and errors together. A device read failure must end the
	// request promptly and still return whatever readings arrived before it;
	// waiting on results alone would deadlock when the producer aborts.
	for {
		select {
		case value, ok := <-results:
			if !ok {
				// results channel was closed: drain any pending error.
				if err := <-errors; err != nil {
					return collected, err
				}
				return collected, nil
			}
			collected = append(collected, value)
		case err := <-errors:
			// An error aborts the run immediately. Drain any readings the
			// producer already buffered so they are not lost, then stop.
			for value := range results {
				collected = append(collected, value)
			}
			if err != nil {
				return collected, err
			}
			return collected, nil
		}
	}
}
