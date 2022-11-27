package middleware

import "net/http"

func FileSizeLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)
		next.ServeHTTP(w, r)
	})
}
