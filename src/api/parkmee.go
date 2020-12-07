package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

/*

	All routes

	/
	/healthcheck
	/home
	/create
	/edit/{uuid}

*/
func handleRequests() {
	http.HandleFunc("/healthcheck", healthCheck)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

/*

	Route handlers

*/
func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "👍🏻")
}

