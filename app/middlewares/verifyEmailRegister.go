package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func VerifyEmailRegister(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramEmail := r.Header.Get("email")
		fmt.Println(`Verifying email: ` + paramEmail)

		if paramEmail == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		res, err := VerifyEmailFunc(paramEmail)
		fmt.Println("result: ", res)
		if err != nil && res != -1 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		if res != 0 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("ya existe un usuario con ese email")
			return
		}

		next.ServeHTTP(w, r)
	})
}
