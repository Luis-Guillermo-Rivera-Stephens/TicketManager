package middlewares

import (
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
	"github.com/gorilla/mux"
)

func ValidStatus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		statusID, err := general.StringToInt(vars["s_id"])
		if err != nil {
			http.Error(w, "Status ID invalido", http.StatusBadRequest)
			return
		}
		if statusID == 0 || statusID > len(datatypes.StatusSlice)-1 {
			http.Error(w, "Status ID invalido", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
