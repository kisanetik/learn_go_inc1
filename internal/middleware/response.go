package gzip

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (writer gzipWriter) Write(b []byte) (int, error) {
	return writer.Writer.Write(b)
}

func Response(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if !strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(response, request)
			return
		}

		gz, err := gzip.NewWriterLevel(response, gzip.BestSpeed)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			log.Fatal("Failed to read gzipped content!")

			return
		}
		defer func(gz *gzip.Writer) {
			err = gz.Close()
			if err != nil {
				log.Fatal("Failed to send gzipped chunk: %v", err)
			}
		}(gz)

		response.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: response, Writer: gz}, request)
	})
}
