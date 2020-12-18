package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
	http.HandleFunc("/sql", databaseHealthCheck)
	http.HandleFunc("/docs", showDocumentation)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func databaseHealthCheck(w http.ResponseWriter, r *http.Request) {


	db, err := sql.Open("mysql", "admin:1234@pk-db-instance-1.cqbx0szwkzam.us-west-2.rds.amazonaws.com/parkmee")

	if (err != nil) {
		panic(err)
	}
	defer db.Close()

	pingResult := db.Ping()

	if( pingResult != nil) {
		panic(pingResult);
	}

	fmt.Fprint(w, err)
	fmt.Fprint(w, pingResult);

}

/*

	Route handlers

*/
func healthCheck(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "👍🏻")
}

func showDocumentation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><head></head><body><h1>Docs!</h1></body></html>")
}