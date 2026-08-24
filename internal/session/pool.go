package session

import "sync"

type Context struct {
	DeviceID string
	Tags     []string
}

type Pool struct{ pool sync.Pool }

func NewPool() *Pool {
	p := &Pool{}
	p.pool.New = func() any { return &Context{} }
	return p
}

func (p *Pool) Acquire(deviceID string, tags []string) *Context {
	ctx := p.pool.Get().(*Context)
	ctx.DeviceID = deviceID
	ctx.Tags = append(ctx.Tags[:0], tags...)
	return ctx
}

func (p *Pool) Release(ctx *Context) {
	p.pool.Put(ctx)
}
