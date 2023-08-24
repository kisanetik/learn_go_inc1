package logger

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseLogger(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("test response"))
		if err != nil {
			t.Errorf("have error in w.Write: %d", err)
		}
	}

	ResponseLogger(handlerFunc)(rr, req)
}
