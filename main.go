package main

import (
	"github.com/rotta-f/ticketingApi/database"
	"github.com/rotta-f/ticketingApi/handlers"
	"github.com/rotta-f/ticketingApi/router"
	"log"
	"net/http"
	"time"
)

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		tFinal := time.Now().Sub(t1)

		log.Printf("[%s] %s %s", r.Method, r.RequestURI, tFinal.String())
	}
}

func main() {
	database.PrintDatabse()

	routerAuth := router.NewRouter()
	routerAuth.AddRoute("POST", "/v1/auth/login", handlers.AuthLogin)
	http.Handle("/v1/auth/", withLogging(router.UseRouter(routerAuth)))

	/*routerAuth.AddRoute("POST", `/v1/auth/+[\d]`, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("signup"))
	})*/

	http.ListenAndServe(":3000", nil)
}
