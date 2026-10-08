package routes

import (
	"context"
	"encoding/json"
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

func GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()
	for _, val := range books {
		select {
		case <-ctx.Done():
			http.Error(w, "Timeout", http.StatusBadRequest)
			return
		default:
		}
		err := json.NewEncoder(w).Encode(&val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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
	ctx.Done()
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, Books{ID: req.Id, Name: req.Name})
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
	ctx.Done()
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for i, val := range books {
		idint, _ := strconv.Atoi(id)
		if val.ID == idint {
			json.NewEncoder(w).Encode(&val)
			books = append(books[:i], books[i+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusBadGateway)
	json.NewEncoder(w).Encode("There is no content")
}
