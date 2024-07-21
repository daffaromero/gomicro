package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			wrw := NewWrappedResponseWriter(w)
			wrw.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrw, r)
			defer wrw.Flush()

			return
		}

		next.ServeHTTP(w, r)
	})
}

type WrappedResponseWriter struct {
	write http.ResponseWriter
	gzipw *gzip.Writer
}

func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(w)

	return &WrappedResponseWriter{write: w, gzipw: gw}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.write.Header()
}

func (wr *WrappedResponseWriter) Write(b []byte) (int, error) {
	return wr.gzipw.Write(b)
}

func (wr *WrappedResponseWriter) WriteHeader(statusCode int) {
	wr.write.WriteHeader(statusCode)
}

func (wr *WrappedResponseWriter) Flush() {
	wr.gzipw.Flush()
	wr.gzipw.Close()
}
