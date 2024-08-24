package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/general"
)

func PasswordAuthentification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var Pass NewPassword
		json.NewDecoder(r.Body).Decode(&Pass)

		id, err := general.StringToInt(r.Header.Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		User, errRes := general.GetAccount(id)
		if errRes != nil {
			http.Error(w, errRes.Message, http.StatusUnauthorized)
			return
		}

		err = User.VerifPassword(Pass.Password)
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
