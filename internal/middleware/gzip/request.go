package gzip

import (
	"compress/gzip"
	"github.com/kisanetik/learn_go_inc1/internal/logger"
	"net/http"
	"strings"
)

func Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				logger.Errorf("new reader is error: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			r.Body = gz
			defer func() {
				if err = gz.Close(); err != nil {
					logger.Errorf("GzipRequest gz.Close() failed: %s", err)
				}
			}()
		}

		next.ServeHTTP(w, r)
	})
}
