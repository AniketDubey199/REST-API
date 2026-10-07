package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingDetails(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Print(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func TimeDetails(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		fmt.Println("Request Took: ", time.Since(start))
	})
}
