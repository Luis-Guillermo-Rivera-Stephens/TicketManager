package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
)

func VerifyEmail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramEmail := r.Header.Get("email")
		res, err := VerifyEmailFunc(paramEmail)
		if res == -1 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		} else if res == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		r.Header.Set("id", strconv.Itoa(int(res)))

		next.ServeHTTP(w, r)
	})
}

func VerifyEmailFunc(email string) (int, error) {
	var res int
	db, err := data.GetDataBase()
	if err != nil {
		return -1, err
	}

	db.Raw("SELECT dbo.ACCOUNT_EXIST(?)", email).Scan(&res)

	if res == -1 {
		return 0, fmt.Errorf("cuenta inexistente")

	}

	return res, nil
}
