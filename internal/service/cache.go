package service

import ("sync"; "time")

type cacheItem struct{ value interface{}; expiresAt time.Time }
type CacheService struct{ mu sync.RWMutex; items map[string]cacheItem }

func NewCacheService() *CacheService {
	c := &CacheService{items: make(map[string]cacheItem)}
	go c.evict()
	return c
}

func (c *CacheService) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock(); defer c.mu.Unlock()
	c.items[key] = cacheItem{value, time.Now().Add(ttl)}
}

func (c *CacheService) Get(key string) (interface{}, bool) {
	c.mu.RLock(); defer c.mu.RUnlock()
	item, ok := c.items[key]
	if !ok || time.Now().After(item.expiresAt) { return nil, false }
	return item.value, true
}

func (c *CacheService) Delete(key string) {
	c.mu.Lock(); defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *CacheService) evict() {
	for range time.Tick(time.Minute) {
		c.mu.Lock()
		for k, v := range c.items { if time.Now().After(v.expiresAt) { delete(c.items, k) } }
		c.mu.Unlock()
	}
}
