package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/newyu_transactions/config"
)

// Init initializes gorrila router
func Init() {
	router := mux.NewRouter()
	// Routes consist of a path and a handler function.
	router.HandleFunc("/v1/transaction", transaction).Methods(http.MethodPost)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(viper.GetString(config.ListeningPort), router))
}
