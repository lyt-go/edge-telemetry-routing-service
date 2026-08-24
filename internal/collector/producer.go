package collector

import "fmt"

func Produce(inputs []string) (<-chan string, <-chan error) {
	results := make(chan string)
	errors := make(chan error, 1)
	go func() {
		for _, input := range inputs {
			if input == "bad" {
				errors <- fmt.Errorf("device read failed")
				return
			}
			results <- input
		}
		close(results)
		close(errors)
	}()
	return results, errors
}
