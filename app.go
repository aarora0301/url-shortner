package main

import (
	"github.com/gorilla/mux"
	"github.com/poc/url-shortner/cassandra"
	routes "github.com/poc/url-shortner/router"
	"log"
	"net/http"
)

/**
Main Application
*/
func main() {
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()
	router := mux.NewRouter().StrictSlash(true)
	/**
	  Expose all the routes
	*/
	routes.HandleHttpRequest(router)
	log.Println("Application started")
	//cassandra.CreateKeys()
	log.Fatal(http.ListenAndServe(":8090", router))
	log.Println("Application started on : 8090")
}
