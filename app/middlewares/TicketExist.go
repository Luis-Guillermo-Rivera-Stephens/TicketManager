package middlewares

import (
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func TicketExist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := data.GetDataBase()
		if err != nil {
			http.Error(w, "Error getting the database", http.StatusInternalServerError)
			return
		}
		id, err := general.StringToInt(r.Header.Get("id_ticket"))
		if err != nil {
			http.Error(w, "Error getting the id", http.StatusBadRequest)
			return
		}

		result := db.Raw("SELECT dbo.TICKET_EXIST(?)", id).Scan(&id)

		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		if id == 0 {
			http.Error(w, "Ticket does not exist", http.StatusNotFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
