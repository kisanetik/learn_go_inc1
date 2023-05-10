package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodPost(t *testing.T) {
	emptyBodyReq := httptest.NewRequest("POST", "/", nil)

	emptyBodyRes := httptest.NewRecorder()

	methodPost(emptyBodyRes, emptyBodyReq)

	if emptyBodyRes.Code != http.StatusBadRequest {
		t.Errorf("Ожидаемый статус код %v, получен %v", http.StatusBadRequest, emptyBodyRes.Code)
	}
}
