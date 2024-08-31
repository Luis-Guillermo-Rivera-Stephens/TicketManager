package general

import (
	"errors"
	"fmt"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
)

func ValidTicketInfo(t datatypes.TICKET) error {
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.T_Description == "" {
		return errors.New("description is empty")
	}
	if idDep, ok := datatypes.DepartmentsHash[t.Department]; !ok {
		fmt.Println(idDep)
		return errors.New("department must be between 1 and 10")
	}

	return nil
}
