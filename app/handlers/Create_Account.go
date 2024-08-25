package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func Create_Account(w http.ResponseWriter, r *http.Request) {
	db, err := data.GetDataBase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var Account datatypes.ACCOUNT
	err = json.NewDecoder(r.Body).Decode(&Account)
	if err != nil {
		http.Error(w, "Error al decodificar el body", http.StatusInternalServerError)
	}

	Account.Passwd = r.Header.Get("password")

	if err = general.ValidAccountInfo(Account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(Account)
	var id int

	result := db.Raw("EXEC CREATE_ACCOUNT ?, ?, ?, ?", Account.Name, Account.Email, Account.Passwd, Account.Department_FK).Scan(&id)

	fmt.Println(id)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "No se pudo crear la cuenta", http.StatusInternalServerError)
		return
	}

	Account, err = general.GetAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Account)
}
