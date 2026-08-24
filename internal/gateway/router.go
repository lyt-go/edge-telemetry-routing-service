package gateway

import (
	"context"

	"edgetelemetry/internal/transport"
)

type Router struct {
	client *transport.Client
}

func NewRouter(client *transport.Client) *Router {
	return &Router{client: client}
}

func (r *Router) Dispatch(ctx context.Context, payload string) error {
	return r.client.Send(ctx, payload)
}
