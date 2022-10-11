package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/palermo/mux-api/api"
)

func handleIndex (w http.ResponseWriter, r *http.Request){
	// fmt.Fprintf(w, "hello world")
	json.NewEncoder(w).Encode("{\"message\": \"Hello World\"}")
}


func main() {

	router := mux.NewRouter() 

	a := &api.API{}
	a.RegisterRoutes(router)

	router.HandleFunc("/", handleIndex)
	svr := &http.Server {
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Listening ....")
	svr.ListenAndServe()

}