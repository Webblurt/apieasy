package apieasy

import (
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
}

func newResponseWriter(w http.ResponseWriter) ResponseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	return rw.ResponseWriter.Write(data)
}
