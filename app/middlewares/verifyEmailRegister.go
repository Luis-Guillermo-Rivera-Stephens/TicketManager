package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func VerifyEmailRegister(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramEmail := r.Header.Get("email")

		res, err := VerifyEmailFunc(paramEmail)
		if res != -1 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		if res != 0 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(fmt.Errorf("ya existe un usuario con ese email"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
