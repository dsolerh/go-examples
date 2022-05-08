package gziphandler

import (
	"compress/gzip"
	"net/http"
)

type GzipResponseWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := w.Header()["Content-Type"]; !ok {
		// If the content type is not set, infer it from the uncompressed body.
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.gw.Write(b)
}

type DeflateResponseWriter struct {
	df *gzip.Writer
	http.ResponseWriter
}
