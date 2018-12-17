package routes

import (

	"github.com/gorilla/mux"
	"idempotence/handlers"
)

/**
	This function initialises a mux router with set of handler functions
	returns: a router with configured routes and handler functions
**/
func ConfigRouter() *mux.Router {

    versionPath := "/v1/"
    router := mux.NewRouter().StrictSlash(true)

	IdempotenceRouter := router.PathPrefix(versionPath).Subrouter()
	IdempotenceRouter.HandleFunc("/", handlers.GenerateHandler).Methods("POST")
	IdempotenceRouter.HandleFunc("/{client_identifier}", handlers.GetHandler).Methods("GET")
	IdempotenceRouter.HandleFunc("/{client_identifier}", handlers.DeleteHandler).Methods("DELETE")

	return router
}