// package main

// import "net/http"

// func myCorsHandler(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
// 		if r.Method == "OPTIONS" {
// 			w.Header().Set("Access-Control-Allow-Headers", "X-PINGOTHER, Content-Type")
// 			w.Header().Set("Access-Control-Max-Age", "86400")
// 			w.WriteHeader(200)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

package main

import (
	"log"
	"net/http"
	"time"
)

var allowedOrigin = "*"

const (
	allowedHeaders = "X-PINGOTHER, Content-Type"
	maxAge         = "86400"
)

// myCorsHandler is a middleware that adds CORS headers to the HTTP responses
func myCorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set Access-Control-Allow-Origin header to allow requests from a specific origin
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)

		// If the request method is OPTIONS, set appropriate CORS headers and return immediately
		switch r.Method {
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
			w.Header().Set("Access-Control-Max-Age", maxAge)
			w.WriteHeader(http.StatusOK)
			return
		default:
			// Call the next handler in the chain to handle the actual request
			next.ServeHTTP(w, r)
		}
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(err)
			}
		}()
		t1 := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.RequestURI, time.Since(t1))
	})
}
