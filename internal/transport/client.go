package transport

import (
	"context"
	"sync"
)

type Client struct {
	mu    sync.Mutex
	calls int
}

func (c *Client) Send(ctx context.Context, payload string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	c.mu.Lock()
	c.calls++
	c.mu.Unlock()
	return nil
}

func (c *Client) Calls() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.calls
}
