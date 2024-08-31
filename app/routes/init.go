package routes

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/middlewares"
)

func (api *API) ViewsRouter() {

	file_service := api.router.NewRoute().Subrouter()
	file_service.Use(middlewares.SetContentTypeMiddleware)

	stylesDir := http.Dir("app/views/styles/")
	stylesHandler := http.StripPrefix("/styles/", http.FileServer(stylesDir))
	file_service.PathPrefix("/styles/").Handler(stylesHandler).Methods("GET")

	controllersDir := http.Dir("app/views/controllers/")
	controllersHandler := http.StripPrefix("/controllers/", http.FileServer(controllersDir))
	file_service.PathPrefix("/controllers/").Handler(controllersHandler).Methods("GET")

	file_service.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		filePath, err := filepath.Abs(filepath.Join("app", "views/register.html"))
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		fmt.Println("filepath: ", filePath)

		http.ServeFile(w, r, filePath)
	})

	file_service.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		filePath, err := filepath.Abs(filepath.Join("app", "views/login.html"))
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		fmt.Println("filepath: ", filePath)

		http.ServeFile(w, r, filePath)
	})

	file_service.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		filePath, err := filepath.Abs(filepath.Join("app", "views/updatepass.html"))
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		fmt.Println("filepath: ", filePath)

		http.ServeFile(w, r, filePath)
	})

	file_service.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		filePath, err := filepath.Abs(filepath.Join("app", "views/home.html"))
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		fmt.Println("filepath: ", filePath)

		http.ServeFile(w, r, filePath)
	})

}
