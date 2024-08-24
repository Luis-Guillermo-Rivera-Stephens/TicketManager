package general

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"name":       name,
		"exp":        time.Now().Add(time.Minute * 120).Unix(),
	})
	secretKey, err := GetJWT_sk()
	if err != nil {
		fmt.Println("Error al obtener la signing key")
		return "", err
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error al generar el token")
		return "", err
	}

	return tokenString, nil
}
