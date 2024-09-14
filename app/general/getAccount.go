package general

import (
	"fmt"

	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
)

func GetAccount(id int) (datatypes.ACCOUNT, error, int) {
	db, err := data.GetDataBase()
	if err != nil {
		return datatypes.ACCOUNT{}, err, http.StatusInternalServerError
	}
	var User datatypes.ACCOUNT
	result := db.Raw("EXEC GET_ACCOUNT ?", id).Scan(&User)

	if result.Error != nil {
		return datatypes.ACCOUNT{}, result.Error, http.StatusInternalServerError
	}
	if result.RowsAffected == 0 {
		return datatypes.ACCOUNT{}, fmt.Errorf("no se encontro la cuenta"), http.StatusNotFound

	}
	return User, nil, http.StatusOK
}
