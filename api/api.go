package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type API struct{}

var books = []string{"Boook 1", "Boook 2", "Boook 3"}

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {

	limitParant := r.URL.Query().Get("limit")

	if limitParant == "" {
		json.NewEncoder(w).Encode(books)
		return
	}

	limit, err := strconv.Atoi(limitParant)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if limit < 0 || limit > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(books[:limit])
}

func (a *API) getBook(w http.ResponseWriter, r *http.Request){
}