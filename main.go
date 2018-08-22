package main

import (
	"github.com/rotta-f/ticketingApi/handlers"
	"github.com/rotta-f/ticketingApi/router"
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

	/*routerAuth.AddRoute("POST", `/v1/auth/+[\d]`, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("signup"))
	})*/

	http.ListenAndServe(":3000", nil)
}
