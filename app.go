package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/poc/url-shortner/cassandra"
	"github.com/poc/url-shortner/handler"
	"log"
	"net/http"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
func main() {
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", heartbeat)
//	router.HandleFunc("/create", handler.GetHash)
	router.HandleFunc("/toko/{pattern}", handler.RedirectURL)
	log.Println("Application started")
	//cassandra.CreateKeys()
	log.Fatal(http.ListenAndServe(":8090", router))
}
