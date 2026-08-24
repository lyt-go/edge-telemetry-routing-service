package gateway

import (
	"context"

	"edgetelemetry/internal/transport"
)

type Router struct {
	client *transport.Client
	first  context.Context
}

func NewRouter(client *transport.Client) *Router {
	return &Router{client: client}
}

func (r *Router) Dispatch(ctx context.Context, payload string) error {
	if r.first == nil {
		r.first = ctx
	}
	return r.client.Send(r.first, payload)
}
