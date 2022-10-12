package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type API struct{}

type BooksParams struct {
	Limit int `schema:"limit"`
	Offset int `schema:"offset"`
}

type PostBook struct {
	Title string `schema:"title"`
}

var (
	books = []string{"Boook 1", "Boook 2", "Boook 3"}
	decoder = schema.NewDecoder()
)

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {

	params := &BooksParams{}

	err := decoder.Decode(params, r.URL.Query())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if  params.Offset > len(books) || params.Offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Limit < 0 || params.Limit > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Offset == 0 && params.Limit == 0 {
		json.NewEncoder(w).Encode(books[:])
		return
	}

	json.NewEncoder(w).Encode(books[params.Offset:params.Limit])
}

func (a *API) getBook(w http.ResponseWriter, r *http.Request){

	pathParms := mux.Vars(r)

	idParam := pathParms["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	index := id -1

	if index < 0 || index > len(books) -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(books[index])

}

func (a *API) postBook(w http.ResponseWriter, r *http.Request){
	book := &PostBook{}

	err := json.NewDecoder(r.Body).Decode(book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	books = append(books, book.Title)
	w.WriteHeader(http.StatusCreated)


}