package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type request struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type response struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

func getBooks(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "List of all books")
}
func getBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	io.WriteString(w, "List of books with id: "+id)
}
func createBook(w http.ResponseWriter, r *http.Request) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var res response
	res = response{
		Name:      req.Name,
		CreatedAt: time.Now().Format(time.UnixDate),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	io.WriteString(w, "Book id: "+id+" got deleted.")
}
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", getBooks)
	mux.HandleFunc("GET /books/{id}", getBook)
	mux.HandleFunc("POST /books", createBook)
	mux.HandleFunc("DELETE /books/{id}", deleteBook)

	http.ListenAndServe(":8080", mux)
}
