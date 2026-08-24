package resource

import "sync"

type Pool struct {
	mu       sync.Mutex
	open     int
	limit    int
	released int
}

type Handle struct {
	pool *Pool
	once sync.Once
}

func NewPool(limit int) *Pool { return &Pool{limit: limit} }

func (p *Pool) Acquire() (*Handle, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.open >= p.limit {
		return nil, false
	}
	p.open++
	return &Handle{pool: p}, true
}

func (h *Handle) Close() {
	h.once.Do(func() {
		h.pool.mu.Lock()
		h.pool.open--
		h.pool.released++
		h.pool.mu.Unlock()
	})
}

func (p *Pool) Stats() (open, released int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.open, p.released
}
