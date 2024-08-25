package routes

import (
	"fmt"
	"net/http"

	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/handlers"
	"github.com/Luis-Guillermo-Rivera-Stephens/TicketManager/app/middlewares"
)

func (api *API) InitRoutes() {
	fmt.Println("Starting the routes")

	public := api.router.NewRoute().Subrouter()
	protected := api.router.NewRoute().Subrouter()
	protected.Use(
		middlewares.VerifyJWT,
		middlewares.VerifyEmail)

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

	protected.HandleFunc("/account", handlers.Delete_Account).Methods(http.MethodDelete)

	update_router := protected.NewRoute().Subrouter()
	update_router.Use(
		middlewares.PasswordAuthentification,
		middlewares.HashPassword)

	update_router.HandleFunc("/account", handlers.Update_Password).Methods(http.MethodPut)
}
