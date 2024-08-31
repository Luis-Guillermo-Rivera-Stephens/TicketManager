package general

import (
	"fmt"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/data"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
)

func GetAccount(id int) (datatypes.ACCOUNT, error) {
	db, err := data.GetDataBase()
	if err != nil {
		return datatypes.ACCOUNT{}, err
	}
	var User datatypes.ACCOUNT
	result := db.Raw("EXEC GET_ACCOUNT ?", id).Scan(&User)

	if result.Error != nil {
		return datatypes.ACCOUNT{}, result.Error
	}
	if result.RowsAffected == 0 {
		return datatypes.ACCOUNT{}, fmt.Errorf("no se encontro la cuenta")

	}
	return User, nil
}
