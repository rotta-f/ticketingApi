package handlers

import (
	"context"
	"github.com/rotta-f/ticketingApi/database"
	"github.com/rotta-f/ticketingApi/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *logResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Log all request
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		lrw := &logResponseWriter{w, http.StatusOK}

		next.ServeHTTP(lrw, r)
		tFinal := time.Now().Sub(t1)

		log.Printf("[%s]\t%s\tstatus:%d\ttime:%s", r.Method, r.RequestURI, lrw.statusCode, tFinal.String())
	}
}

// Add a context to Request structure
func WithContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, "values", map[string]interface{}{})

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}

func CtxSetValue(r *http.Request, key string, value interface{}) {
	v := r.Context().Value("values").(map[string]interface{})
	v[key] = value
}

func CtxGetValue(r *http.Request, key string) interface{} {
	v := r.Context().Value("values").(map[string]interface{})
	return v[key]
}

// Middleware to verify authentication and store the current user to request context
func WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authT := strings.Split(r.Header.Get("Authorization"), " ")

		// Check if token exist
		if len(authT) == 2 && authT[0] == "Bearer" {

			// Get the current user and store it in context
			a := database.GetAuthToken(authT[1])
			if a != nil {
				CtxSetValue(r, STORE_AUTH, a.User)
				next.ServeHTTP(w, r)
				return
			}
		}
		utils.WriteError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Invalid token")
	}
}
