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
		fmt.Println("Verifying email: ", paramEmail)
		res, err := VerifyEmailFunc(paramEmail)
		if err != nil {
			if res == -1 {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err.Error())
				return
			} else if res == 0 {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
				return
			}
		}

		r.Header.Set("id", strconv.Itoa(int(res)))
		fmt.Println("Email verified")

		next.ServeHTTP(w, r)
	})
}

func VerifyEmailFunc(email string) (int, error) {
	db, err := data.GetDataBase()
	if err != nil {
		return -1, err
	}

	var res int

	db.Raw("SELECT dbo.ACCOUNT_EXIST(?)", email).Scan(&res)

	if res == -1 {
		return 0, fmt.Errorf("cuenta inexistente")
	}
	fmt.Println("Result: ", res)

	return res, nil
}
