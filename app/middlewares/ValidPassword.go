package middlewares

import (
	"fmt"

	"net/http"
	"strings"
)

func ValidPassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		password := r.Header.Get("password")
		fmt.Println("Validating password: ", password)
		paslen := len(password)

		if paslen == 0 {
			http.Error(w, "Password is required", http.StatusBadRequest)
			return
		} else if paslen < 10 {
			http.Error(w, "Password must be at least 10 characters", http.StatusBadRequest)
			return
		}
		if strings.Contains(password, "-") {
			http.Error(w, "Password must not contain '-'", http.StatusBadRequest)
			return
		}

		fmt.Println("Password is ok")
		next.ServeHTTP(w, r)
	})
}
