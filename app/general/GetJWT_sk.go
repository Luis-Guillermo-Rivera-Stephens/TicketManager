package general

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetJWT_sk() (string, error) {
	err := godotenv.Load("./JWT.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}
	JWT_SK := os.Getenv("JWT_SECRET_KEY")
	if JWT_SK == "" {
		fmt.Println("No se encontro el url")
		return "", err
	}

	return JWT_SK, nil
}
