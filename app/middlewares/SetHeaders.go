package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SetHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Setting headers")

		// Leer el cuerpo del request
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error al leer el Body del request", http.StatusInternalServerError)
			return
		}

		// Restablecer el cuerpo del request para que pueda ser leído nuevamente
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		var data map[string]interface{}
		if err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&data); err != nil {
			http.Error(w, "Error decodificando el bpdy", http.StatusInternalServerError)
			return
		}

		email, email_ok := data["email"].(string)
		password, pass_ok := data["password"].(string)

		if !email_ok || !pass_ok {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		r.Header.Set("email", email)
		r.Header.Set("password", password)

		fmt.Println("Headers ready")
		fmt.Println("Email: ", r.Header.Get("email"))
		fmt.Println("Password: ", r.Header.Get("password"))
		next.ServeHTTP(w, r)
	})
}
