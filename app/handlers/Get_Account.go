package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	id, err := strconv.Atoi(r.Header.Get("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	Account, err := general.GetAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pass := r.Header.Get("password")

	err = Account.VerifPassword(pass)
	if err != nil {
		http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
		return
	}

	token, err := general.GenerateJWT(string(Account.Name + Account.Passwd))
	if err != nil {
		http.Error(w, "Error creating the token", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("token", token)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Account)
}
