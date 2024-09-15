package datatypes

type TICKET struct {
	ID_Ticket     int    `json:"id_ticket"`
	Title         string `json:"title"`
	T_Description string `json:"t_description"`
	ID_Account    int    `json:"id_account"`
	Owner         string `json:"owner"`
	Department    string `json:"department"`
	Status        string `json:"status"`
	Priority      string `json:"priority"`
}
