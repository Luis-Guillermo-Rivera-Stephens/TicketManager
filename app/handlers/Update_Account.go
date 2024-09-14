package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func Update_Password(w http.ResponseWriter, r *http.Request) {
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

	password := r.Header.Get("password")
	if password == "" {
		http.Error(w, "password is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Usuario a actualizar: ", id)

	var resultID int

	result := db.Raw("EXEC UPDATE_ACCOUNT_PASSWORD ?, ?", id, password).Scan(&resultID)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	Account, err, stat := general.GetAccount(resultID)
	if err != nil {
		http.Error(w, err.Error(), stat)
		return
	}

	token, err := general.GenerateJWT(string(Account.Name + Account.Passwd))
	if err != nil {
		http.Error(w, "Error creating the token", http.StatusBadRequest)
		return
	}

	fmt.Println(Account)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Token", token)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Account)
}
