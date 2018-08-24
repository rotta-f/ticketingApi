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
	routerUser.AddRoute("PATCH", `/v1/users/+[\d]+`, handlers.UserUpdate)
	http.Handle("/v1/users/", handlers.WithLogging(handlers.WithContext(handlers.WithAuth(router.UseRouter(routerUser)))))

	routerTicket := router.NewRouter()
	routerTicket.AddRoute("POST", "/v1/tickets/create", handlers.CreateTicket)
	routerTicket.AddRoute("GET", "/v1/tickets/", handlers.GetTickets)
	routerTicket.AddRoute("GET", "/v1/tickets", handlers.GetTickets)
	routerTicket.AddRoute("GET", `/v1/tickets/+[\d]+`, handlers.GetTicketById)
	routerTicket.AddRoute("PATCH", `/v1/tickets/+[\d]+`, handlers.EditTicket)
	routerTicket.AddRoute("POST", `/v1/tickets/+[\d]+/close`, handlers.CloseTicket)
	routerTicket.AddRoute("POST", `/v1/tickets/+[\d]+/archive`, handlers.ArchiveTicket)
	http.Handle("/v1/tickets/", handlers.WithLogging(handlers.WithContext(handlers.WithAuth(router.UseRouter(routerTicket)))))
	http.Handle("/v1/tickets", handlers.WithLogging(handlers.WithContext(handlers.WithAuth(router.UseRouter(routerTicket)))))

	routerMessage := router.NewRouter()
	routerMessage.AddRoute("POST", `/v1/messages/ticket/+[\d]+`, handlers.NewMessageToTicket)
	routerMessage.AddRoute("GET", `/v1/messages/ticket/+[\d]+`, handlers.GetTicketMessages)
	routerMessage.AddRoute("GET", `/v1/messages/+[\d]+`, handlers.GetMessage)
	routerMessage.AddRoute("PATCH", `/v1/messages/+[\d]+`, handlers.EditMessage)
	http.Handle("/v1/messages/", handlers.WithLogging(handlers.WithContext(handlers.WithAuth(router.UseRouter(routerMessage)))))


	log.Println("Ready to listen and serve on port 3000.")
	http.ListenAndServe(":3000", nil)
}
