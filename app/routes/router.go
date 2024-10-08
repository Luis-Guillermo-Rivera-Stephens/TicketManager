package routes

import (
	"fmt"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/handlers"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/middlewares"
)

func (api *API) InitRoutes() {
	fmt.Println("Starting the routes")

	//ROUTERS
	public := api.router.NewRoute().Subrouter()
	protected := api.router.NewRoute().Subrouter()
	account_protected := protected.NewRoute().Subrouter()

	protected.Use(middlewares.VerifyJWT)
	account_protected.Use(middlewares.VerifyEmail)

	login := public.NewRoute().Subrouter()
	login.Use(
		middlewares.ValidEmail,
		middlewares.VerifyEmail)

	login.HandleFunc("/account", handlers.GetAccount).Methods(http.MethodGet)

	register := public.NewRoute().Subrouter()
	register.Use(
		middlewares.SetHeaders,
		middlewares.ValidEmail,
		middlewares.VerifyEmailRegister,
		middlewares.ValidPassword,
		middlewares.HashPassword)

	register.HandleFunc("/account", handlers.Create_Account).Methods(http.MethodPost)

	account_protected.HandleFunc("/account", handlers.Delete_Account).Methods(http.MethodDelete)

	//ROUTER
	update_router := account_protected.NewRoute().Subrouter()
	update_router.Use(
		middlewares.PasswordAuthentification,
		middlewares.ValidPassword,
		middlewares.HashPassword)

	update_router.HandleFunc("/account", handlers.Update_Password).Methods(http.MethodPut)

	//ROUTER
	tickets := protected.NewRoute().Subrouter()
	tickets.Use(middlewares.AccountExistByID)

	//ROUTER
	department_tickets_router := tickets.NewRoute().Subrouter()
	department_tickets_router.Use(middlewares.VerifyDepartment)

	tickets.HandleFunc("/tickets/all", handlers.GetAllTickets).Methods(http.MethodGet)
	tickets.HandleFunc("/tickets/open", handlers.GetOpenTickets).Methods(http.MethodGet)
	tickets.HandleFunc("/tickets/closed", handlers.GetClosedTickets).Methods(http.MethodGet)
	tickets.HandleFunc("/tickets/unassigned", handlers.GetAllUnassignedTickets).Methods(http.MethodGet)
	tickets.HandleFunc("/tickets/owner", handlers.GetTicketsByOwner).Methods(http.MethodGet)

	department_tickets_router.HandleFunc("/tickets/department/{d_id}", handlers.GetTicketsByDepartment).Methods(http.MethodGet)
	department_tickets_router.HandleFunc("/tickets/unassigned/department/{d_id}", handlers.GetUnassignedTicketsByDepartment).Methods(http.MethodGet)

	//ROUTER
	PM_ticket_router := tickets.NewRoute().Subrouter()
	PM_ticket_router.Use(middlewares.IsPM)

	PM_ticket_router.HandleFunc("/tickets", handlers.Create_Ticket).Methods(http.MethodPost)

	one_ticket := tickets.NewRoute().Subrouter()
	one_ticket.Use(middlewares.TicketExist)
	one_ticket.HandleFunc("/tickets/update/owner", handlers.Update_Owner).Methods(http.MethodPut)
	one_ticket_status := one_ticket.NewRoute().Subrouter()
	one_ticket_status.Use(middlewares.ValidStatus)
	one_ticket_status.HandleFunc("/tickets/update/status/{s_id}", handlers.Update_Status).Methods(http.MethodPut)

}
