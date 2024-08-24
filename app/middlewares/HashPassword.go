package middlewares

import (
	"fmt"

	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hashing password")
		password := r.Header.Get("password")

		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r.Header.Set("password", string(hashed))
		fmt.Println("Password hashed")
		next.ServeHTTP(w, r)
	})
}
