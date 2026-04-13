package cache_test

import ("testing"; "time")
"github.com/example/ts-background-jobs-1/internal/service"

func TestCacheSetGet(t *testing.T) {
	c := service.NewCacheService()
	c.Set("key", "value", time.Minute)
	v, ok := c.Get("key")
	if !ok { t.Fatal("expected hit") }
	if v != "value" { t.Fatalf("expected value, got %v", v) }
}

func TestCacheTTL(t *testing.T) {
	c := service.NewCacheService()
	c.Set("k", "v", time.Millisecond)
	time.Sleep(5 * time.Millisecond)
	_, ok := c.Get("k")
	if ok { t.Fatal("expected cache miss after TTL") }
}

func TestCacheDelete(t *testing.T) {
	c := service.NewCacheService()
	c.Set("k", "v", time.Minute)
	c.Delete("k")
	_, ok := c.Get("k")
	if ok { t.Fatal("expected miss after delete") }
}
