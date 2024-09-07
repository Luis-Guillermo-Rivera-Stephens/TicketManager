package middlewares

import (
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
	"github.com/gorilla/mux"
)

func VerifyDepartment(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		departmentID, err := general.StringToInt(vars["d_id"])
		if err != nil {
			http.Error(w, "Department ID invalido", http.StatusBadRequest)
			return
		}
		if departmentID == 0 || departmentID > len(datatypes.DepartmentsSlice)-1 {
			http.Error(w, "Department ID invalido", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
