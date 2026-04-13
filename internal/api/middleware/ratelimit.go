package middleware

import ("net/http"; "sync"; "time")

type windowEntry struct { count int; reset time.Time }
var (
	rateMu  sync.Mutex
	windows = map[string]*windowEntry{}
)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		rateMu.Lock()
		e, ok := windows[ip]
		if !ok || time.Now().After(e.reset) {
			windows[ip] = &windowEntry{1, time.Now().Add(time.Minute)}
		} else {
			e.count++
			if e.count > 100 {
				rateMu.Unlock()
				w.Header().Set("Retry-After", "60")
				http.Error(w, `{"error":"rate limit exceeded"}`, http.StatusTooManyRequests)
				return
			}
		}
		rateMu.Unlock()
		next.ServeHTTP(w, r)
	})
}
