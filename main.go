package main

import (
	"net/http"

	"github.com/AniketDubey199/middleware"
	"github.com/AniketDubey199/routes"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", routes.GetBooks)
	mux.HandleFunc("GET /books/{id}", routes.GetBook)
	mux.HandleFunc("POST /books", routes.CreateBook)
	mux.HandleFunc("DELETE /books/{id}", routes.DeleteBook)

	handler := middleware.TimeDetails(mux)
	handler = middleware.LoggingDetails(handler)

	http.ListenAndServe(":8080", handler)
}
