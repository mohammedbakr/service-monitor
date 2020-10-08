package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/k8-proxy/service-monitor/back-end/check"
)

func main() {
	check.ConnectDatabse()

	go check.F()

	router := mux.NewRouter()

	router.HandleFunc("/api/stats", check.Handlercheck).Methods("GET")
	router.HandleFunc("/api/urls/{id}", check.GetItemEndpoint).Methods("GET")

	router.HandleFunc("/api/urls", check.InsertUser).Methods("POST")

	router.HandleFunc("/api/urls/{id}", check.DeleteUser).Methods("DELETE")

	router.HandleFunc("/api/urls/{id}", check.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}
