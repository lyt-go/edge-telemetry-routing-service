package sink

import "sync"

type Recorder struct {
	mu   sync.Mutex
	seen map[string]bool
	ids  []string
}

func NewRecorder() *Recorder { return &Recorder{seen: make(map[string]bool)} }

func (r *Recorder) Emit(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seen[key] = true
	r.ids = append(r.ids, key)
}

func (r *Recorder) IDs() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]string(nil), r.ids...)
}
