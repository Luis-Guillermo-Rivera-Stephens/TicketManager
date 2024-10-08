package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"gorm.io/gorm"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
)

type API struct {
	ServerName string `json:"servername"`
	DB         *gorm.DB
	Port       int16 `json:"port"`
	router     *mux.Router
	Started    bool `json:"started"`
}

var ApiInfo API = API{"", nil, 0, nil, false}

func (api *API) Initialize(name string, port int16) (err error) {
	fmt.Println("API initialized")
	api.ServerName = name
	api.router = mux.NewRouter()
	fmt.Println("New router")
	api.Port = port
	api.DB, err = data.GetDataBase()

	if err != nil {
		fmt.Println("Error getting database")
		return err
	}
	fmt.Println("Database connected")
	api.Started = true
	return nil
}

func (api *API) Listen() {
	fmt.Println("Listening on port", api.Port)
	address := ":" + strconv.Itoa(int(api.Port))
	if err := http.ListenAndServe(address, api.router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func GetAPI() (API, error) {
	var err error = nil
	if !ApiInfo.Started {
		err = ApiInfo.Initialize("TicketManagerServer", 8080)
	}
	return ApiInfo, err
}
