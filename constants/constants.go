package constants

const (
        DB_INFO = "user=postgres password=xenopsila dbname=idempotence sslmode=disable port=5000"

        MESSAGE_FOR_404 = "No such client identifier"
        MESSAGE_FOR_200 = "Task completed"

        STMT_DELETE_IDENTIFIER = "DELETE FROM uuids WHERE client_id=$1"
        STMT_INSERT_UUIDS = "INSERT INTO uuids(client_id, uuid) VALUES($1,$2);"

        QUERY_UUIDS_GIVEN_CLIENT = `SELECT uuid FROM uuids WHERE client_id=$1`
        QUERY_CLIENT_ID = `SELECT client_id FROM uuids WHERE client_id=$1`
)

