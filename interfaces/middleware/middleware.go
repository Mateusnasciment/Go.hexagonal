package middleware

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("Finished handling %s %s", r.Method, r.URL.Path)
	})
}
