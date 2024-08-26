package middlewares

import (
	"net/http"
	"strings"
)

func SetContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") {

			w.Header().Set("Content-Type", "text/css")

		} else if strings.HasSuffix(r.URL.Path, ".js") {

			w.Header().Set("Content-Type", "text/javascript")
		}
		next.ServeHTTP(w, r)
	})
}
