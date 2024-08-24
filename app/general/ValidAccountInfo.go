package general

import (
	"fmt"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/datatypes"
)

func ValidAccountInfo(u datatypes.ACCOUNT) error {
	if u.Name == "" {
		return fmt.Errorf("el campo 'Name' es obligatorio")
	}

	if u.Email == "" {
		return fmt.Errorf("el campo 'Email' es obligatorio")
	}
	if u.Passwd == "" {
		return fmt.Errorf("el campo 'Password' es obligatorio")
	}
	if u.Department_FK < 0 && u.Department_FK > 8 {
		return fmt.Errorf("el campo 'Department' no es válido")
	}

	return nil
}
