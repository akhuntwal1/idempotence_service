package operations

import (
	"encoding/json"
	"net/http"

	models "idempotence/models"
)

type GenerateResponse struct {
	/*
	  In: Body
	*/
	Payload *models.OutputSpec `json:"body,omitempty"`

	StatusCode int
	
	ErrorMessage string
}

func NewGenerateResponse() *GenerateResponse {
	return &GenerateResponse{}
}

func (o *GenerateResponse) WithPayload(payload *models.OutputSpec) *GenerateResponse {
	o.Payload = payload
	return o
}

func (o *GenerateResponse) WithStatusCode_error(statuscode int, error_message string) *GenerateResponse {
	o.StatusCode = statuscode
	o. ErrorMessage = error_message
	return o
}


func (o *GenerateResponse) WriteResponse(w http.ResponseWriter) {
	if o.StatusCode == 0 {
		if o.Payload != nil {
			o.StatusCode = http.StatusOK
		} else {
			o.StatusCode = http.StatusInternalServerError
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(o.StatusCode)
	if o.Payload != nil {
		payload := o.Payload
		json.NewEncoder(w).Encode(payload)
	} else {
		json.NewEncoder(w).Encode(o.ErrorMessage)
	}
}