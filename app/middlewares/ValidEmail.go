package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ValidEmail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramEmail := r.Header.Get("email")
		if paramEmail == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Necesitas un email para ingresar")
			return
		}

		atIndex := strings.Index(paramEmail, "@")
		if atIndex == -1 {
			// No se encontró el carácter '@'
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Email invalido: falta el @")
			return
		}

		// Buscar el punto después del carácter '@'
		dotIndex := strings.Index(paramEmail[atIndex:], ".")
		if dotIndex == -1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Email invalido: falta el .")
			return

		}

		// Verificar que el punto no esté justo después del '@' (debe haber al menos un carácter entre '@' y '.')
		if dotIndex <= 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Email invalido: no hay dominio")
			return
		}
		fmt.Println("Email válido")
		next.ServeHTTP(w, r)
	})
}
