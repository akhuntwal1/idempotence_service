package handlers

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"

	"idempotence/controllers"
	"idempotence/operations"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Came to Generate handler")

	params := operations.NewClientParams(r)
	fmt.Println("Created Params %d, %d ", params.INPUT.ClientIdentifier, params.INPUT.Count)

	response := controllers.GenerateUUIDs(*params)
	fmt.Println("Made response %d", response.Payload.UUIDList)


	response.WriteResponse(w)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Came to get handler")

	vars := mux.Vars(r)
	fmt.Println("Received client identifier %d", vars["client_identifier"])

	response := controllers.GetUUIDs(vars["client_identifier"])

	response.WriteResponse(w)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Came to get handler")

	vars := mux.Vars(r)
	fmt.Println("Received client identifier %d", vars["client_identifier"])

	response := controllers.DeleteUUIDs(vars["client_identifier"])

	response.WriteResponse(w)
}