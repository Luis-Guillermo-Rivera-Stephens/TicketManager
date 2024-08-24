package handlers

import (
	"fmt"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
)

func Delete_Account(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting user")
	db, err := data.GetDataBase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := r.Header.Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Usuario a eliminar: ", id)

	result := db.Exec("EXEC DELETE_ACCOUNT ?", id)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Usuario eliminado")
}
