package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
	"github.com/gorilla/mux"
)

func GetUnassignedTicketsByDepartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	departmentID, err := general.StringToInt(vars["d_id"])
	if err != nil {
		http.Error(w, "Department ID invalido", http.StatusBadRequest)
		return
	}

	db, err := data.GetDataBase()
	if err != nil {
		http.Error(w, "Database connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var tickets []datatypes.TICKET
	result := db.Raw("EXEC GET_UNASSIGNED_TICKETS_BY_DEPARTMENT ?", departmentID).Find(&tickets)

	if result.Error != nil {
		http.Error(w, "Database execution error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	token, err := general.GenerateJWT("G3tt1Ck3TsB_YD3p4r7MenT")
	if err != nil {
		http.Error(w, "Error generating token: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Token", token)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tickets)
}
