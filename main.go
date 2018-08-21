package main

import (
	"github.com/rotta-f/ticketingApi/database"
	"net/http"
	"regexp"
	"time"
	"log"
	"github.com/rotta-f/ticketingApi/router"
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
	routerAuth.AddRoute("GET", "/v1/auth/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login"))
	})
	routerAuth.AddRoute("POST", "/v1/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("signup"))
	})
	routerAuth.AddRoute("POST", `/v1/auth/+[\d]`, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("signup"))
	})

	http.Handle("/v1/auth/", withLogging(router.UseRouter(routerAuth)))

	http.ListenAndServe(":3000", nil)

	rg, err := regexp.Compile(`^/ticket/+[\d]+/create/*$`)
	bool := rg.Match([]byte("/ticket/19/create"))
	println(bool, err)
	rg, err = regexp.Compile(`^/ticket/+[\d]/*$`)
	bool = rg.Match([]byte("/ticket/1"))
	println(bool, err)
}

