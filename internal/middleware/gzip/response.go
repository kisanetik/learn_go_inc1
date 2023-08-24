package gzip

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/kisanetik/learn_go_inc1/internal/logger"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Response(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			logger.Errorf("New writer level error: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer func() {
			if err = gz.Close(); err != nil {
				logger.Errorf("gzip.Response gz.Close() failed: %s", err)
			}
		}()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}
