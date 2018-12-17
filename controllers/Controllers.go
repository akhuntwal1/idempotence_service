package controllers

import (
    "database/sql"

    _ "github.com/lib/pq"

   	"idempotence/models"
	"idempotence/operations"
	. "idempotence/constants"
)


func GenerateUUIDs (params operations.ClientParams) *operations.GenerateResponse {

	Res := models.OutputSpec{
		params.INPUT.ClientIdentifier,
		params.INPUT.Count,
	    NewUUIDs(params.INPUT.ClientIdentifier, params.INPUT.Count)}
	return operations.NewGenerateResponse().WithPayload(&Res)
}

func GetUUIDs (identifier string) *operations.GenerateResponse {

	var UUIDlist []string
	UUIDlist = FetchUUIDlist(identifier)
	if (len(UUIDlist) != 0){
    	Res := models.OutputSpec{
    		identifier,
    		int64(len(UUIDlist)),
    		UUIDlist}
    	return operations.NewGenerateResponse().WithPayload(&Res)
    } else{
    	return operations.NewGenerateResponse().WithStatusCode_error(404, MESSAGE_FOR_404)
    }
}

func DeleteUUIDs (identifier string) *operations.GenerateResponse {

	if !CheckIdentifierExistence(identifier){
		return operations.NewGenerateResponse().WithStatusCode_error(404, MESSAGE_FOR_404)
	} else{
		db, err := sql.Open("postgres", DB_INFO)
		_, err = db.Exec(STMT_DELETE_IDENTIFIER, identifier)
        checkErr(err)

        return operations.NewGenerateResponse().WithStatusCode_error(200, MESSAGE_FOR_200)
	}
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
