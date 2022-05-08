package gziphandler

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
	next http.Handler
}

func (h *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodings := r.Header.Get("Accept-Encoding")
	if strings.Contains(encodings, "gzip") {
		h.serveGzipped(w, r)
	} else if strings.Contains(encodings, "deflate") {
		panic("Deflate not implemented")
	} else {
		h.servePlain(w, r)
	}
}

func (h *GzipHandler) serveGzipped(w http.ResponseWriter, r *http.Request) {
	gzw := gzip.NewWriter(w)
	defer gzw.Close()
	w.Header().Set("Content-Encoding", "gzip")
	h.next.ServeHTTP(GzipResponseWriter{gzw, w}, r)
}

func (h *GzipHandler) servePlain(w http.ResponseWriter, r *http.Request) {
	h.next.ServeHTTP(w, r)
}
