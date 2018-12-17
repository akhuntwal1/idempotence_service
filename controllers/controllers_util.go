package controllers

import (
    "fmt"
    "strings"
    "strconv"
    "database/sql"

    _ "github.com/lib/pq"

    . "idempotence/constants"
)

func NewUUIDs (identifier string, count int64) []string{

	var UUIDlist []string
    var iterator int64

    UUIDlist = FetchUUIDlist(identifier)

    if int64(len(UUIDlist)) >= count{
    	//no need to generate new uuids. Use existing ones
        return UUIDlist[:count]
    } else{
    	//generate count - len number of uuids

    	for iterator=int64(len(UUIDlist)); iterator < count ; iterator++{

    		uuid := strings.Join([]string{identifier,
     	     strconv.FormatInt(iterator, 10)}, "-")
   
   			db, err := sql.Open("postgres", DB_INFO)

   			_, err = db.Exec(STMT_INSERT_UUIDS, identifier, uuid)
    		checkErr(err)

        }
        return FetchUUIDlist(identifier)   
    }
}

func FetchUUIDlist(identifier string) []string{

	db, err := sql.Open("postgres", DB_INFO)

	var UUIDlist []string

	rows, err := db.Query(QUERY_UUIDS_GIVEN_CLIENT, identifier)
    checkErr(err)

    var uuid string
 	for rows.Next() {
            err = rows.Scan(&uuid)
            checkErr(err)
            fmt.Println(uuid)
            UUIDlist = append(UUIDlist, uuid)
    }
    return UUIDlist
}

func CheckIdentifierExistence (identifier string) bool{

	db, err := sql.Open("postgres", DB_INFO)

	rows, err := db.Query(QUERY_CLIENT_ID, identifier)
    checkErr(err)

    if rows.Next(){
    	return true
    } else {
    	return false
    }
}
