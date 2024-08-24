package datatypes

import "time"

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
