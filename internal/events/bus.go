package events

import "sync"

type Bus struct {
	mu     sync.Mutex
	events []string
	fail   bool
}

func (b *Bus) SetFail(fail bool) { b.fail = fail }

func (b *Bus) Publish(event string) error {
	if b.fail {
		return errPublish
	}
	b.mu.Lock()
	b.events = append(b.events, event)
	b.mu.Unlock()
	return nil
}

func (b *Bus) Events() []string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return append([]string(nil), b.events...)
}

type publishError string

func (e publishError) Error() string { return string(e) }

const errPublish publishError = "publish failed"
