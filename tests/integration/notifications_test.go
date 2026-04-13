package integration

import ("bytes"; "encoding/json"; "net/http"; "net/http/httptest"; "testing")

func TestCreateNotification(t *testing.T) {
	srv := newTestServer(t)
	token := loginTestUser(t, srv)
	body, _ := json.Marshal(map[string]string{"title":"Test","body":"msg","type":"info"})
	req := newAuthRequest(t, "POST", "/api/v1/notifications", bytes.NewReader(body), token)
	rr := httptest.NewRecorder()
	srv.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rr.Code)
	}
}

func TestListNotifications(t *testing.T) {
	srv := newTestServer(t)
	token := loginTestUser(t, srv)
	req := newAuthRequest(t, "GET", "/api/v1/notifications", nil, token)
	rr := httptest.NewRecorder()
	srv.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
