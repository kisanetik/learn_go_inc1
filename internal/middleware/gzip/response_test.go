package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponse(t *testing.T) {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, World!"))
	}

	handler := Response(http.HandlerFunc(handlerFunc))

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() failed: %v", err)
	}

	req.Header.Set("Accept-Encoding", "gzip")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "Hello, World!"
	reader, err := gzip.NewReader(bytes.NewReader(rr.Body.Bytes()))
	if err != nil {
		t.Fatalf("gzip.NewReader() failed: %v", err)
	}
	defer func(reader *gzip.Reader) {
		err := reader.Close()
		if err != nil {
			t.Fatalf("reader cannot closed")
		}
	}(reader)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, reader)
	if err != nil {
		t.Fatalf("io.Copy() failed: %v", err)
	}

	if buf.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", buf.String(), expected)
	}
}
