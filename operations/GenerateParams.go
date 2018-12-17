package operations

import (
	"encoding/json"
	"net/http"

	models "idempotence/models"
)

type ClientParams struct {
	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	INPUT *models.InputSpec
}

func NewClientParams(r *http.Request) *ClientParams {
	params := new(ClientParams)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params.INPUT)
	params.HTTPRequest = r
	return params
}