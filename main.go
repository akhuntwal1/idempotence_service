package main

import (
	"fmt"
	"net/http"
	"database/sql"

    _ "github.com/mattn/go-sqlite3"

	routes "idempotence/routes"
	. "idempotence/constants"
)

func main() {

    db, err := sql.Open("postgres", DB_INFO)

    sqlStmt := "CREATE TABLE uuids (client_id VARCHAR(64), uuid VARCHAR(64) NULL);"

    _, err = db.Exec(sqlStmt)
    checkErr(err)

	router := routes.ConfigRouter()
	Port := 7003
    fmt.Println("Starting idempotence service at port %d", Port)

	http.ListenAndServe(fmt.Sprintf(":%d", Port), router)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}