package data

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDataBase() (*gorm.DB, error) {
	var err error
	var url string
	if DB == nil {
		url, err = getDataBaseUrl()
		if err != nil {
			fmt.Println("No se pudo obtener el link de la base de datos")
			return nil, err
		}

		DB, err := gorm.Open(sqlserver.Open(url), &gorm.Config{})

		if err != nil {
			fmt.Println("Error connecting to database")
			return nil, err
		}

		return DB, nil
	}

	sqls, err := DB.DB()

	if err != nil {
		fmt.Println("Error al comprobar la conexion de la base de datos")
		return nil, err
	}

	if err = sqls.Ping(); err == nil {
		fmt.Println("DB is already conected")
	}

	return DB, nil
}

func getDataBaseUrl() (string, error) {
	err := godotenv.Load("C:\\Users\\memor\\OneDrive\\Escritorio\\programacionPersonal\\TicketManager\\app\\data\\DB.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		fmt.Println("No se encontro el url")
		return "", err
	}

	return databaseURL, nil
}
