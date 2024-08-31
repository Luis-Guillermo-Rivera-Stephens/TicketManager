package middlewares

import (
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func IsPM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := general.StringToInt(r.Header.Get("id"))
		if err != nil {
			http.Error(w, "Error getting the id", http.StatusBadRequest)
			return
		}

		Account, err := general.GetAccount(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !Account.IsPM {
			http.Error(w, "You are not a PM", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
