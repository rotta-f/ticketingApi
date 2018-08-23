package main

import (
	"github.com/rotta-f/ticketingApi/handlers"
	"github.com/rotta-f/ticketingApi/router"
	"log"
	"net/http"
)

func main() {
	routerAuth := router.NewRouter()
	routerAuth.AddRoute("POST", "/v1/auth/login", handlers.AuthLogin)
	routerAuth.AddRoute("POST", "/v1/auth/signup", handlers.AuthSignup)
	http.Handle("/v1/auth/", handlers.WithLogging(router.UseRouter(routerAuth)))

	routerUser := router.NewRouter()
	routerUser.AddRoute("POST", "/v1/users/create/support", handlers.UserCreateSupport)
	routerUser.AddRoute("POST", "/v1/users/create/client", handlers.UserCreateClient)
	routerUser.AddRoute("PATCH", `/v1/users/edit/+[\d]`, handlers.UserUpdate)
	http.Handle("/v1/users/", handlers.WithLogging(handlers.WithContext(handlers.WithAuth(router.UseRouter(routerUser)))))

	routerTicket := router.NewRouter()
	routerTicket.AddRoute("POST", "/v1/tickets/create", handlers.CreateTicket)
	routerTicket.AddRoute("GET", "/v1/tickets", nil)
	routerTicket.AddRoute("GET", `/v1/tickets/+[\d]`, nil)
	routerTicket.AddRoute("PATCH", `/v1/tickets/+[\d]`, nil)
	routerTicket.AddRoute("POST", `/v1/tickets/+[\d]/close`, nil)
	routerTicket.AddRoute("POST", `/v1/tickets/+[\d]/archive`, nil)
	http.Handle("/v1/tickets/", handlers.WithLogging(handlers.WithContext(handlers.WithAuth(router.UseRouter(routerTicket)))))

	/*routerAuth.AddRoute("POST", `/v1/auth/+[\d]`, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("signup"))
	})*/

	log.Println("Ready to listen and serve on port 3000.")
	http.ListenAndServe(":3000", nil)
}
