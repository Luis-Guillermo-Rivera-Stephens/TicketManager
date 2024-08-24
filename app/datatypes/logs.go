package datatypes

import "time"

type LOGS struct {
	ID_Log          int       `json:"id_log"`
	ID_Ticket       int       `json:"id_ticket"`
	Log_Date        time.Time `json:"log_date"`
	Log_Description string    `json:"log_description"`
}
