package datatypes

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type ACCOUNT struct {
	ID_Account    int    `json:"id_account"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Passwd        string `json:"passwd"`
	IsStarted     int    `json:"isstarted"`
	Department_FK int    `json:"department_fk"`
	IsPM          int    `json:"ispm"`
}

func (a *ACCOUNT) VerifPassword(parampassword string) error {
	if parampassword == "" {
		return errors.New("password is empty")
	}
	err := bcrypt.CompareHashAndPassword([]byte(a.Passwd), []byte(parampassword))
	return err
}
