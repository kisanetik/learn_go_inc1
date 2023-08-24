package gzip

import (
	"bytes"
	gz "compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequest(t *testing.T) {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}

		if string(body) != "test body" {
			t.Fatalf("Request body is incorrect: %s", string(body))
		}

		w.WriteHeader(http.StatusOK)
	}

	buf := new(bytes.Buffer)
	gzWriter := gz.NewWriter(buf)
	_, err := gzWriter.Write([]byte("test body"))
	if err != nil {
		t.Fatalf("Failed to write to gzip writer: %v", err)
	}
	err = gzWriter.Close()
	if err != nil {
		t.Fatalf("Failed to close gzip writer: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/", buf)
	req.Header.Set("Content-Encoding", "gzip")

	rec := httptest.NewRecorder()

	Request(http.HandlerFunc(handlerFunc)).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}
