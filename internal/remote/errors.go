package remote

import "fmt"

type TemporaryError struct{ Message string }

func (e *TemporaryError) Error() string { return e.Message }

type RejectedError struct{ Message string }

func (e *RejectedError) Error() string { return e.Message }

func Normalize(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("remote delivery: %v", err)
}
