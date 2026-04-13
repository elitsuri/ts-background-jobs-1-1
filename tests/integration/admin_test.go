package integration

import ("net/http"; "net/http/httptest"; "testing")

func TestAdminListUsers(t *testing.T) {
	srv := newTestServer(t)
	token := loginAdminUser(t, srv)
	req := newAuthRequest(t, "GET", "/api/v1/admin/users", nil, token)
	rr := httptest.NewRecorder()
	srv.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rr.Code, rr.Body.String())
	}
}
