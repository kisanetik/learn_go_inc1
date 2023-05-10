package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMethodEmptyPost(t *testing.T) {
	emptyBodyReq := httptest.NewRequest("POST", "/", nil)
	emptyBodyRes := httptest.NewRecorder()
	methodPost(emptyBodyRes, emptyBodyReq)
	if emptyBodyRes.Code != http.StatusBadRequest {
		t.Errorf("Ожидаемый статус код %v, получен %v", http.StatusBadRequest, emptyBodyRes.Code)
	}
}

func TestMethodOkPost(t *testing.T) {
	bodyReader := strings.NewReader("http://ya.ru")
	BodyReq := httptest.NewRequest("POST", "/", bodyReader)
	BodyRes := httptest.NewRecorder()
	methodPost(BodyRes, BodyReq)
	if BodyRes.Code != http.StatusCreated {
		t.Errorf("Ожидаемый статус код %v, получен %v", http.StatusCreated, BodyRes.Code)
	}
}
