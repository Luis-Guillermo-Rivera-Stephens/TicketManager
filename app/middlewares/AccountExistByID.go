package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func AccountExistByID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verifying the ID")
		db, err := data.GetDataBase()
		if err != nil {
			http.Error(w, "Error getting the database", http.StatusInternalServerError)
			return
		}
		id, err := general.StringToInt(r.Header.Get("id"))
		if err != nil {
			http.Error(w, "Error getting the id", http.StatusBadRequest)
			return
		}

		result := db.Raw("SELECT dbo.ACCOUNT_EXIST_BY_ID(?)", id).Scan(&id)

		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		if id == 0 {
			http.Error(w, "Account does not exist", http.StatusNotFound)
			return
		}
		fmt.Println("Account exist")
		next.ServeHTTP(w, r)
	})
}
