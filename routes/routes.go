package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Books struct {
	ID   int
	Name string
}

type request struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type response struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

var books []Books

func GetBooks(w http.ResponseWriter, _ *http.Request) {
	for _, val := range books {
		err := json.NewEncoder(w).Encode(&val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for _, val := range books {
		idint, _ := strconv.Atoi(id)
		if val.ID == idint {
			err := json.NewEncoder(w).Encode(&val)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, Books{req.Id, req.Name})
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
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	io.WriteString(w, "Book id: "+id+" got deleted.")
}
