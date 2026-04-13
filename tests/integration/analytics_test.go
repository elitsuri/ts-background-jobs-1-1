package integration

import ("net/http"; "net/http/httptest"; "testing")

func TestAnalyticsOverview(t *testing.T) {
	srv := newTestServer(t)
	token := loginTestUser(t, srv)
	req := newAuthRequest(t, "GET", "/api/v1/analytics/overview", nil, token)
	rr := httptest.NewRecorder()
	srv.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rr.Code, rr.Body.String())
	}
}

func TestAnalyticsTimeseries(t *testing.T) {
	srv := newTestServer(t)
	token := loginTestUser(t, srv)
	req := newAuthRequest(t, "GET", "/api/v1/analytics/timeseries?days=7", nil, token)
	rr := httptest.NewRecorder()
	srv.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
