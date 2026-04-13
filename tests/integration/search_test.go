package integration

import ("net/http"; "net/http/httptest"; "testing")

func TestSearch(t *testing.T) {
	srv := newTestServer(t)
	token := loginTestUser(t, srv)
	req := newAuthRequest(t, "GET", "/api/v1/search?q=test&type=items", nil, token)
	rr := httptest.NewRecorder()
	srv.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
