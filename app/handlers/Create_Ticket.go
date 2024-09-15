package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func Create_Ticket(w http.ResponseWriter, r *http.Request) {
	db, err := data.GetDataBase()
	if err != nil {
		http.Error(w, "Database connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var Ticket datatypes.TICKET
	err = json.NewDecoder(r.Body).Decode(&Ticket)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Department: ", Ticket.Department)
	fmt.Println("Department ID: ", datatypes.DepartmentsHash[Ticket.Department])
	if err = general.ValidTicketInfo(Ticket); err != nil {
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(Ticket)

	result := db.Exec("EXEC CREATE_TICKET ?, ?, ?, ?, ?", Ticket.Title, Ticket.T_Description, datatypes.DepartmentsHash[Ticket.Department], time.Now().Format("2006-01-02 15:04:05.000"), Ticket.Priority)

	if result.Error != nil {
		http.Error(w, "Database execution error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "No rows affected, ticket not created", http.StatusInternalServerError)
		return
	}

	token, err := general.GenerateJWT(string(Ticket.Title))
	if err != nil {
		http.Error(w, "Error generating token: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Token", token)
	w.WriteHeader(http.StatusCreated)
}
