package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	//"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func Update_Owner(w http.ResponseWriter, r *http.Request) {

	id, err := general.StringToInt(r.Header.Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id_ticket, err := general.StringToInt(r.Header.Get("id_ticket"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(id, id_ticket)

	db, err := data.GetDataBase()
	if err != nil {
		http.Error(w, "Database connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(db)
	result := db.Exec("EXEC UPDATE_TICKET_OWNER ?, ?, ?", id_ticket, id, time.Now().Format("2006-01-02 15:04:05.000"))

	if result.Error != nil {
		http.Error(w, "Database execution error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "No rows affected, ticket not created", http.StatusInternalServerError)
		return
	}

	token, err := general.GenerateJWT(string("Update_Owner"))
	if err != nil {
		http.Error(w, "Error generating token: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Token", token)
	w.WriteHeader(http.StatusAccepted)
}
