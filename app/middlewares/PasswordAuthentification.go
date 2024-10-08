package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func PasswordAuthentification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var Pass NewPassword
		err := json.NewDecoder(r.Body).Decode(&Pass)
		if err != nil {
			http.Error(w, "Error trying to decode the body", http.StatusInternalServerError)
		}

		id, err := general.StringToInt(r.Header.Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		Account, err, stat := general.GetAccount(id)
		if err != nil {
			http.Error(w, err.Error(), stat)
			return
		}

		err = Account.VerifPassword(Pass.Password)
		if err != nil {
			http.Error(w, "La contraseña no es correcta", http.StatusUnauthorized)
			return
		}

		r.Header.Set("password", Pass.NewPassword)
		next.ServeHTTP(w, r)
	})
}

type NewPassword struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
