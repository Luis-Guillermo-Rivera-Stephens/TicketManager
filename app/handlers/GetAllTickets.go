package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	db, err := data.GetDataBase()
	if err != nil {
		http.Error(w, "Database connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var tickets []datatypes.TICKET
	result := db.Raw("EXEC GET_ALL_TICKETS").Find(&tickets)

	if result.Error != nil {
		http.Error(w, "Database execution error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	token, err := general.GenerateJWT("G3t4L1t1Ck3Ts")
	if err != nil {
		http.Error(w, "Error generating token: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Token", token)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tickets)
}
