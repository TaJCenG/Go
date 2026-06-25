package middleware

import (
	"bytes"
	"io"
	"net/http"
)

func JSONValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
			return
		}
		if r.Body == nil {
			http.Error(w, "Request body required", http.StatusBadRequest)
			return
		}
		// Peek body length
		b, _ := io.ReadAll(r.Body)
		if len(b) == 0 {
			http.Error(w, "Empty JSON body", http.StatusBadRequest)
			return
		}
		// Reset body so handler can read it again
		r.Body = io.NopCloser(bytes.NewReader(b))
		next.ServeHTTP(w, r)
	})
}
