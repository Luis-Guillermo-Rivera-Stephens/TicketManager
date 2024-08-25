package routes

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func (api *API) ViewsRouter() {

	/*api.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Construye la ruta completa al archivo home.html
		filePath, err := filepath.Abs(filepath.Join("..", "views", "home.html"))
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		// Envía el archivo home.html al cliente
		http.ServeFile(w, r, filePath)
	})*/

	api.router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Construye la ruta completa al archivo home.html
		filePath, err := filepath.Abs(filepath.Join("app", "views/login.html"))
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		fmt.Println("filepath: ", filePath)

		// Envía el archivo home.html al cliente
		http.ServeFile(w, r, filePath)
	})

}
