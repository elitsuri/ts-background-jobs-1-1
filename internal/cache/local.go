package cache

// Thread-safe in-memory cache with TTL.
import ("sync"; "time")

type entry struct{ v interface{}; exp time.Time }
type LocalCache struct{ mu sync.RWMutex; m map[string]entry }

func NewLocalCache() *LocalCache {
	c := &LocalCache{m: make(map[string]entry)}
	go func() {
		for range time.Tick(time.Minute) {
			c.mu.Lock()
			for k, v := range c.m { if time.Now().After(v.exp) { delete(c.m, k) } }
			c.mu.Unlock()
		}
	}()
	return c
}

func (c *LocalCache) Set(k string, v interface{}, ttl time.Duration) {
	c.mu.Lock(); defer c.mu.Unlock()
	c.m[k] = entry{v, time.Now().Add(ttl)}
}

func (c *LocalCache) Get(k string) (interface{}, bool) {
	c.mu.RLock(); defer c.mu.RUnlock()
	e, ok := c.m[k]
	if !ok || time.Now().After(e.exp) { return nil, false }
	return e.v, true
}

func (c *LocalCache) Flush() { c.mu.Lock(); c.m = make(map[string]entry); c.mu.Unlock() }
