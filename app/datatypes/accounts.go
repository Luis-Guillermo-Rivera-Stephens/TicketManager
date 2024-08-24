package datatypes

import "time"

type ACCOUNT struct {
	ID_Account    int    `json:"id_account"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Passwd        string `json:"passwd"`
	IsStarted     int    `json:"isstarted"`
	Department_FK int    `json:"department_fk"`
	IsPM          int    `json:"ispm"`
}

type TICKET struct {
	ID_Ticket     int       `json:"id_ticket"`
	Title         string    `json:"title"`
	T_Description string    `json:"t_description"`
	CreationDate  time.Time `json:"creation_date"`
	ID_Account    int       `json:"id_account"`
	OwnerName     string    `json:"owner"`
	Department    string    `json:"department"`
	Status        string    `json:"status"`
}

type LOGS struct {
	ID_Log          int       `json:"id_log"`
	ID_Ticket       int       `json:"id_ticket"`
	Log_Date        time.Time `json:"log_date"`
	Log_Description string    `json:"log_description"`
}

/*
currentTime := time.Now()
formattedTime := currentTime.Format("2006-01-02 15:04:05")

*/
