package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("token")
		if tokenString == "" {
			fmt.Println("No token provided")
			http.Error(w, "No token provided", http.StatusBadRequest)
			return
		}
		token, err := verifyJWTFunc(tokenString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}
		fmt.Println(token)

		next.ServeHTTP(w, r)
	})
}

func verifyJWTFunc(tokenString string) (*jwt.Token, error) {
	secretKey, err := general.GetJWT_sk()
	if err != nil {
		fmt.Println("Error al obtener la signing key")
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token is valid")
		fmt.Println("name:", claims["name"])
		fmt.Println("Authorized:", claims["authorized"])
		return token, nil
	}

	return nil, fmt.Errorf("invalid token")
}
