package catalog

import (
	"sync"

	"edgetelemetry/internal/model"
)

type Cache struct {
	mu       sync.RWMutex
	profiles map[string]*model.DeviceProfile
}

func NewCache() *Cache { return &Cache{profiles: make(map[string]*model.DeviceProfile)} }

func (c *Cache) Put(profile *model.DeviceProfile) {
	if profile == nil {
		return
	}
	copyOfProfile := &model.DeviceProfile{ID: profile.ID, Ready: profile.Ready, Labels: make(map[string]string)}
	for key, value := range profile.Labels {
		copyOfProfile.Labels[key] = value
	}
	c.mu.Lock()
	c.profiles[profile.ID] = copyOfProfile
	c.mu.Unlock()
}

func (c *Cache) Get(id string) *model.DeviceProfile {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.profiles[id]
}
