package logger

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

var sugar = zap.NewExample().Sugar()

func RequestLogger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)

		duration := time.Since(start)
		sugar.Infoln(
			"path =", r.RequestURI,
			"| method =", r.Method,
			"| time duration =", duration,
		)
	}
}
