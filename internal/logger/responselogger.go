package logger

import (
	"net/http"
)

type responseData struct {
	status int
	size   int
}

type loggResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (r *loggResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *loggResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func ResponseLogger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := loggResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		h.ServeHTTP(&lw, r)

		sugar.Infoln(
			"status code =", responseData.status,
			"| size body =", responseData.size,
		)
	}
}
